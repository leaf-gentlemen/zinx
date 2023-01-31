package control

import (
	"zinx/utils"
	"zinx/ziface"
	"zinx/znet"

	"go.uber.org/zap"
)

type User struct {
	znet.BaseHandler
}

func (l *User) Handle(r ziface.IRequest) {
	logger := utils.Logger
	conn := r.GetConnection()
	logger.Debug("client message", zap.String("data", string(r.GetData())))
	var msgID uint32 = 201
	if err := conn.SendMsg(msgID, []byte("hello word...")); err != nil {
		logger.Error("send client fail", zap.Error(err))
	}
}
