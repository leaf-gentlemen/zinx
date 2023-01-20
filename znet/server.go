package znet

import (
	"fmt"
	"log"
	"net"
	"zinx/ziface"
)

type Server struct {
	//
	//  Name
	//  @Description: 服务器名称
	//
	Name string

	//
	//  IPVersion
	//  @Description: 服务器绑定的IP版本
	//
	IPVersion string
	//
	//  Addr
	//  @Description: 地址
	//
	Addr string
	//
	//  Port
	//  @Description: 端口
	//
	Port int
	//
	//  Router
	//  @Description: 路由
	//
	Router ziface.IRouter
}

func (s *Server) Start() {
	log.Printf("[Start] Server Listenner at IP :%s, Port %d, is starting\n", s.Addr, s.Port)

	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Addr, s.Port))
	if err != nil {
		log.Printf("reslove tcp addr error:%s\n", err)
		return
	}

	listen, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		log.Printf("listen tcp error:%s\n", err)
		return
	}
	var cid uint32 = 1
	log.Println("start Zinx server succeed", s.Name, "server listening...")
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Printf("accpetTCP error:%s", err)
			continue
		}

		delConn := NewConnection(conn, cid, s.Router)
		go delConn.Start()
	}
}

func (s *Server) Stop() {
	//TODO 将一些服务器资源停止，
}

func (s *Server) Serve() {
	s.Start()

	// 阻塞
	select {}
}

func (s *Server) AddRouter(r ziface.IRouter) {
	// TODO 添加多个路由
	s.Router = r
}

func NewServe(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		Addr:      "0.0.0.0",
		Port:      8081,
	}
	return s
}
