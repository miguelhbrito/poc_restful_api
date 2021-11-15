package transfer

import (
	"encoding/json"
	"net/http"

	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

func CreateTransfersHandler(w http.ResponseWriter, r *http.Request) {
	mctx := mcontext.NewFrom(r.Context())
	mlog.Debug(mctx).Msg("receive request to create a transfer")

	var req request.Transfer
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		mlog.Error(mctx).Msgf("Error to decode from json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	transferEntity := req.GenerateEntity()
	transferManager := Manager{}
	transfer, err := transferManager.Create(mctx, transferEntity)
	if err != nil {
		mlog.Error(mctx).Msgf("Error to create new transfer, err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transfer.Response())

	return
}
