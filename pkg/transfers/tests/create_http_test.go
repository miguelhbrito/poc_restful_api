package transfers

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
	"github.com/stone_assignment/pkg/transfers"
)

func createAccount(cpf string) response.Account {
	reqAccount := request.CreateAccount{
		Name:     "any_name",
		Cpf:      cpf,
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

func createLogin() response.LoginToken {
	reqLogin := request.LoginRequest{
		Cpf:    "96483478593",
		Secret: "password",
	}

	body, _ := json.Marshal(reqLogin)

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		fmt.Errorf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login.LoginHandler)

	handler.ServeHTTP(rr, req)

	var resp response.LoginToken
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return resp
}

func TestCreateTransfersHandlerWhenEverythingIsOkThenSuccess(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	_ = createAccount("96483478593")
	resp := createAccount("58847923603")

	token := createLogin()

	reqTransfer := request.TransferRequest{
		AccountDestId: resp.Id,
		Ammount:       20.50,
	}

	body, _ := json.Marshal(reqTransfer)

	req, err := http.NewRequest("POST", "/transfers", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("authorization", token.Token)

	fmt.Println(req)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transfers.CreateTransfersHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 201 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	fmt.Println(rr.Body.String())
	var respTransfer response.Transfer
	err = json.NewDecoder(rr.Body).Decode(&respTransfer)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response body is what we expect.
	expected := `{}`

	if fmt.Sprint(respTransfer) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			fmt.Sprint(respTransfer), expected)
	}
}
