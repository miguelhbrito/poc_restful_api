package accounts

import (
	"errors"

	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

const (
	initialAmmount = 100.00
)

var (
	errPasswordHash = errors.New("Error to generate password hash")
)

type Account interface {
	Create(mctx mcontext.Context, ac entity.Account) (entity.Account, error)
	GetById(mctx mcontext.Context, id string) (entity.Account, error)
	GetByCpf(mctx mcontext.Context, cpf string) (entity.Account, error)
	List(mctx mcontext.Context) (entity.Accounts, error)
	Delete(mctx mcontext.Context, id string) error
	Update(mctx mcontext.Context, ac entity.Account) error
}
