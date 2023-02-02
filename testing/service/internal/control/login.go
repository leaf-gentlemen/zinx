package control

import (
	"github.com/leaf-gentlemen/zinx/utils"
	"github.com/leaf-gentlemen/zinx/ziface"
	"github.com/leaf-gentlemen/zinx/znet"

	"go.uber.org/zap"
)

type Login struct {
	znet.BaseHandler
}

func (l *Login) Handle(r ziface.IRequest) {
	logger := utils.Logger
	conn := r.GetConnection()
	logger.Debug("client message", zap.String("data", string(r.GetData())))
	var msgID uint32 = 201
	if err := conn.SendMsg(msgID, []byte("login succeed...")); err != nil {
		logger.Error("send client fail", zap.Error(err))
	}
}
