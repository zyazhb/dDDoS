package wserver

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

/*
	onclose | onmessage | onopen | onerror
*/

type Conn struct {
	Connection *websocket.Conn

	AfterReadFunc   func(messageType int, r []byte)
	BeforeCloseFunc func()

	stopChan chan struct{}
}

func (c *Conn) Write(text []byte) error {
	select {
	case <-c.stopChan:
		return errors.New("[x] connection is closed, can't be written")
	default:
		w, err := c.Connection.NextWriter(websocket.TextMessage)
		if err != nil {
			return err
		}

		w.Write(text)

		return nil
	}
}

func (c *Conn) Close() error {
	select {
	case <-c.stopChan:
		return errors.New("[+] connection has been closed")
	default:
		c.Connection.Close()
		close(c.stopChan)

		return nil
	}
}

func (c *Conn) Listen() {
	// TODO: 能用，但是连接关闭控制还要想想能不能实现连接池进行管理
	c.Connection.SetCloseHandler(func(code int, text string) error {
		if c.BeforeCloseFunc != nil {
			c.BeforeCloseFunc()
		}

		if err := c.Close(); err != nil {
			log.Fatalln(err)
			return err
		}

		message := websocket.FormatCloseMessage(code, "Close the conn")
		c.Connection.WriteControl(code, message, time.Now().Add(time.Second))
		return nil
	})

Exit:
	for {
		select {
		case <-c.stopChan:
			break Exit
		default:
			messageType, r, err := c.Connection.ReadMessage()
			if err != nil {
				log.Fatalln("[x] Read next message error!")
				break
			}

			if c.AfterReadFunc != nil {
				// process get message
				c.AfterReadFunc(messageType, r)
			}
		}
	}
}

func NewConn(conn *websocket.Conn) *Conn {
	return &Conn{
		Connection: conn,
		stopChan:   make(chan struct{}),
	}
}
