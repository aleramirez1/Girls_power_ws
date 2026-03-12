package entities

type AlertPayload struct {
	Users []struct {
		UsuarioID int `json:"usuario_id"`
	} `json:"users"`
}