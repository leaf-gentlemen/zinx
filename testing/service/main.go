package main

import (
	"zinx/testing/service/internal/control"
	"zinx/utils"
	"zinx/znet"
)

func main() {
	utils.InitConf("/Users/yesheng/zinx/testing/config")
	s := znet.NewServe(utils.Interface().Name)
	s.AddRouter(0, &control.Login{})
	s.AddRouter(1, &control.User{})
	s.Serve()
}
