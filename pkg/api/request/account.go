package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/stone_assignment/pkg/api/entity"
)

type CreateAccount struct {
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Password string `json:"password"`
}

func (c CreateAccount) GenerateEntity() entity.Account {
	return entity.Account{
		Id:        uuid.New().String(),
		Name:      c.Name,
		Cpf:       c.Cpf,
		Secret:    c.Password,
		CreatedAt: time.Now(),
	}
}
