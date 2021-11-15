package entity

import (
	"time"

	"github.com/stone_assignment/pkg/api/response"
)

type Transfer struct {
	Id              string    `json:"id"`
	AccountOriginId string    `json:"accountOriginId"`
	AccountDestId   string    `json:"accountDestId`
	Ammount         float64   `json:"ammount`
	CreatedAt       time.Time `json:"CreatedAt"`
}

type Transfers []Transfer

func (t Transfer) Response() response.Transfer {
	return response.Transfer{
		Id:              t.Id,
		AccountOriginId: t.AccountOriginId,
		AccountDestId:   t.AccountDestId,
		Ammount:         t.Ammount,
		CreatedAt:       t.CreatedAt,
	}
}

func (trs Transfers) Response() []response.Transfer {
	resp := make([]response.Transfer, 0)
	for i := range trs {
		resp = append(resp, trs[i].Response())
	}
	return resp
}
