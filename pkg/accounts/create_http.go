package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/merrors"
	"github.com/stone_assignment/pkg/mhttp"
	"github.com/stone_assignment/pkg/mlog"
)

type (
	CreateAccountHTPP struct {
		accountManager Account
	}
)

func NewCreateAccountHTPP(
	accountsManager Account,
) mhttp.HttpHandler {
	return CreateAccountHTPP{
		accountManager: accountsManager,
	}
}

func (h CreateAccountHTPP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mctx := mcontext.NewFrom(r.Context())
		mlog.Debug(mctx).Msg("receive request to create account")

		var req request.CreateAccount
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to decode from json")
			merrors.Handler(mctx, w, 500,
				fmt.Errorf("Error to decode from json, err:%s", err.Error()))
			return
		}

		err = req.Validate()
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to validate fields from account")
			merrors.Handler(mctx, w, 400, err)
			return
		}

		accountEntity := req.GenerateEntity()
		accountResult, err := h.accountManager.Create(mctx, accountEntity)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to create new account")
			merrors.Handler(mctx, w, 500, err)
			return
		}

		if err := mhttp.WriteJsonResponse(w, accountResult.Response(), http.StatusCreated); err != nil {
			merrors.Handler(mctx, w, http.StatusCreated, err)
			return
		}
	}
}
