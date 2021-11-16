package storage

import (
	"database/sql"

	dbconnect "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

type AccountPostgres struct{}

func (a AccountPostgres) SaveAccount(mctx mcontext.Context, ac entity.Account) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `INSERT INTO account VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, ac.Id, ac.Name, ac.Cpf, ac.Secret, ac.Balance, ac.CreatedAt)
	if err != nil {
		mlog.Error(mctx).Err(err).Msgf("Error to insert an new account into db")
		return err
	}
	return nil
}

func (a AccountPostgres) GetByIdAccount(mctx mcontext.Context, id string) (entity.Account, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var ac entity.Account
	sqlStatement := `SELECT id, name, cpf, secret_login, balance, created_at FROM account WHERE id = $1`
	result := db.QueryRow(sqlStatement, id)
	err := result.Scan(&ac.Id, &ac.Name, &ac.Cpf, &ac.Secret, &ac.Balance, &ac.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			mlog.Error(mctx).Err(err).Msgf("Not found account with id %s", id)
			return entity.Account{}, err
		}
		mlog.Error(mctx).Err(err).Msgf("Error to get account from db, with id %s", id)
		return entity.Account{}, err
	}
	return ac, nil
}

func (a AccountPostgres) GetByCpfAccount(mctx mcontext.Context, cpf string) (entity.Account, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var ac entity.Account
	sqlStatement := `SELECT id, name, cpf, secret_login, balance, created_at FROM account WHERE cpf = $1`
	result := db.QueryRow(sqlStatement, cpf)
	err := result.Scan(&ac.Id, &ac.Name, &ac.Cpf, &ac.Secret, &ac.Balance, &ac.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			mlog.Error(mctx).Err(err).Msgf("Not found account with cpf %s", cpf)
			return entity.Account{}, err
		}
		mlog.Error(mctx).Err(err).Msgf("Error to get account from db, with cpf %s", cpf)
		return entity.Account{}, err
	}
	return ac, nil
}

func (a AccountPostgres) ListAccount(mctx mcontext.Context) ([]entity.Account, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var acs []entity.Account
	sqlStatement := `SELECT id, name, cpf, secret_login, balance, created_at FROM account`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		mlog.Error(mctx).Err(err).Msg("Error to get all accounts from db")
		return nil, err
	}
	for rows.Next() {
		var ac entity.Account
		err := rows.Scan(&ac.Id, &ac.Name, &ac.Cpf, &ac.Secret, &ac.Balance, &ac.CreatedAt)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to extract result from row")
		}
		acs = append(acs, ac)
	}
	return acs, nil
}

func (a AccountPostgres) DeleteAccount(mctx mcontext.Context, id string) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `DELETE FROM notebook WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		mlog.Error(mctx).Err(err).Msg("Error to delete account from db")
		return err
	}
	return nil
}

func (a AccountPostgres) UpdateAccount(mctx mcontext.Context, ac entity.Account) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `UPDATE account SET balance=$2 WHERE id=$1`
	_, err := db.Exec(sqlStatement, ac.Id, ac.Balance)
	if err != nil {
		mlog.Error(mctx).Err(err).Msg("Error to update account from db")
		return err
	}
	return nil
}

func (a AccountPostgres) GetCredentials(mctx mcontext.Context, cpf string) (entity.LoginEntity, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var lr entity.LoginEntity
	sqlStatement := `SELECT cpf, secret_login FROM account WHERE cpf = $1`
	result := db.QueryRow(sqlStatement, cpf)
	err := result.Scan(&lr.Cpf, &lr.Secret)
	if err != nil {
		if err == sql.ErrNoRows {
			mlog.Error(mctx).Err(err).Msgf("Not found account with cpf %s", cpf)
			return entity.LoginEntity{}, err
		}
		mlog.Error(mctx).Err(err).Msgf("Error to get credentials from db, with id %s", cpf)
		return entity.LoginEntity{}, err
	}
	return lr, nil
}
