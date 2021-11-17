package transfers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/api/request"
	"github.com/stone_assignment/pkg/api/response"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/merrors"
	"github.com/stone_assignment/pkg/transfers"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransferHTPP_Handler(t *testing.T) {
	tests := []struct {
		name    string
		manager transfers.Transfer
		h       transfers.CreateTransferHTPP
		body    []byte
		request request.TransferRequest
		want    http.HandlerFunc
	}{
		{
			name: "Success",
			manager: TransferCustomMock{
				CreateMock: func(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
					return entity.Transfer{
						Id:        "any_id",
						CreatedAt: time.Time{},
					}, nil
				},
			},
			request: request.TransferRequest{
				AccountDestId: "any_id",
				Ammount:       10.50,
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				_ = json.NewEncoder(w).Encode(response.Transfer{
					Id:        "any_id",
					CreatedAt: time.Time{},
				})
			},
		},
		{
			name: "Error to decode json",
			manager: TransferCustomMock{
				CreateMock: func(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
					return entity.Transfer{}, nil
				},
			},
			body: []byte(""),
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(merrors.HTTPError{Msg: "Error to decode from json, err:EOF"})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error to validate fields",
			manager: TransferCustomMock{
				CreateMock: func(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
					return entity.Transfer{}, nil
				},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				data, _ := json.Marshal(merrors.HTTPError{Msg: "account destination id is required,ammount to be transfer need to be greater than 0"})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error on save tranfers into database",
			manager: TransferCustomMock{
				CreateMock: func(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
					return entity.Transfer{}, errors.New("some error")
				},
			},
			request: request.TransferRequest{
				AccountDestId: "any_id",
				Ammount:       10.50,
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(merrors.HTTPError{Msg: errors.New("some error").Error()})
				_, _ = w.Write(data)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := transfers.NewCreateTransferHTPP(tt.manager)

			body, _ := json.Marshal(tt.request)
			if tt.body != nil {
				body = tt.body
			}
			req, _ := http.NewRequest(http.MethodPost, "/transfers", bytes.NewReader(body))

			w := httptest.NewRecorder()

			tt.want.ServeHTTP(w, req)

			g := httptest.NewRecorder()

			h.Handler()(g, req)

			assert.Equal(t, w.Code, g.Result().StatusCode, fmt.Sprintf("expected status code %v ", w.Code))

			assert.Equal(t, w.Body.String(), g.Body.String(), "body was not equal as expected")
		})
	}
}
