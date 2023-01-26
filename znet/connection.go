package znet

import (
	"io"
	"log"
	"net"
	"zinx/utils"
	"zinx/ziface"

	"github.com/pkg/errors"

	"go.uber.org/zap"
)

type Connection struct {
	// Conn 当前连接的socket TCP
	Conn *net.TCPConn
	// ConnID 当前连接ID
	ConnID uint32
	// isClose 是否关闭
	isClose bool
	// handle 绑定的业务方法
	handle ziface.IRouter
	// ExitChan 通知当前连接已退出
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, handle ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClose:  false,
		handle:   handle,
		ExitChan: make(chan bool),
	}
}

func (c *Connection) Start() {
	log.Printf("Conn Start()... ConnID:%d \n", c.GetConnID())
	// 启动读数据业务
	c.StartReader()
	// TODO 启动写的业务
}

func (c *Connection) Stop() {
	log.Printf("Conn stop()... ConnID = %d \n", c.GetConnID())
	if c.isClose {
		return
	}

	c.isClose = true
	if err := c.Conn.Close(); err != nil {
		log.Printf("err:%s\n", err)
	}
	close(c.ExitChan)
}

func (c *Connection) GetConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(msgID uint32, buf []byte) error {
	if c.isClose {
		return errors.WithStack(errors.New("conn is clone"))
	}

	msg := NewMessage(msgID, buf)
	dp := NewDataPack()
	bufData, err := dp.Pack(msg)
	if err != nil {
		return err
	}

	_, err = c.GetConn().Write(bufData)
	if err != nil {
		return err
	}
	return nil
}

func (c *Connection) StartReader() {
	logger := utils.Logger
	log.Printf("reader goruntine is running...\n")
	defer log.Printf("connID:%d, reader is exit, remote addr is %s\n",
		c.GetConnID(), c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		dp := NewDataPack()
		head := make([]byte, dp.GetHeadLen())
		cnt, err := io.ReadFull(c.GetConn(), head)
		if err != nil || cnt != dp.GetHeadLen() {
			if err.Error() == "EOF" {
				logger.Debug("client msg EOF")
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
			if err := c.handle.DoMessage(r); err != nil {
				logger.Error("do message fail", zap.Error(err))
			}
		}(r)
	}
}
