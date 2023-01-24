package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"

	"github.com/pkg/errors"
	"go.uber.org/zap"
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
	logger := utils.Logger
	logger.Debug("[Start] Server Listening...", zap.String("IP", s.Addr), zap.Int("port", s.Port))
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Addr, s.Port))
	if err != nil {
		logger.Error("resolve tcp addr", zap.Error(errors.WithStack(err)))
		return
	}

	listen, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		logger.Error("listen tcp", zap.Error(errors.WithStack(err)))
		return
	}
	var cid uint32 = 1
	logger.Debug(fmt.Sprintf("start Zinx server succeed  %s server listening...", s.Name))
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			logger.Error("acceptTCP", zap.Error(errors.WithStack(err)))
			continue
		}

		delConn := NewConnection(conn, cid, s.Router)
		go delConn.Start()
	}
}

func (s *Server) Stop() {
	// TODO 将一些服务器资源停止，
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
		Addr:      utils.Interface().Host,
		Port:      utils.Interface().Port,
	}
	return s
}
