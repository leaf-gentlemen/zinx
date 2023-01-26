package main

import (
	"zinx/utils"
	"zinx/ziface"
	"zinx/znet"

	"go.uber.org/zap"
)

type Router struct {
	znet.BaseRouter
}

func (r *Router) Handle(req ziface.IRequest) {
	logger := utils.Logger
	logger.Debug("client msg", zap.Any("data", string(req.GetData())))
	conn := req.GetConnection()
	if err := conn.Send(1, []byte("ping...")); err != nil {
		logger.Error("send message client fail", zap.Error(err))
	}
}

func main() {
	utils.InitConf("/Users/yesheng/zinx/demo/config")
	s := znet.NewServe(utils.Interface().Name)

	r := &Router{}
	s.AddRouter(r)
	s.Serve()
}
