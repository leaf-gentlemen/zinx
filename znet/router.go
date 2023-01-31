package znet

import (
	"fmt"
	"zinx/utils"
	"zinx/ziface"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Router struct {
	router         map[uint32]ziface.IHandler
	taskQueue      []chan ziface.IRequest
	workerPoolSize uint32
}

func NewRouter() ziface.IRouter {
	return &Router{
		router:         make(map[uint32]ziface.IHandler),
		taskQueue:      make([]chan ziface.IRequest, utils.Interface().WorkerPoolSize),
		workerPoolSize: utils.Interface().WorkerPoolSize,
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

func (r *Router) AddRoute(msgID uint32, h ziface.IHandler) error {
	logger := utils.Logger
	if _, ok := r.router[msgID]; ok {
		return errors.WithStack(errors.New(fmt.Sprintf("msgID:%d exist", msgID)))
	}
	r.router[msgID] = h
	logger.Debug(fmt.Sprintf("add router msgID: %d", msgID), zap.Any("control", r))
	return nil
}

func (r *Router) StartWorkerPool() {
	for i := 0; i < int(r.workerPoolSize); i++ {
		r.taskQueue[i] = make(chan ziface.IRequest, utils.Interface().WorkerMsgLen)
		go r.startOneWorker(i, r.taskQueue[i])
	}
}

func (r *Router) startOneWorker(workerID int, reqChan chan ziface.IRequest) {
	logger := utils.Logger
	logger.Debug("start worker", zap.Int("workerID", workerID))
	for req := range reqChan {
		if err := r.DoMessage(req); err != nil {
			logger.Error("do message fail", zap.Error(err))
		}
	}
}

func (r *Router) SendMsgToTaskQueue(req ziface.IRequest) {
	logger := utils.Logger
	workerID := req.GetConnection().GetConnID() % r.workerPoolSize
	r.taskQueue[workerID] <- req
	logger.Debug("send msg to task queue", zap.Uint32("workerID", workerID), zap.Int("taskMsgLen", len(r.taskQueue[workerID])))
}
