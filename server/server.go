package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsConnections []*websocket.Conn

func WSEndPoint(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	wsConnections = append(wsConnections, ws)
	go reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		msg := string(p)
		broadcast(conn, msg)
	}
}

func broadcast(conn *websocket.Conn, msg string) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println(err)
	}
}
