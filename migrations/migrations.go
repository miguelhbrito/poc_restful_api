package migrations

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"fmt"
)

func InitMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	//linux users :
	/*pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", pwd),
		"postgres", driver)
	*/

	//only for windows
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://C:/Users/migue/Documents/dev/github/stone_assignment/migrations/"),
		"postgres", driver)

	m.Down()
	m.Up()
}
