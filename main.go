package main

import (
	"log"
	"net/http"
	"os"

	"ws-server/internal/alerts/application"
	"ws-server/internal/alerts/infrastructure"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")

	hub := infrastructure.NewHub()
	processAlertUC := application.NewProcessAlertUseCase(hub)
	alertHandler := infrastructure.NewAlertHandler(processAlertUC, hub, jwtSecret)

	mux := http.NewServeMux()
	infrastructure.RegisterRoutes(mux, alertHandler)

	log.Println("Server running on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
