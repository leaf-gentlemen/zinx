package znet

type Message struct {
	ID   uint32
	Len  uint32
	Data []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		ID:   id,
		Len:  uint32(len(data)),
		Data: data,
	}
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetMsgLen() uint32 {
	return m.Len
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(id uint32) {
	m.ID = id
}

func (m *Message) SetMsgLen(msgLen uint32) {
	m.Len = msgLen
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
