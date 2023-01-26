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
}
