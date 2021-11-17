package main

import (
	"fmt"
	"net/http"

	db "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/migrations"
	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/login"
	"github.com/stone_assignment/pkg/middleware"
	"github.com/stone_assignment/pkg/storage"
	"github.com/stone_assignment/pkg/transfers"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	//Starting rounter
	router := mux.NewRouter()
	addrHttp := fmt.Sprintf(":%d", 3000)

	//Auth manager
	authManager := auth.NewManager()

	//Starting accounts postgres and manager
	accountPostgres := storage.NewAccountPostgres()
	accountManager := accounts.NewManager(accountPostgres, authManager)

	//Starting login manager
	loginManager := login.NewManager(accountPostgres, authManager)

	//Starting transfer postgres and manager
	transferPostgres := storage.NewTransferPostgres()
	transferManager := transfers.NewManager(transferPostgres, accountManager)

	//HANDLERS ----------------

	//Starting accounts handlers
	accountCreate := accounts.NewCreateAccountHTPP(accountManager)
	accountList := accounts.NewListAccountsHTPP(accountManager)
	accountById := accounts.NewByIdAccountHTPP(accountManager)

	//Starting login handlers
	loginCreate := login.NewLoginHTPP(loginManager)

	//Starting login handlers
	transferCreate := transfers.NewCreateTransferHTPP(transferManager)
	transferList := transfers.NewListTransferHTPP(transferManager)

	//ENDPOINTS----------------

	//Accounts endpoints
	router.HandleFunc("/accounts", middleware.Authorization(accountCreate.Handler())).Methods("POST")
	router.HandleFunc("/accounts", middleware.Authorization(accountList.Handler())).Methods("GET")
	router.HandleFunc("/accounts/{account_id}/balance", middleware.Authorization(accountById.Handler())).Methods("GET")

	//Login endpoints
	router.HandleFunc("/login", middleware.Authorization(loginCreate.Handler())).Methods("POST")

	//Transfers endpoints
	router.HandleFunc("/transfers", middleware.Authorization(transferCreate.Handler())).Methods("POST")
	router.HandleFunc("/transfers", middleware.Authorization(transferList.Handler())).Methods("GET")

	//Starting gateway
	log.Fatal().Err(http.ListenAndServe(addrHttp, router)).Msg("failed to start gateway")
}
