package entity

import (
	"time"

	"github.com/stone_assignment/pkg/api/response"
)

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"Cpf"`
	Secret    string    `json:"Secret"`
	Balance   float64   `json:"Balance"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type Accounts []Account

func (ac Account) Response() response.Account {
	return response.Account{
		Id:        ac.Id,
		Name:      ac.Name,
		Cpf:       ac.Cpf,
		Balance:   ac.Balance,
		CreatedAt: ac.CreatedAt,
	}
}

func (ac Accounts) Response() []response.Account {
	resp := make([]response.Account, 0)
	for i := range ac {
		resp = append(resp, ac[i].Response())
	}
	return resp
}
