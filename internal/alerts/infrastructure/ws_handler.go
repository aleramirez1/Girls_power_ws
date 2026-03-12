package infrastructure

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"ws-server/internal/alerts/application"
	"ws-server/internal/alerts/domain/entities"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type AlertHandler struct {
	processAlertUC *application.ProcessAlertUseCase
	hub            *Hub
	jwtSecret      string
}

func NewAlertHandler(uc *application.ProcessAlertUseCase, hub *Hub, secret string) *AlertHandler {
	return &AlertHandler{
		processAlertUC: uc,
		hub:            hub,
		jwtSecret:      secret,
	}
}

func (h *AlertHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.URL.Query().Get("token")
	if tokenStr == "" {
		http.Error(w, "token requerido", http.StatusUnauthorized)
		return
	}

	userID, err := ExtractUserIDFromToken(tokenStr, h.jwtSecret)
	if err != nil {
		http.Error(w, "token inválido", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	h.hub.Register(userID, conn)
	defer h.hub.Unregister(userID)

	for {
		var payload entities.AlertPayload
		if err := conn.ReadJSON(&payload); err != nil {
			break
		}

		h.processAlertUC.Execute(userID, payload)
	}
}