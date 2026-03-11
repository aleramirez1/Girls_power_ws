package entities

type AlertResponse struct {
	Alert             Alert              `json:"alerta"`
	EmergencyContacts []EmergencyContact `json:"contactos_emergencia"` 
	NearbyUsers       []int              `json:"usuarios_cercanos"` 
}