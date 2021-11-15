package login

import (
	"encoding/json"
	"net/http"

	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	mctx := mcontext.NewFrom(r.Context())
	mlog.Debug(mctx).Msg("receive request to login into system")

	var req request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		mlog.Error(mctx).Msgf("Error to decode from json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	loginEntity := req.GenerateEntity()
	loginManager := Manager{}
	token, err := loginManager.LoginIntoSystem(mctx, loginEntity)
	if err != nil {
		mlog.Error(mctx).Msgf("Error to login into system err %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
	return
}
