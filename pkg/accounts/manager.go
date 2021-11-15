package accounts

import (
	"errors"

	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/storage"
)

type Manager struct {
	accountStorage storage.AccountPostgres
}

func (m Manager) Create(mctx mcontext.Context, ac entity.Account) error {
	newPassword, err := auth.GenerateHashPassword(ac.Secret)
	if err != nil {
		return errors.New("error in password hash")
	}
	ac.Secret = newPassword
	ac.Balance = initialAmmount
	return m.accountStorage.SaveAccount(mctx, ac)
}

func (m Manager) GetById(mctx mcontext.Context, id string) (entity.Account, error) {
	account, err := m.accountStorage.GetByIdAccount(mctx, id)
	if err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

func (m Manager) List(mctx mcontext.Context) (entity.Accounts, error) {
	accounts, err := m.accountStorage.ListAccount(mctx)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (m Manager) Delete(mctx mcontext.Context, id string) error {
	return m.accountStorage.DeleteAccount(mctx, id)
}

func (m Manager) Update(mctx mcontext.Context, ac entity.Account) error {
	return m.accountStorage.UpdateAccount(mctx, ac)
}
