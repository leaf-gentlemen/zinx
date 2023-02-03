package znet

import (
	"bytes"
	"encoding/binary"

	"github.com/leaf-gentlemen/zinx/ziface"
	"github.com/pkg/errors"
)

// DataPack
// @Description: 消息拆包，封包（tlv）。采用小端
type DataPack struct {
}

func NewDataPack() ziface.IDataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() int {
	headLen := 8
	return headLen
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	//  写入消息长度
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, errors.WithStack(err)
	}
	//  写入消息 ID
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, errors.WithStack(err)
	}
	//  写入消息内容
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, errors.WithStack(err)
	}
	return buf.Bytes(), nil
}

func (d *DataPack) UnPack(buf []byte) (ziface.IMessage, error) {
	bufRead := bytes.NewReader(buf)
	msg := &Message{}

	if err := binary.Read(bufRead, binary.LittleEndian, &msg.Len); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := binary.Read(bufRead, binary.LittleEndian, &msg.ID); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := binary.Read(bufRead, binary.LittleEndian, &msg.Data); err != nil {
		return nil, errors.WithStack(err)
	}

	// TODO 处理包体过大
	return msg, nil
}
