package accounts

import (
	brdocs "github.com/brazanation/go-documents"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
	"github.com/stone_assignment/pkg/storage"
)

type Manager struct {
	accountStorage storage.Account
	auth           auth.Auth
}

func NewManager(accountStorage storage.Account, auth auth.Auth) Account {
	return Manager{
		accountStorage: accountStorage,
		auth:           auth,
	}
}

func (m Manager) Create(mctx mcontext.Context, ac entity.Account) (entity.Account, error) {
	doc, err := brdocs.NewCpf(ac.Cpf)
	if err != nil {
		return entity.Account{}, err
	}
	mlog.Debug(mctx).Msgf("Checked validation of cpf {%s}", doc.Format())

	newPassword, err := m.auth.GenerateHashPassword(ac.Secret)
	if err != nil {
		return entity.Account{}, errPasswordHash
	}

	ac.Secret = newPassword
	ac.Balance = initialAmmount
	err = m.accountStorage.SaveAccount(mctx, ac)
	if err != nil {
		return entity.Account{}, err
	}
	return ac, nil
}

func (m Manager) GetById(mctx mcontext.Context, id string) (entity.Account, error) {
	account, err := m.accountStorage.GetByIdAccount(mctx, id)
	if err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

func (m Manager) GetByCpf(mctx mcontext.Context, cpf string) (entity.Account, error) {
	account, err := m.accountStorage.GetByCpfAccount(mctx, cpf)
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
