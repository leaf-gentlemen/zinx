package znet

import (
	"log"
	"net"
	"zinx/ziface"
)

type Connection struct {
	// Conn 当前连接的socket TCP
	Conn *net.TCPConn
	// ConnID 当前连接ID
	ConnID uint32
	// isClose 是否关闭
	isClose bool
	// handle 绑定的业务方法
	handle ziface.HandleFunc
	// ExitChan 通知当前连接已退出
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, handleFunc ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClose:  false,
		handle:   handleFunc,
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

func (c *Connection) Send(buf []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) StartReader() {
	log.Printf("reader goruntine is running...\n")
	defer log.Printf("connID:%d, reader is exit, remote addr is %s\n",
		c.GetConnID(), c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.GetConn().Read(buf)
		if err != nil {
			log.Printf("recv fail err:%s\n", err)
			continue
		}

		if err := c.handle(c.GetConn(), buf, cnt); err != nil {
			log.Printf("handel function fail err:%s\n", err)
		}
	}
}
