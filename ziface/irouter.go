package ziface

// IRouter
// @Description: 路由抽象层
type IRouter interface {
	//
	// PerHandle
	//  @Description: 处理 conn 业务之前的钩子方法代Hook
	//  @param request
	//
	PerHandle(request IRequest)
	//
	// Handle
	//  @Description: 处理 conn 业务的主方法 Hook
	//  @param request
	//
	Handle(request IRequest)
	//
	// PostHandle
	//  @Description: 处理 conn 业务之后的钩子方法 Hook
	//  @param request
	//
	PostHandle(request IRequest)
}
