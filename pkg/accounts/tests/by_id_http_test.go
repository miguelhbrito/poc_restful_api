package accounts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	db "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/migrations"
	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/api/response"
)

func createAccount() response.Account {
	reqAccount := request.CreateAccount{
		Name:     "any_name",
		Cpf:      "96483478593",
		Password: "password",
	}

	body, _ := json.Marshal(reqAccount)

	reqCreate, err := http.NewRequest("POST", "/accounts", bytes.NewReader(body))
	if err != nil {
		fmt.Errorf(err.Error())
	}

	rrCreate := httptest.NewRecorder()
	handlerCreate := http.HandlerFunc(accounts.CreateAccountsHandler)

	handlerCreate.ServeHTTP(rrCreate, reqCreate)

	var resp response.Account
	err = json.NewDecoder(rrCreate.Body).Decode(&resp)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	return resp
}

func TestGetByIdAccountsHandlerWhenEverythingIsOkThenSuccess(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	//Add an account into system
	resp := createAccount()

	req, err := http.NewRequest("GET", fmt.Sprintf("/accounts/%s/balance", resp.Id), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accounts.GetByIdAccountsHandler)

	vars := map[string]string{
		"account_id": resp.Id,
	}

	req = mux.SetURLVars(req, vars)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var respCheck response.Account
	err = json.NewDecoder(rr.Body).Decode(&respCheck)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body is what we expect.
	expected := `{` + resp.Id + ` any_name 96483478593 100 ` + respCheck.CreatedAt + `}`

	if fmt.Sprint(respCheck) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			fmt.Sprint(respCheck), expected)
	}
}

func TestGetByIdAccountsHandlerWhenNowFoundAccountThenFail(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	req, err := http.NewRequest("GET", "/accounts/"+"123"+"/balance", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accounts.GetByIdAccountsHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{}\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
