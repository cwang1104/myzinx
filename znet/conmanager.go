package znet

import (
	"errors"
	"fmt"
	"myzinx/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理的连接信息
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	c.connections[conn.GetConnID()] = conn

	fmt.Println("conn add to connmanager successfully conn num = ", c.Len())

}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.connections, conn.GetConnID())
	fmt.Println("conn delete to connmanager successfully conn num = ", c.Len())

}

func (c *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connections not found")
	}
}

// 获取当前连接
func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connId, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connId)
	}
	fmt.Println("Clear All Connections successfully: conn num = ", c.Len())
}
