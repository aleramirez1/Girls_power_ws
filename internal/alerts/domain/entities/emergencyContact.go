package entities

type EmergencyContact struct {
	PersonaID int    `json:"persona_id"`
	UsuarioID *int   `json:"usuario_id"` 
	Name      string `json:"nombre"` 
}