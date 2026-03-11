package infrastructure

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"ws-server/internal/alerts/application"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type AlertHandler struct {
	processAlertUC *application.ProcessAlertUseCase
}

func NewAlertHandler(uc *application.ProcessAlertUseCase) *AlertHandler {
	return &AlertHandler{processAlertUC: uc}
}

func (h *AlertHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	var req struct { UsuarioID int `json:"usuario_id"` }
	if err := conn.ReadJSON(&req); err != nil {
		return
	}

	err = h.processAlertUC.Execute(req.UsuarioID)
	if err != nil {
		conn.WriteJSON(map[string]string{"error": "Failed to process"})
		return
	}

	for {
		var incomingMsg map[string]interface{}
		if err := conn.ReadJSON(&incomingMsg); err != nil {
			break 
		}
	}
}