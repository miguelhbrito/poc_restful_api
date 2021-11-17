package transfers

import (
	"errors"

	dbconnect "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
	"github.com/stone_assignment/pkg/storage"
)

type manager struct {
	transferStorage storage.Transfer
	accountManager  accounts.Account
}

func NewManager(transferStorage storage.Transfer, accountManager accounts.Account) Transfer {
	return manager{
		accountManager:  accountManager,
		transferStorage: transferStorage,
	}
}

func (m manager) Create(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
	mlog.Debug(mctx).Msgf("Starting transfers action between two accounts!")

	//Getting account origin
	accountOrigin, err := m.accountManager.GetByCpf(mctx, mctx.Cpf().String())
	if err != nil {
		return entity.Transfer{}, err
	}
	tr.AccountOriginId = accountOrigin.Id

	//Check if it is same account
	if accountOrigin.Id == tr.AccountDestId {
		return entity.Transfer{}, errSameAccount
	}

	//Getting account destination
	accountDest, err := m.accountManager.GetById(mctx, tr.AccountDestId)
	if err != nil {
		return entity.Transfer{}, err
	}

	//Transfer actions
	err = m.transferBetweenTwoAccounts(mctx, accountOrigin, accountDest, tr)
	return tr, nil
}

func (m manager) List(mctx mcontext.Context) (entity.Transfers, error) {
	transfers, err := m.transferStorage.ListTransfers(mctx)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

func (m manager) transferBetweenTwoAccounts(mctx mcontext.Context, origin, destination entity.Account, tr entity.Transfer) error {
	//Check origin ammount
	newOriginBalance, err := checkOriginAmmount(origin.Balance, tr.Ammount)
	if err != nil {
		return err
	}

	mlog.Debug(mctx).Msgf("Begin tx manager!")
	//Begin tx manager to check if all transfers actions on database was done ok, otherwise none of them will be commited
	db := dbconnect.InitDB()
	defer db.Close()
	tx, err := db.BeginTx(mctx, nil)
	txc := mcontext.WithValue(mctx, "tx", tx)

	//Update balance on origin account
	origin.Balance = newOriginBalance
	err = m.accountManager.Update(txc, origin)
	if err != nil {
		return err
	}

	//Update balance on destination account
	destination.Balance = destination.Balance + tr.Ammount
	err = m.accountManager.Update(txc, destination)
	if err != nil {
		return err
	}

	//Save transfer on database
	err = m.transferStorage.SaveTransfer(txc, tr)
	if err != nil {
		return err
	}

	//Then commit all transactions on database
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
