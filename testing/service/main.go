package main

import (
	"fmt"
	"github.com/leaf-gentlemen/zinx/testing/service/internal/control"
	"github.com/leaf-gentlemen/zinx/utils"
	"github.com/leaf-gentlemen/zinx/ziface"
	"github.com/leaf-gentlemen/zinx/znet"
)

// 创建连接的时候执行
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
	conn.SetProperty("name", "test")
}

// 连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {
	name, _ := conn.GetProperty("name")
	fmt.Printf("DoConneciotnLost is Called ... name:%s \n", name)
}

func main() {
	utils.InitConf("testing/config")
	s := znet.NewServe(utils.Interface().Name)
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	_ = s.AddRouter(0, &control.Login{})
	_ = s.AddRouter(1, &control.User{})
	s.Serve()
}
