package znet

import (
	"github.com/leaf-gentlemen/zinx/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
		connLock:    sync.RWMutex{},
	}
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	defer c.connLock.Unlock()
	c.connLock.Lock()
	if _, ok := c.connections[conn.GetConnID()]; ok {
		return
	}
	c.connections[conn.GetConnID()] = conn
}

func (c *ConnManager) Remove(connID uint32) {
	defer c.connLock.Unlock()
	c.connLock.Lock()
	delete(c.connections, connID)
}

func (c *ConnManager) Get(connID uint32) (ziface.IConnection, bool) {
	defer c.connLock.RUnlock()
	c.connLock.RLock()
	conn, ok := c.connections[connID]
	return conn, ok
}

func (c *ConnManager) Len() int {
	defer c.connLock.RUnlock()
	c.connLock.RLock()
	return len(c.connections)
}

func (c *ConnManager) Clear() {
	defer c.connLock.Unlock()
	c.connLock.Lock()
	for _, conn := range c.connections {
		conn.Stop()
		delete(c.connections, conn.GetConnID())
	}
}
