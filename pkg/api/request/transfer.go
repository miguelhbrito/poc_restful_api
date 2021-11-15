package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/stone_assignment/pkg/api/entity"
)

type Transfer struct {
	AccountDestId string  `json:"accountDestId`
	Ammount       float64 `json:"ammount`
}

func (t Transfer) GenerateEntity() entity.Transfer {
	return entity.Transfer{
		Id:            uuid.New().String(),
		AccountDestId: t.AccountDestId,
		Ammount:       t.Ammount,
		CreatedAt:     time.Now(),
	}
}
