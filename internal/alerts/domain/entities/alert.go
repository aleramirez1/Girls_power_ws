package entities

import "time"

type Alert struct {
	ID        int       `json:"id"`        
	UsuarioID int       `json:"usuario_id"` 
	Status    string    `json:"estatus"`    
	CreatedAt time.Time `json:"created_at"` 
}