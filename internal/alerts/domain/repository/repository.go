package repository

import "ws-server/internal/alerts/domain/entities"

type AlertRepository interface {
	RegisterDanger(usuarioID int) (*entities.AlertResponse, error)
}

type WSNotifier interface {
	NotifyUser(usuarioID int, event string, payload interface{})
	NotifyMultiple(usuarioIDs []int, event string, payload interface{})
}