package user

import "time"

type ResUsuario struct {
	ID          *int64     `json:"id"`
	Nome        *string    `json:"nome"`
	Email       *string    `json:"email"`
	DataCriacao *time.Time `json:"data_criacao"`
}
