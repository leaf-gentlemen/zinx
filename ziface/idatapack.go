package ziface

type IDataPack interface {
	//
	// GetHeadLen
	//  @Description: 获取消息头部长度
	//  @return uint32
	//
	GetHeadLen() int
	//
	// Pack
	//  @Description: 消息封包
	//  @param IMessage
	//  @return []byte
	//  @return error
	//
	Pack(IMessage) ([]byte, error)
	//
	// UnPack
	//  @Description: 消息解包
	//  @param []byte
	//  @return IMessage
	//  @return error
	//
	UnPack([]byte) (IMessage, error)
}
