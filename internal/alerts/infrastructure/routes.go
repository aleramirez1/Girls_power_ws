package infrastructure

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler *AlertHandler) {
	mux.HandleFunc("/ws/alerts", handler.HandleWebSocket)
}