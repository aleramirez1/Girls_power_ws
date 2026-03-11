package main

import (
	"log"
	"net/http"
	"ws-server/internal/alerts/application"
	"ws-server/internal/alerts/infrastructure"
)

func main() {

	apiClient := &infrastructure.RESTClient{BaseURL: "http://localhost:8080/api/v1"}
	wsHub := &infrastructure.Hub{Clients: make(map[int]*infrastructure.ConnectedClient)}

	processAlertUC := application.NewProcessAlertUseCase(apiClient, wsHub)

	alertHandler := infrastructure.NewAlertHandler(processAlertUC)
	
	mux := http.NewServeMux()
	infrastructure.RegisterRoutes(mux, alertHandler)

	log.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}