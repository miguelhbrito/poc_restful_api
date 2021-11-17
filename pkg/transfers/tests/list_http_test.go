package transfers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/stone_assignment/db_connect"
	"github.com/stone_assignment/migrations"
	"github.com/stone_assignment/pkg/transfers"
)

func TestListTransfersHandlerWhenEverythingIsOkThenSuccess(t *testing.T) {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	_ = createAccount("96483478593")

	token := createLogin()

	req, err := http.NewRequest("GET", "/transfers", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("authorization", token.Token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transfers.ListTransfersHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body is what we expect.
	expected := "[]\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
