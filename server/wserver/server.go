package wserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Server struct {
	Connection *Conn
	Upgrader   *websocket.Upgrader
}

func (s *Server) WebsocketHandler(c *gin.Context) {
	s.Upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := s.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalln("[x] Upgrade to websocket connection error!")
		return
	}
	s.Connection = NewConn(conn)

	s.Connection.BeforeCloseFunc = func() {
		s.Connection.Close()
	}

	s.Connection.AfterReadFunc = func(messageType int, r []byte) {
		log.Printf("[+] Receive message: %s\n", string(r))
	}

	// s.Connection.AfterReadFunc = nil

	s.Connection.Listen()
}

func NewServer() Server {
	return Server{}
}

