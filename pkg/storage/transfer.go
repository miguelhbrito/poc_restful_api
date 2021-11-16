package storage

import (
	dbconnect "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

type TransferPostgres struct{}

func (t TransferPostgres) SaveTransfer(mctx mcontext.Context, tr entity.Transfer) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `INSERT INTO transfer_account VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, tr.Id, tr.AccountOriginId, tr.AccountDestId, tr.Ammount, tr.CreatedAt)
	if err != nil {
		mlog.Error(mctx).Err(err).Msg("Error to insert an new transfer into db")
		return err
	}
	return nil
}

func (t TransferPostgres) ListTransfers(mctx mcontext.Context) ([]entity.Transfer, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var trs []entity.Transfer
	sqlStatement := `SELECT id, account_origin_id, account_destination_id, amount, created_at FROM transfer_account`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		mlog.Error(mctx).Err(err).Msg("Error to get all transfers from db")
		return nil, err
	}
	for rows.Next() {
		var tr entity.Transfer
		err := rows.Scan(&tr.Id, &tr.AccountOriginId, &tr.AccountDestId, &tr.Ammount, &tr.CreatedAt)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to extract result from row")
		}
		trs = append(trs, tr)
	}
	return trs, nil
}
