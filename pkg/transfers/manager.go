package transfers

import (
	"errors"

	dbconnect "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
	"github.com/stone_assignment/pkg/storage"
)

type Manager struct {
	transferStorage storage.TransferPostgres
	accountStorage  storage.AccountPostgres
}

func (m Manager) Create(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
	mlog.Debug(mctx).Msgf("Starting transfers action between two accounts!")

	accountOrigin, err := m.accountStorage.GetByCpfAccount(mctx, mctx.Cpf().String())
	if err != nil {
		return entity.Transfer{}, err
	}
	tr.AccountOriginId = accountOrigin.Id

	accountDest, err := m.accountStorage.GetByIdAccount(mctx, tr.AccountDestId)
	if err != nil {
		return entity.Transfer{}, err
	}

	err = m.transferBetweenTwoAccounts(mctx, accountOrigin, accountDest, tr)
	return tr, nil
}

func (m Manager) List(mctx mcontext.Context) (entity.Transfers, error) {
	transfers, err := m.transferStorage.ListTransfers(mctx)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

func (m Manager) transferBetweenTwoAccounts(mctx mcontext.Context, origin, destiny entity.Account, tr entity.Transfer) error {
	newOriginBalance, err := checkOriginAmmount(origin.Balance, tr.Ammount)
	if err != nil {
		return err
	}

	db := dbconnect.InitDB()
	defer db.Close()
	tx, err := db.BeginTx(mctx, nil)
	txc := mcontext.WithValue(mctx, "tx", tx)

	origin.Balance = newOriginBalance
	err = m.accountStorage.UpdateAccount(txc, origin)
	if err != nil {
		return err
	}

	destiny.Balance = destiny.Balance + tr.Ammount
	err = m.accountStorage.UpdateAccount(txc, destiny)
	if err != nil {
		return err
	}

	err = m.transferStorage.SaveTransfer(txc, tr)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		mlog.Error(mctx).Err(err).Msg("Error to commit transfer action")
		return errors.New("error commit")
	}

	return nil
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
