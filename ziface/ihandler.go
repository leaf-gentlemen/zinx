package ziface

type IHandler interface {
	//
	// DoMessage
	//  @Description: 执行对应的路由处理逻辑
	//  @param msgID
	//  @param r
	//
	DoMessage(msgID uint32, r IRequest) error
	//
	// AddRoute
	//  @Description: 添加路由
	//  @param msgID
	//  @param r
	//
	AddRoute(msgID uint32, r IRouter) error
}
