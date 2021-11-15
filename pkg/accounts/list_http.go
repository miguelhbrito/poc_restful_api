package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

func ListAccountsHandler(w http.ResponseWriter, r *http.Request) {
	mctx := mcontext.NewFrom(r.Context())
	mlog.Debug(mctx).Msg("receive request to list all accounts")

	accountsManager := Manager{}
	acs, err := accountsManager.List(mctx)
	if err != nil {
		mlog.Error(mctx).Msgf("Error to list all accounts")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acs.Response())
	return
}
