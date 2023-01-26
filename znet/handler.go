package znet

import (
	"fmt"
	"zinx/ziface"

	"github.com/pkg/errors"
)

type Handler struct {
	router map[uint32]ziface.IRouter
}

func NewHandler() ziface.IHandler {
	return &Handler{
		router: make(map[uint32]ziface.IRouter),
	}
}

func (h *Handler) DoMessage(msgID uint32, r ziface.IRequest) error {
	handle, ok := h.router[msgID]
	if ok {
		return errors.WithStack(errors.New(fmt.Sprintf("msgID:%d not exist.", msgID)))
	}
	handle.PerHandle(r)
	handle.Handle(r)
	handle.PostHandle(r)
	return nil
}

func (h *Handler) AddRoute(msgID uint32, r ziface.IRouter) error {
	if _, ok := h.router[msgID]; ok {
		return errors.WithStack(errors.New(fmt.Sprintf("msgID:%d exist", msgID)))
	}
	h.router[msgID] = r
	return nil
}
