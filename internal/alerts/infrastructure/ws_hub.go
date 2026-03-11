package infrastructure

import (
	"log"
	"github.com/gorilla/websocket"
)

type ConnectedClient struct {
	Conn *websocket.Conn
}

type Hub struct {
	Clients map[int]*ConnectedClient
}

func (h *Hub) NotifyUser(usuarioID int, event string, payload interface{}) {
	if client, exists := h.Clients[usuarioID]; exists {
		
		msg := map[string]interface{}{
			"event":   event,
			"payload": payload,
		}

		err := client.Conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error sending message to user %d: %v\n", usuarioID, err)
		}
	}
}

func (h *Hub) NotifyMultiple(usuarioIDs []int, event string, payload interface{}) {
	for _, id := range usuarioIDs {
		h.NotifyUser(id, event, payload)
	}
}