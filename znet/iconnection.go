package znet

import (
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
