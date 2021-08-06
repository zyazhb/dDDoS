package wserver

import (
	"errors"
	"io"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

/*
	onclose | onmessage | onopen | onerror
*/

type Conn struct {
	Connection *websocket.Conn

	AfterReadFunc   func(messageType int, r io.Reader)
	BeforeCloseFunc func()

	stopChan chan struct{}
}

func (c *Conn) Write(text []byte) error {
	select {
	case <-c.stopChan:
		return errors.New("[x] connection is closed, can't be written")
	default:
		err := c.Connection.WriteMessage(websocket.TextMessage, text)
		if err != nil {
			return err
		}

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
			messageType, r, err := c.Connection.NextReader()
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
