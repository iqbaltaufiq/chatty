package handler

import "golang.org/x/net/websocket"

type WebsocketHandler interface {
	Connect(conn *websocket.Conn)
}
