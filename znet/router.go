package znet

import "zinx/ziface"

type BaseRouter struct{}

func (r *BaseRouter) PerHandle(request ziface.IRequest) {}

func (r *BaseRouter) Handle(request ziface.IRequest) {}

func (r *BaseRouter) PostHandle(request ziface.IRequest) {}
