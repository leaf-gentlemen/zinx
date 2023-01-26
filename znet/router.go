package znet

import (
	"fmt"
	"zinx/utils"
	"zinx/ziface"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

type Router struct {
	router map[uint32]ziface.IHandler
}

func NewRouter() ziface.IRouter {
	return &Router{
		router: make(map[uint32]ziface.IHandler),
	}
}

func (r *Router) DoMessage(req ziface.IRequest) error {
	handle, ok := r.router[req.GetMessage().GetMsgID()]
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("msgID:%d not exist.", req.GetMessage().GetMsgID())))
	}
	handle.PerHandle(req)
	handle.Handle(req)
	handle.PostHandle(req)
	return nil
}

func (h *Router) AddRoute(msgID uint32, r ziface.IHandler) error {
	logger := utils.Logger
	if _, ok := h.router[msgID]; ok {
		return errors.WithStack(errors.New(fmt.Sprintf("msgID:%d exist", msgID)))
	}
	h.router[msgID] = r
	logger.Debug(fmt.Sprintf("add router msgID: %d", msgID), zap.Any("control", r))
	return nil
}
