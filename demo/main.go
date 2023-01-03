package main

import (
	"zinx/znet"
)

func main() {
	s := znet.NewServe("game server")
	s.Serve()
}
