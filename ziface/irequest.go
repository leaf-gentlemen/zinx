package ziface

// IRequest
// @Description: 请求体
type IRequest interface {
	//
	// GetConnection
	//  @Description: 获取绑定的 connect 链接
	//  @return IConnection
	//
	GetConnection() IConnection
	//
	// GetData
	//  @Description: 获取数据
	//  @return []byte
	//
	GetData() []byte
	//
	// GetMessage
	//  @Description: 获取 message 对象
	//  @return IMessage
	//
	GetMessage() IMessage
}
