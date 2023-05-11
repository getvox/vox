package server

import "sync"

type ConnManager struct {
	sync.Mutex
	conns map[string]*Connection
}

func NewConnManager() *ConnManager {
	cm := new(ConnManager)
	cm.conns = make(map[string]*Connection)
	return cm
}

func (cm *ConnManager) Add(c *Connection) {
	cm.Lock()
	cm.conns[c.ID] = c
	cm.Unlock()
}

func (cm *ConnManager) Get(id string) *Connection {
	cm.Lock()
	defer cm.Unlock()
	c := cm.conns[id]
	return c
}

func (cm *ConnManager) Remove(c *Connection) {
	cm.Lock()
	delete(cm.conns, c.ID)
	cm.Unlock()
}
