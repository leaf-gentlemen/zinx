package znet

import "github.com/leaf-gentlemen/zinx/ziface"

type BaseHandler struct{}

func (r *BaseHandler) PerHandle(request ziface.IRequest) {}

func (r *BaseHandler) Handle(request ziface.IRequest) {}

func (r *BaseHandler) PostHandle(request ziface.IRequest) {}
