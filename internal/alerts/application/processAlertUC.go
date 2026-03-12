package application

import (
	domain "ws-server/internal/alerts/domain/repository"
	"ws-server/internal/alerts/domain/entities"
)

type ProcessAlertUseCase struct {
	wsHub domain.WSNotifier
}

func NewProcessAlertUseCase(ws domain.WSNotifier) *ProcessAlertUseCase {
	return &ProcessAlertUseCase{wsHub: ws}
}

func (uc *ProcessAlertUseCase) Execute(senderID int, payload entities.AlertPayload) {
	ids := make([]int, 0, len(payload.Users))
	for _, u := range payload.Users {
		ids = append(ids, u.UsuarioID)
	}
	uc.wsHub.NotifyMultiple(ids, "NEARBY_ALERT", payload)
}