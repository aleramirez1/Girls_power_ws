package infrastructure

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ConnectedClient struct {
	Conn *websocket.Conn
}

type Hub struct {
	mu      sync.RWMutex
	Clients map[int]*ConnectedClient
}

func NewHub() *Hub {
	return &Hub{Clients: make(map[int]*ConnectedClient)}
}

func (h *Hub) Register(userID int, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Clients[userID] = &ConnectedClient{Conn: conn}
	log.Printf("Usuario %d conectado\n", userID)
}

func (h *Hub) Unregister(userID int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.Clients, userID)
	log.Printf("Usuario %d desconectado\n", userID)
}

func (h *Hub) NotifyUser(usuarioID int, event string, payload interface{}) {
	h.mu.RLock()
	client, exists := h.Clients[usuarioID]
	h.mu.RUnlock()

	if !exists {
		return
	}
	msg := map[string]interface{}{
		"event":   event,
		"payload": payload,
	}
	if err := client.Conn.WriteJSON(msg); err != nil {
		log.Printf("Error enviando a usuario %d: %v\n", usuarioID, err)
	}
}

func (h *Hub) NotifyMultiple(usuarioIDs []int, event string, payload interface{}) {
	for _, id := range usuarioIDs {
		h.NotifyUser(id, event, payload)
	}
}