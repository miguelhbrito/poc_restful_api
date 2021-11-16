package request

import (
	"errors"

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

func (l LoginRequest) Validate() error {
	var errs = ""
	if l.Cpf == "" {
		errs += "cpf is required"
	}
	if l.Secret == "" {
		errs += ",secret is required"
	}
	if len(errs) > 0 {
		return errors.New(errs)
	}
	return nil
}
