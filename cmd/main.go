package main

import (
	"fmt"
	"net/http"

	db "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/migrations"
	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/login"
	"github.com/stone_assignment/pkg/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {

	dbconnection := db.InitDB()
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	router := mux.NewRouter()

	addrHttp := fmt.Sprintf(":%d", 3000)

	router.HandleFunc("/accounts", middleware.Authorization(accounts.CreateAccountsHandler)).Methods("POST")
	router.HandleFunc("/accounts", middleware.Authorization(accounts.ListAccountsHandler)).Methods("GET")
	router.HandleFunc("/accounts/{account_id}/balance", middleware.Authorization(accounts.ListAccountsHandler)).Methods("GET")

	router.HandleFunc("/login", middleware.Authorization(login.LoginHandler)).Methods("POST")

	log.Fatal().Err(http.ListenAndServe(addrHttp, router)).Msg("failed to start gateway")
}
