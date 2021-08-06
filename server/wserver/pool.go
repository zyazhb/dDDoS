package wserver

import "sync"

type WsPool struct {
	Lock *sync.RWMutex
	Connections []Conn
}

func NewPool() WsPool {
	return WsPool{
		Lock: new(sync.RWMutex),
	}
}

func (wsp *WsPool) InsertPool(connection Conn) {
	wsp.Lock.Lock()

	wsp.Connections = append(wsp.Connections, connection)

	wsp.Lock.Unlock()
}

func (wsp *WsPool) DeleteFromPool(connection Conn) {
	wsp.Lock.Lock()

	for i, v := range wsp.Connections {
		if v.Connection == connection.Connection {
			wsp.Connections = append(wsp.Connections[:i], wsp.Connections[i+1:]...)
		}
	}

	wsp.Lock.Unlock()
}
