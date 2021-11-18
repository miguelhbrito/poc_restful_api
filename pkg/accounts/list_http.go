package accounts

import (
	"net/http"

	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/merrors"
	"github.com/stone_assignment/pkg/mhttp"
	"github.com/stone_assignment/pkg/mlog"
)

type (
	ListAccountsHTPP struct {
		accountsManager Account
	}
)

func NewListAccountsHTPP(
	accountsManager Account,
) mhttp.HttpHandler {
	return ListAccountsHTPP{
		accountsManager: accountsManager,
	}
}

func (h ListAccountsHTPP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mctx := mcontext.NewFrom(r.Context())
		mlog.Debug(mctx).Msg("receive request to list all accounts")

		acs, err := h.accountsManager.List(mctx)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to list all accounts")
			merrors.Handler(mctx, w, 500, err)
			return
		}

		if err := mhttp.WriteJsonResponse(w, acs.Response(), http.StatusOK); err != nil {
			merrors.Handler(mctx, w, http.StatusOK, err)
			return
		}
	}
}
