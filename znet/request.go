package znet

import "zinx/ziface"

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage
	cnt  int
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMessage() ziface.IMessage {
	return r.msg
}
