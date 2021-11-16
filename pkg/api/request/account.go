package request

import (
	"errors"
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

func (c CreateAccount) Validate() error {
	var errs = ""
	if c.Name == "" {
		errs += "name is required"
	}
	if c.Cpf == "" {
		errs += ",cpf is required,"
	}
	if c.Password == "" {
		errs += ",password can not be nil"
	}
	if len(errs) > 0 {
		return errors.New(errs)
	}
	return nil
}
