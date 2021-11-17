package accounts

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
)

func TestCreateAccountsHandlerWhenEverythingIsOkThenSuccess(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	reqAccount := request.CreateAccount{
		Name:     "any_name",
		Cpf:      "96483478593",
		Password: "password",
	}

	body, _ := json.Marshal(reqAccount)

	req, err := http.NewRequest("POST", "/accounts", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accounts.CreateAccountsHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 201 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var resp response.Account
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body is what we expect.
	expected := `{` + resp.Id + ` any_name 96483478593 100 ` + resp.CreatedAt + `}`

	if fmt.Sprint(resp) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			fmt.Sprint(resp), expected)
	}
}

func TestCreateAccountsHandlerWhenInvalidCpfThenFail(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	reqAccount := request.CreateAccount{
		Name:     "any_name",
		Cpf:      "12345678910",
		Password: "password",
	}

	body, _ := json.Marshal(reqAccount)

	req, err := http.NewRequest("POST", "/accounts", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accounts.CreateAccountsHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 500 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := `{"message":"Cpf(12345678910) is invalid"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateAccountsHandlerWhenMissingInputsThenFail(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	reqAccount := request.CreateAccount{
		Name:     "",
		Cpf:      "",
		Password: "",
	}

	body, _ := json.Marshal(reqAccount)

	req, err := http.NewRequest("POST", "/accounts", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accounts.CreateAccountsHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 400 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	expected := `{"message":"name is required,cpf is required,,password can not be nil"}{"message":"Cpf must not be empty"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateAccountsHandlerWhenMissingBodyThenFail(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	body, _ := json.Marshal(".")

	req, err := http.NewRequest("POST", "/accounts", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(accounts.CreateAccountsHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 500 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := `{"message":"json: cannot unmarshal string into Go value of type request.CreateAccount"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


