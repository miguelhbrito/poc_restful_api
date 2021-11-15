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
	UpdateBalanceAccount(mctx mcontext.Context, balance float64) error
}

type Transfer interface {
	SaveTransfer(mctx mcontext.Context, ac entity.Account) error
	ListTransfers(mctx mcontext.Context) ([]entity.Account, error)
}
