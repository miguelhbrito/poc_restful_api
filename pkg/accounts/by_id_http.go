package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/merrors"
	"github.com/stone_assignment/pkg/mhttp"
	"github.com/stone_assignment/pkg/mlog"
)

type (
	ByIdAccountHTPP struct {
		accountsManager Account
	}
)

func NewByIdAccountHTPP(
	accountsManager Account,
) mhttp.HttpHandler {
	return ByIdAccountHTPP{
		accountsManager: accountsManager,
	}
}

func (h ByIdAccountHTPP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mctx := mcontext.NewFrom(r.Context())
		mlog.Debug(mctx).Msg("receive request to get balance by account id")

		params := mux.Vars(r)
		id := params["account_id"]

		ac, err := h.accountsManager.GetById(mctx, id)
		if err != nil {
			mlog.Error(mctx).Err(err).Msgf("Error to get balance account {%s}", id)
			merrors.Handler(mctx, w, 500, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ac.Response())
		return
	}
}
