package transfers

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
	CreateTransferHTPP struct {
		transferManager Transfer
	}
)

func NewCreateTransferHTPP(
	transferManager Transfer,
) mhttp.HttpHandler {
	return CreateTransferHTPP{
		transferManager: transferManager,
	}
}

func (h CreateTransferHTPP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mctx := mcontext.NewFrom(r.Context())
		mlog.Debug(mctx).Msg("receive request to create a transfer")

		var req request.TransferRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to decode from json")
			merrors.Handler(mctx, w, 500, errors.New(
				fmt.Sprintf("Error to decode from json, err:%s", err.Error())))
			return
		}

		err = req.Validate()
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to validate fields from transfer")
			merrors.Handler(mctx, w, 400, err)
			return
		}

		transferEntity := req.GenerateEntity()
		transfer, err := h.transferManager.Create(mctx, transferEntity)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to create new transfer")
			merrors.Handler(mctx, w, 500, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(transfer.Response())

		return
	}
}
