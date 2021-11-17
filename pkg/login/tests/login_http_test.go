package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/migrations"
	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/api/response"
	"github.com/stone_assignment/pkg/login"
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

func TestLoginHandlerWhenEverythingIsOkThenSuccess(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	_ = createAccount()

	reqLogin := request.LoginRequest{
		Cpf:    "96483478593",
		Secret: "password",
	}

	body, _ := json.Marshal(reqLogin)

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login.LoginHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var resp response.LoginToken
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}
	// Check the response body is what we expect.
	expected := `{` + resp.Token + " " + fmt.Sprint(resp.ExpTime) + `}`

	if fmt.Sprint(resp) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			fmt.Sprint(resp), expected)
	}
}

func TestLoginHandlerWhenInvalidPasswordThenFail(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	_ = createAccount()

	reqLogin := request.LoginRequest{
		Cpf:    "96483478593",
		Secret: "not a password",
	}

	body, _ := json.Marshal(reqLogin)

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login.LoginHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 401 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := `{"message":"Username or Password is incorrect"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestLoginHandlerWhenUserNotFoundThenFail(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	reqLogin := request.LoginRequest{
		Cpf:    "96483478593",
		Secret: "password",
	}

	body, _ := json.Marshal(reqLogin)

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login.LoginHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 401 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := `{"message":"sql: no rows in result set"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
