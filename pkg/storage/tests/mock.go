package tests

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

type AccountCustomMock struct {
	SaveAccountMock     func(mctx mcontext.Context, ac entity.Account) error
	GetByIdAccountMock  func(mctx mcontext.Context, id string) (entity.Account, error)
	GetByCpfAccountMock func(mctx mcontext.Context, cpf string) (entity.Account, error)
	ListAccountMock     func(mctx mcontext.Context) ([]entity.Account, error)
	DeleteAccountMock   func(mctx mcontext.Context, id string) error
	UpdateAccountMock   func(mctx mcontext.Context, ac entity.Account) error
	GetCredentialsMock  func(mctx mcontext.Context, cpf string) (entity.LoginEntity, error)
}

type TransferCustomMock struct {
	SaveTransferMock  func(mctx mcontext.Context, tr entity.Transfer) error
	ListTransfersMock func(mctx mcontext.Context) ([]entity.Transfer, error)
}

func (a AccountCustomMock) SaveAccount(mctx mcontext.Context, ac entity.Account) error {
	return a.SaveAccountMock(mctx, ac)
}

func (a AccountCustomMock) GetByIdAccount(mctx mcontext.Context, id string) (entity.Account, error) {
	return a.GetByIdAccountMock(mctx, id)
}

func (a AccountCustomMock) GetByCpfAccount(mctx mcontext.Context, cpf string) (entity.Account, error) {
	return a.GetByCpfAccountMock(mctx, cpf)
}

func (a AccountCustomMock) ListAccount(mctx mcontext.Context) ([]entity.Account, error) {
	return a.ListAccountMock(mctx)
}

func (a AccountCustomMock) DeleteAccount(mctx mcontext.Context, id string) error {
	return a.DeleteAccountMock(mctx, id)
}

func (a AccountCustomMock) UpdateAccount(mctx mcontext.Context, ac entity.Account) error {
	return a.UpdateAccountMock(mctx, ac)
}

func (a AccountCustomMock) GetCredentials(mctx mcontext.Context, cpf string) (entity.LoginEntity, error) {
	return a.GetCredentialsMock(mctx, cpf)
}

func (t TransferCustomMock) SaveTransfer(mctx mcontext.Context, tr entity.Transfer) error {
	return t.SaveTransferMock(mctx, tr)
}

func (t TransferCustomMock) ListTransfers(mctx mcontext.Context) ([]entity.Transfer, error) {
	return t.ListTransfersMock(mctx)
}
