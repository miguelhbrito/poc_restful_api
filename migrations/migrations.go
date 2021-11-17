package migrations

import (
	"bufio"
	"database/sql"
	"os"
	"runtime"

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

	var m *migrate.Migrate
	if runtime.GOOS == "windows" {
		fmt.Print("Windows OS detected, please enter project path:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		m, err = migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://"+text+"stone_assignment/migrations/"),
			"postgres", driver)

	} else {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		m, err = migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", pwd),
			"postgres", driver)
	}

	m.Down()
	m.Up()

	fmt.Println("Successfully migrations applied!")
}
