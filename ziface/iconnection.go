package ziface

import "net"

type IConnection interface {
	//
	// Start
	//  @Description: 启动连接
	//
	Start()

	//
	// Stop
	//  @Description: 关连接
	//
	Stop()

	//
	// GetConn
	//  @Description:  获取链接
	//  @return *net.TCPConn
	//
	GetConn() *net.TCPConn

	//
	// GetConnID
	//  @Description: 获取连接ID
	//  @return uint32
	//
	GetConnID() uint32

	//
	// GetRemoteAddr
	//  @Description: 获取连接地址
	//  @return net.Addr
	//
	GetRemoteAddr() net.Addr

	//
	// Send
	//  @Description: 发送客户端数据
	//  @param buf
	//  @return error
	//
	Send(uint32, []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
