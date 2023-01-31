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
	// Name 服务器名称
	Name string
	// IPVersion 服务器绑定的IP版本
	IPVersion string
	// Addr 地址
	Addr string
	// Port 端口
	Port int
	// Router 路由
	Router ziface.IRouter
	// ConnManager 连接管理器
	ConnManager ziface.IConnManager
	// onConnStop 连接关闭前 hook 函数
	onConnStop func(c ziface.IConnection)
	// onConnStart  创建连接前 hook 函数
	onConnStart func(c ziface.IConnection)
}

func NewServe(name string) ziface.IServer {
	s := &Server{
		Name:        name,
		IPVersion:   "tcp4",
		Addr:        utils.Interface().Host,
		Port:        utils.Interface().Port,
		Router:      NewRouter(),
		ConnManager: NewConnManager(),
	}
	return s
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
	s.Router.StartWorkerPool()
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			logger.Error("acceptTCP", zap.Error(errors.WithStack(err)))
			continue
		}

		if s.GetConnManager().Len() >= utils.Interface().MaxConn {
			logger.Warn("Connection exceeds online", zap.Int("maxConn", utils.Interface().MaxConn))
			continue
		}

		delConn := NewConnection(s, conn, cid, s.Router)
		go delConn.Start()
	}
}

func (s *Server) SetOnConnStart(hookFun func(ziface.IConnection)) {
	s.onConnStart = hookFun
}

func (s *Server) SetOnConnStop(hookFun func(ziface.IConnection)) {
	s.onConnStop = hookFun
}

func (s *Server) GetConnManager() ziface.IConnManager {
	return s.ConnManager
}

func (s *Server) Stop() {
	// TODO 将一些服务器资源停止，
	s.ConnManager.Clear()
}

func (s *Server) Serve() {
	s.Start()

	// 阻塞
	select {}
}

func (s *Server) AddRouter(msgID uint32, r ziface.IHandler) error {
	return s.Router.AddRoute(msgID, r)
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.onConnStart != nil {
		s.onConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.onConnStop != nil {
		s.onConnStop(conn)
	}
}
