package ziface

type IRouter interface {
	//
	// DoMessage
	//  @Description: 执行对应的路由处理逻辑
	//  @param r
	//
	DoMessage(IRequest) error
	//
	// AddRoute
	//  @Description: 添加路由
	//  @param msgID
	//  @param r
	//
	AddRoute(uint32, IHandler) error
	//
	// StartWorkerPool
	//  @Description: 启动worker线程池
	//
	StartWorkerPool()
	//
	// SendMsgToTaskQueue
	//  @Description: 投递消息队列
	//  @param IRequest
	//
	SendMsgToTaskQueue(IRequest)
}
