package ziface

type IConnManager interface {
	//
	// Add
	//  @Description: 添加连接先
	//  @param conn
	//
	Add(conn IConnection)
	//
	// Remove
	//  @Description: 移除连接
	//  @param connID
	//
	Remove(connID uint32)
	//
	// Get
	//  @Description: 获取连接
	//  @param connID
	//  @return IConnection
	//  @return bool
	//
	Get(connID uint32) (IConnection, bool)
	//
	// Len
	//  @Description: 连接长度
	//  @return int
	//
	Len() int
	//
	// Clear
	//  @Description: 情况连接
	//
	Clear()
}
