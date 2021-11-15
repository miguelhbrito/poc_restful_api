package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

func CreateAccountsHandler(w http.ResponseWriter, r *http.Request) {
	mctx := mcontext.NewFrom(r.Context())
	mlog.Debug(mctx).Msg("receive request to create account")

	var req request.CreateAccount
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		mlog.Error(mctx).Msgf("Error to decode from json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accountEntity := req.GenerateEntity()
	accountsManager := Manager{}
	err = accountsManager.Create(mctx, accountEntity)
	if err != nil {
		mlog.Error(mctx).Msgf("Error to create new account")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(accountEntity.Response())

	return
}
