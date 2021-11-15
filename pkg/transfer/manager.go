package transfer

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/storage"
)

type Manager struct {
	transferStorage storage.TransferPostgres
	accountStorage  storage.AccountPostgres
}

func (m Manager) Create(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
	accountOrigin, err := m.accountStorage.GetByCpfAccount(mctx, mctx.Username().String())
	if err != nil {
		return entity.Transfer{}, err
	}
	tr.AccountOriginId = accountOrigin.Id
	newOriginBalance, err := checkOriginAmmount(accountOrigin.Balance, tr.Ammount)
	if err != nil {
		return entity.Transfer{}, err
	}
	err = m.accountStorage.UpdateAccount(mctx, entity.Account{Id: accountOrigin.Id, Balance: newOriginBalance})
	if err != nil {
		return entity.Transfer{}, err
	}

	accountDest, err := m.accountStorage.GetByIdAccount(mctx, tr.AccountDestId)
	if err != nil {
		return entity.Transfer{}, err
	}
	err = m.accountStorage.UpdateAccount(mctx, entity.Account{Id: accountDest.Id, Balance: accountDest.Balance + tr.Ammount})
	if err != nil {
		return entity.Transfer{}, err
	}

	err = m.transferStorage.SaveTransfer(mctx, tr)
	if err != nil {
		return entity.Transfer{}, err
	}

	return tr, nil
}

func (m Manager) List(mctx mcontext.Context) (entity.Transfers, error) {
	transfers, err := m.transferStorage.ListTransfers(mctx)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

func checkOriginAmmount(originBalance, ammount float64) (float64, error) {
	if originBalance <= 0 {
		return originBalance, errBalanceLowerThan0
	}
	resultBalance := originBalance - ammount
	if resultBalance <= 0 {
		return originBalance, errBalanceCantAffort
	}
	return resultBalance, nil
}
