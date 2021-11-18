package request

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/stone_assignment/pkg/api/entity"
)

type TransferRequest struct {
	AccountDestId string  `json:"accountDestId"`
	Ammount       float64 `json:"ammount"`
}

func (t TransferRequest) GenerateEntity() entity.Transfer {
	return entity.Transfer{
		Id:            uuid.New().String(),
		AccountDestId: t.AccountDestId,
		Ammount:       t.Ammount,
		CreatedAt:     time.Now(),
	}
}

func (t TransferRequest) Validate() error {
	var errs = ""
	if t.AccountDestId == "" {
		errs += "account destination id is required"
	}
	if t.Ammount <= 0 {
		errs += ",ammount to be transfer need to be greater than 0"
	}
	if len(errs) > 0 {
		return errors.New(errs)
	}
	return nil
}
