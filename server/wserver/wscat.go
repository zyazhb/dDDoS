package wserver

import (
	"log"

	"github.com/gorilla/websocket"
)

type WsCat struct {
	conn *websocket.Conn
}

func (wsc *WsCat) Connect(wspath string) error {
	conn, _, err := websocket.DefaultDialer.Dial(wspath, nil)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	wsc.conn = conn
	return nil
}

func (wsc WsCat) WriteMessage(p string) {
	wsc.conn.WriteMessage(websocket.TextMessage, []byte(p))
}

func (wsc WsCat) WriteJsonMessage(jp []byte) {
	wsc.conn.WriteMessage(websocket.TextMessage, jp)
}
