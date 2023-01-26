package control

import (
	"zinx/utils"
	"zinx/ziface"
	"zinx/znet"

	"go.uber.org/zap"
)

type Login struct {
	znet.BaseHandler
}

func (l *Login) Handle(r ziface.IRequest) {
	logger := utils.Logger
	conn := r.GetConnection()
	logger.Debug("client message", zap.String("data", string(r.GetData())))
	if err := conn.Send(200, []byte("login succeed...")); err != nil {
		logger.Error("send client fail", zap.Error(err))
	}
}
