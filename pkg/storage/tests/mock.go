package tests

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

type AccountCustomMock struct {
	SaveAccountMock          func(mctx mcontext.Context, ac entity.Account) error
	GetByIdAccountMock       func(mctx mcontext.Context, id string) (entity.Account, error)
	ListAccountMock          func(mctx mcontext.Context) ([]entity.Account, error)
	DeleteAccountMock        func(mctx mcontext.Context, id string) error
	UpdateBalanceAccountMock func(mctx mcontext.Context, balance float64) error
}

func (a AccountCustomMock) SaveAccount(mctx mcontext.Context, ac entity.Account) error {
	return a.SaveAccountMock(mctx, ac)
}

func (a AccountCustomMock) GetByIdAccount(mctx mcontext.Context, id string) (entity.Account, error) {
	return a.GetByIdAccountMock(mctx, id)
}

func (a AccountCustomMock) ListAccount(mctx mcontext.Context) ([]entity.Account, error) {
	return a.ListAccountMock(mctx)
}

func (a AccountCustomMock) DeleteAccount(mctx mcontext.Context, id string) error {
	return a.DeleteAccountMock(mctx, id)
}

func (a AccountCustomMock) UpdateBalanceAccount(mctx mcontext.Context, balance float64) error {
	return a.UpdateBalanceAccountMock(mctx, balance)
}
