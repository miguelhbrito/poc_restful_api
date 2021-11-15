package request

import (
	"github.com/stone_assignment/pkg/api/entity"
)

type LoginRequest struct {
	Cpf      string `json:"cpf"`
	Password string `json:"password`
}

func (l LoginRequest) GenerateEntity() entity.LoginEntity {
	return entity.LoginEntity{
		Cpf:    l.Cpf,
		Secret: l.Password,
	}
}
