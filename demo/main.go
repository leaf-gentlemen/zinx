package main

import (
	"log"
	"zinx/utils"
	"zinx/ziface"
	"zinx/znet"
)

type Router struct {
	znet.BaseRouter
}

func (r *Router) Handle(req ziface.IRequest) {
	log.Printf("client msg %s \n", req.GetData())
	conn := req.GetConnection()
	if _, err := conn.GetConn().Write(req.GetData()); err != nil {
		log.Printf(" write client fail err:%s/n", err)
	}
}

func main() {
	utils.InitConf("/Users/yesheng/zinx/demo/config")
	s := znet.NewServe(utils.Interface().Name)

	r := &Router{}
	s.AddRouter(r)
	s.Serve()
}
