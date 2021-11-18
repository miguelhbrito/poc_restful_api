package transfers

import (
	"net/http"

	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/merrors"
	"github.com/stone_assignment/pkg/mhttp"
	"github.com/stone_assignment/pkg/mlog"
)

type (
	ListTransferHTPP struct {
		transferManager Transfer
	}
)

func NewListTransferHTPP(
	transferManager Transfer,
) mhttp.HttpHandler {
	return ListTransferHTPP{
		transferManager: transferManager,
	}
}

func (h ListTransferHTPP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mctx := mcontext.NewFrom(r.Context())
		mlog.Debug(mctx).Msg("receive request to list all transfers")

		trs, err := h.transferManager.List(mctx)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("Error to list all transfers")
			merrors.Handler(mctx, w, 500, err)
			return
		}

		if err := mhttp.WriteJsonResponse(w, trs.Response(), http.StatusOK); err != nil {
			merrors.Handler(mctx, w, http.StatusOK, err)
			return
		}
	}
}
