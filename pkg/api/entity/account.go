package entity

import (
	"time"

	"github.com/stone_assignment/pkg/api/response"
)

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type Accounts []Account

func (ac Account) Response() response.Account {
	return response.Account{
		Id:        ac.Id,
		Name:      ac.Name,
		Cpf:       ac.Cpf,
		Balance:   ac.Balance,
		CreatedAt: ac.CreatedAt.String(),
	}
}

func (ac Accounts) Response() []response.Account {
	resp := make([]response.Account, 0)
	for i := range ac {
		resp = append(resp, ac[i].Response())
	}
	return resp
}
