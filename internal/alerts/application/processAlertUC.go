package application

import(
	domain "ws-server/internal/alerts/domain/repository"
)
type ProcessAlertUseCase struct {
	alertRepo domain.AlertRepository
	wsHub     domain.WSNotifier
}

func NewProcessAlertUseCase(repo domain.AlertRepository, ws domain.WSNotifier) *ProcessAlertUseCase {
	return &ProcessAlertUseCase{
		alertRepo: repo,
		wsHub:     ws,
	}
}

func (uc *ProcessAlertUseCase) Execute(requestingUserID int) error {
	res, err := uc.alertRepo.RegisterDanger(requestingUserID)
	if err != nil {
		return err 
	}

	uc.wsHub.NotifyMultiple(res.NearbyUsers, "NEARBY_ALERT", res.Alert)

	for _, contact := range res.EmergencyContacts {
		if contact.UsuarioID != nil { 
			uc.wsHub.NotifyUser(*contact.UsuarioID, "FAMILY_ALERT", res.Alert)
		}
	}
	
	return nil
}