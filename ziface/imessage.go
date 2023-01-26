package ziface

type IMessage interface {
	//
	// GetMsgID
	//  @Description: 获取消息 ID
	//  @return uint32
	//
	GetMsgID() uint32
	//
	// GetMsgLen
	//  @Description: 获取消息长度
	//  @return uint32
	//
	GetMsgLen() uint32
	//
	// GetData
	//  @Description: 获取消息内容
	//  @return []byte
	//
	GetData() []byte
	//
	// SetMsgID
	//  @Description: 设置消息 ID
	//  @param id
	//
	SetMsgID(id uint32)
	//
	// SetMsgLen
	//  @Description: 设置消息长度
	//  @param msgLen
	//
	SetMsgLen(msgLen uint32)
	//
	// SetData
	//  @Description: 设置消息内容
	//  @param data
	//
	SetData(data []byte)
}
