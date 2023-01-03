package znet

import (
	"fmt"
	"log"
	"net"
	"zinx/ziface"
)

type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的IP版本
	IPVersion string
	// IP 地址
	Addr string
	// 端口号
	Port int
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
	log.Println("start Zinx server succeed", s.Name, "server listening...")
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Printf("accpetTCP error:%s", err)
			continue
		}

		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					log.Printf("recv buf error :%s\n", err)
					return
				}

				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Printf("write back buf error :%s\n", err)
				}
			}
		}()
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

func NewServe(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		Addr:      "0.0.0.0",
		Port:      8081,
	}
	return s
}
