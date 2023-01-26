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
	if err := conn.Send(201, []byte("hello word...")); err != nil {
		logger.Error("send client fail", zap.Error(err))
	}
}
