package ziface

type IRouter interface {
	//
	// DoMessage
	//  @Description: 执行对应的路由处理逻辑
	//  @param r
	//
	DoMessage(r IRequest) error
	//
	// AddRoute
	//  @Description: 添加路由
	//  @param msgID
	//  @param r
	//
	AddRoute(msgID uint32, r IHandler) error
}
