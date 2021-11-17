package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/merrors"
	"github.com/stone_assignment/pkg/mhttp"
	"github.com/stone_assignment/pkg/mlog"
)

type (
	LoginHTPP struct {
		loginManager Login
	}
)

func NewLoginHTPP(
	loginManager Login,
) mhttp.HttpHandler {
	return LoginHTPP{
		loginManager: loginManager,
	}
}

func (h LoginHTPP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mctx := mcontext.NewFrom(r.Context())
		mlog.Debug(mctx).Msg("receive request to login into system")

		var req request.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to decode from json")
			merrors.Handler(mctx, w, 500, errors.New(
				fmt.Sprintf("Error to decode from json, err:%s", err.Error())))
			return
		}

		err = req.Validate()
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to validate fields from login")
			merrors.Handler(mctx, w, 400, err)
			return
		}

		loginEntity := req.GenerateEntity()
		token, err := h.loginManager.LoginIntoSystem(mctx, loginEntity)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to login into system")
			merrors.Handler(mctx, w, 401, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(token)
		return
	}
}
