package storage

import (
	dbconnect "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

type LoginPostgres struct{}

func (l LoginPostgres) SaveLogin(mctx mcontext.Context, le entity.LoginEntity) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `INSERT INTO login_user VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, le.Cpf, le.Secret)
	if err != nil {
		mlog.Error(mctx).Err(err).Msgf("Error to insert an new login into db %v", err)
		return err
	}
	return nil
}
