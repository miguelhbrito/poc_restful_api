package accounts

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

type ManagerCustomMock struct {
	CreateMock  func(mctx mcontext.Context, ac entity.Account) (entity.Account, error)
	GetByIdMock func(mctx mcontext.Context, id string) (entity.Account, error)
	ListMock    func(mctx mcontext.Context) (entity.Accounts, error)
	DeleteMock  func(mctx mcontext.Context, id string) error
	UpdateMock  func(mctx mcontext.Context, ac entity.Account) error
}

func (a ManagerCustomMock) Create(mctx mcontext.Context, ac entity.Account) (entity.Account, error) {
	return a.CreateMock(mctx, ac)
}

func (a ManagerCustomMock) GetById(mctx mcontext.Context, id string) (entity.Account, error) {
	return a.GetByIdMock(mctx, id)
}

func (a ManagerCustomMock) List(mctx mcontext.Context) (entity.Accounts, error) {
	return a.ListMock(mctx)
}

func (a ManagerCustomMock) Delete(mctx mcontext.Context, id string) error {
	return a.DeleteMock(mctx, id)
}

func (a ManagerCustomMock) Update(mctx mcontext.Context, ac entity.Account) error {
	return a.UpdateMock(mctx, ac)
}
