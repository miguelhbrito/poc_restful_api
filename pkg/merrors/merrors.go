package merrors

import (
	"encoding/json"
	"net/http"

	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/mlog"
)

type HTTPError struct {
	Err    error  `json:"-"`
	Msg    string `json:"message"`
	Status int    `json:"-"`
}

func Handler(mctx mcontext.Context, w http.ResponseWriter, status int, err error) {
	mlog.Debug(mctx).Msg("Error handler")

	he := HTTPError{
		Err:    err,
		Msg:    err.Error(),
		Status: status,
	}

	msg, err := json.Marshal(he)
	w.Header().Set("Content-type", "application/json")
	if status != 0 {
		w.WriteHeader(he.Status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err == nil {
		_, _ = w.Write(msg)
	}
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}
