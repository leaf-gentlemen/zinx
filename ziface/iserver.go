package ziface

// IServer
// @Description: server 抽象层
type IServer interface {
	//
	// Start
	//  @Description: 启动服务器
	//
	Start()
	//
	// Stop
	//  @Description: 停止服务器
	//
	Stop()
	//
	// Serve
	//  @Description: 运行服务器
	//
	Serve()
	//
	// AddRouter
	//  @Description: 添加路由
	//  @param uint32
	//  @param IRouter
	//  @return error
	//
	AddRouter(uint32, IHandler) error
	//
	// GetConnManager
	//  @Description: 获取连接管理器
	//  @return IConnManager
	//
	GetConnManager() IConnManager
	//
	// SetOnConnStart
	//  @Description:设置该Server的连接创建时Hook函数
	//  @param func(IConnection)
	//
	SetOnConnStart(func(IConnection))
	//
	// SetOnConnStop
	//  @Description: 设置该Server的连接断开时的Hook函数
	//  @param func(IConnection)
	//
	SetOnConnStop(func(IConnection))
	//
	// CallOnConnStart
	//  @Description: 调用连接OnConnStart Hook函数
	//  @param conn
	//
	CallOnConnStart(conn IConnection)
	//
	// CallOnConnStop
	//  @Description: 	调用连接OnConnStop Hook函数
	//  @param conn
	//
	CallOnConnStop(conn IConnection)
}
