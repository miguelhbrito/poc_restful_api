package accounts

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

type AccountCustomMock struct {
	CreateMock  func(mctx mcontext.Context, ac entity.Account) (entity.Account, error)
	GetByIdMock func(mctx mcontext.Context, id string) (entity.Account, error)
	ListMock    func(mctx mcontext.Context) (entity.Accounts, error)
	DeleteMock  func(mctx mcontext.Context, id string) error
	UpdateMock  func(mctx mcontext.Context, ac entity.Account) error
}

func (a AccountCustomMock) Create(mctx mcontext.Context, ac entity.Account) (entity.Account, error) {
	return a.CreateMock(mctx, ac)
}

func (a AccountCustomMock) GetById(mctx mcontext.Context, id string) (entity.Account, error) {
	return a.GetByIdMock(mctx, id)
}

func (a AccountCustomMock) List(mctx mcontext.Context) (entity.Accounts, error) {
	return a.ListMock(mctx)
}

func (a AccountCustomMock) Delete(mctx mcontext.Context, id string) error {
	return a.DeleteMock(mctx, id)
}

func (a AccountCustomMock) Update(mctx mcontext.Context, ac entity.Account) error {
	return a.UpdateMock(mctx, ac)
}
