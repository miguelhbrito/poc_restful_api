package request

import (
	"github.com/stone_assignment/pkg/api/entity"
)

type LoginRequest struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret`
}

func (l LoginRequest) GenerateEntity() entity.LoginEntity {
	return entity.LoginEntity{
		Cpf:    l.Cpf,
		Secret: l.Secret,
	}
}
