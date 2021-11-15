package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/stone_assignment/pkg/api/entity"
)

type CreateAccount struct {
	Name    string  `json:"name"`
	Cpf     string  `json:"cpf"`
	Balance float64 `json:"balance"`
}

func (c CreateAccount) GenerateEntity() entity.Account {
	return entity.Account{
		Id:        uuid.New().String(),
		Name:      c.Name,
		Cpf:       c.Cpf,
		Balance:   c.Balance,
		CreatedAt: time.Now(),
	}
}
