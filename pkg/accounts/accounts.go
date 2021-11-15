package accounts

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

type AccountManager interface {
	Create(mctx mcontext.Context, ac entity.Account) error
	GetById(mctx mcontext.Context, id string) (entity.Account, error)
	List(mctx mcontext.Context) (entity.Accounts, error)
	Delete(mctx mcontext.Context, id string) error
	Update(mctx mcontext.Context, ac entity.Account) error
}
