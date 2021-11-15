package storage

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

type Account interface {
	SaveAccount(mctx mcontext.Context, ac entity.Account) error
	GetByIdAccount(mctx mcontext.Context, id string) (entity.Account, error)
	ListAccount(mctx mcontext.Context) ([]entity.Account, error)
	DeleteAccount(mctx mcontext.Context, id string) error
	UpdateAccount(mctx mcontext.Context, ac entity.Account) error
}

type Login interface {
	SaveLogin(mctx mcontext.Context, loginEntity entity.LoginEntity) error
}
