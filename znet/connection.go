package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"

	"github.com/pkg/errors"

	"go.uber.org/zap"
)

type Connection struct {
	// srv 当前连接的 server
	srv ziface.IServer
	// Conn 当前连接的socket TCP
	conn *net.TCPConn
	// ConnID 当前连接ID
	connID uint32
	// isClose 是否关闭
	isClose bool
	// handle 绑定的业务方法
	router ziface.IRouter
	// exitChan 通知当前连接已退出
	exitChan chan bool
	// msgChan 写业务消息管道
	msgChan chan []byte
	// msgBuffChan 带缓存写业务消息管道
	msgBuffChan chan []byte
}

func NewConnection(srv ziface.IServer, conn *net.TCPConn, connID uint32, handle ziface.IRouter) *Connection {
	c := &Connection{
		conn:        conn,
		connID:      connID,
		isClose:     false,
		router:      handle,
		exitChan:    make(chan bool),
		msgChan:     make(chan []byte),
		srv:         srv,
		msgBuffChan: make(chan []byte, utils.Interface().MsgBuffChanLen),
	}
	c.srv.GetConnManager().Add(c)
	return c
}

func (c *Connection) Start() {
	logger := utils.Logger
	logger.Debug(fmt.Sprintf("Conn start()... ConnID = %d \n", c.GetConnID()))
	// 启动读数据业务
	go c.StartReader()
	go c.StartWrite()
	c.srv.CallOnConnStart(c)
}

func (c *Connection) Stop() {
	logger := utils.Logger
	logger.Debug(fmt.Sprintf("Conn stop()... ConnID = %d \n", c.GetConnID()))
	if c.isClose {
		return
	}

	c.isClose = true
	if err := c.conn.Close(); err != nil {
		logger.Error("conn close fail", zap.Error(err))
	}

	c.srv.CallOnConnStop(c)
	c.exitChan <- true
	c.srv.GetConnManager().Remove(c.GetConnID())
	close(c.exitChan)
	close(c.msgChan)
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnID() uint32 {
	return c.connID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) StartReader() {
	logger := utils.Logger
	logger.Debug("reader goroutine is running...")
	defer logger.Debug(fmt.Sprintf("connID:%d, reader is exit, remote addr is %s\n",
		c.GetConnID(), c.GetRemoteAddr().String()))
	defer c.Stop()

	for {
		dp := NewDataPack()
		head := make([]byte, dp.GetHeadLen())
		cnt, err := io.ReadFull(c.GetConn(), head)
		if err != nil || cnt != dp.GetHeadLen() {
			if _, ok := err.(net.Error); ok || err == io.EOF {
				logger.Warn("client close", zap.String("remote", c.GetConn().RemoteAddr().String()))
				break
			}
			logger.Error("reader head fail.", zap.Error(err))
			continue
		}

		pack, err := dp.UnPack(head)
		if err != nil {
			logger.Error("unpack  fail.", zap.Error(err))
			continue
		}

		msgData := make([]byte, pack.GetMsgLen())
		cnt, err = io.ReadFull(c.GetConn(), msgData)
		if err != nil || cnt != int(pack.GetMsgLen()) {
			logger.Error("reader data fail.", zap.Error(err), zap.Int("cnt", cnt))
			continue
		}

		r := &Request{
			conn: c,
			msg:  NewMessage(pack.GetMsgID(), msgData),
		}

		go func(r ziface.IRequest) {
			c.router.SendMsgToTaskQueue(r)
		}(r)
	}
}

func (c *Connection) StartWrite() {
	logger := utils.Logger
	logger.Debug("write goroutine running...")
	defer logger.Debug(fmt.Sprintf("connID:%d, write is exit, remote addr is %s\n",
		c.GetConnID(), c.GetRemoteAddr().String()))

	for {
		select {
		case msg := <-c.msgChan:
			_, err := c.GetConn().Write(msg)
			if err != nil {
				logger.Error("write client msg fail", zap.Error(err))
				break
			}
		case msg, ok := <-c.msgBuffChan:
			if ok {
				_, err := c.GetConn().Write(msg)
				if err != nil {
					logger.Error("write client msg fail", zap.Error(err))
					break
				}
			} else {
				logger.Error("msgBuffChan is close")
				break
			}

		case <-c.exitChan:
			return
		}
	}
}

func (c *Connection) SendMsg(msgID uint32, buf []byte) error {
	if c.isClose {
		return errors.WithStack(errors.New("conn is clone"))
	}
	msg := NewMessage(msgID, buf)
	dp := NewDataPack()
	bufData, err := dp.Pack(msg)
	if err != nil {
		return err
	}
	c.msgChan <- bufData
	return nil
}

func (c *Connection) SendBuffMsg(msgID uint32, buf []byte) error {
	if c.isClose {
		return errors.WithStack(errors.New("conn is clone"))
	}

	dp := NewDataPack()
	bufData, err := dp.Pack(NewMessage(msgID, buf))
	if err != nil {
		return err
	}
	c.msgBuffChan <- bufData
	return nil
}
