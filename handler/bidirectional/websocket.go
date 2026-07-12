package bidirectional

import (
	"log"

	"golang.org/x/net/websocket"
)

type Hub struct {
	// connection pool contains all of the connected users.
	clients map[string]*websocket.Conn

	// buffered channels handle user operation queues.
	connectQ    chan Client
	disconnectQ chan Client
	broadcastQ  chan Message
}

func NewHub() *Hub {
	return &Hub{
		clients:     make(map[string]*websocket.Conn),
		connectQ:    make(chan Client, 25),
		disconnectQ: make(chan Client, 25),
		broadcastQ:  make(chan Message, 100),
	}
}

// Run is a long-lived method that runs on the entire server's lifespan.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.connectQ:
			h.clients[client.Username] = client.Conn
			log.Printf("User joined your ws. [Population: %d]\n", len(h.clients))
		case client := <-h.disconnectQ:
			client.Conn.Close()
			delete(h.clients, client.Username)
			log.Printf("User left your ws. [Population: %d]\n", len(h.clients))
		case message := <-h.broadcastQ:
			sendMessage(h, message)
		}
	}
}

// Connect inserts client into the websocket population.
func (h *Hub) Connect(client Client) {
	if client.Username == "" || client.Conn == nil {
		return
	}
	if _, exists := h.clients[client.Username]; exists {
		return
	}
	h.connectQ <- client
}

// Disconnect removes client from websocket population.
func (h *Hub) Disconnect(client Client) {
	if client.Username == "" || client.Conn == nil {
		return
	}
	if _, exists := h.clients[client.Username]; !exists {
		return
	}
	h.disconnectQ <- client
}

// Broadcast sends message to the specified audiences.
func (h *Hub) Broadcast(message Message) {
	if len(message.Audiences) < 1 {
		return
	}
	h.broadcastQ <- message
}
