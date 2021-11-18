package accounts

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
	"github.com/stretchr/testify/assert"
)

func Test_createAccountHTPP_Handler(t *testing.T) {
	tests := []struct {
		manager Account
		name    string
		h       CreateAccountHTPP
		body    []byte
		request request.CreateAccount
		want    http.HandlerFunc
	}{
		{
			name: "Success",
			manager: AccountCustomMock{CreateMock: func(mctx mcontext.Context, ac entity.Account) (entity.Account, error) {
				return entity.Account{
					Id:        "any_id",
					Name:      "any_name",
					Cpf:       "96483478593",
					Secret:    "any_secret",
					Balance:   1,
					CreatedAt: time.Time{},
				}, nil
			},
			},
			request: request.CreateAccount{
				Name:     "any_name",
				Cpf:      "96483478593",
				Password: "any_secret",
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				_ = json.NewEncoder(w).Encode(response.Account{
					Id:        "any_id",
					Name:      "any_name",
					Cpf:       "96483478593",
					Balance:   1,
					CreatedAt: time.Time{}.String(),
				})
			},
		},
		{
			name: "Error to create a new account",
			manager: AccountCustomMock{CreateMock: func(mctx mcontext.Context, ac entity.Account) (entity.Account, error) {
				return entity.Account{}, errors.New("some error")
			},
			},
			request: request.CreateAccount{
				Name:     "any_name",
				Cpf:      "96483478593",
				Password: "any_secret",
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(merrors.HTTPError{Msg: errors.New("some error").Error()})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error to validate fields",
			manager: AccountCustomMock{CreateMock: func(mctx mcontext.Context, ac entity.Account) (entity.Account, error) {
				return entity.Account{}, nil
			},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				data, _ := json.Marshal(merrors.HTTPError{Msg: "name is required,cpf is required,,password can not be nil"})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error to decode json",
			manager: AccountCustomMock{CreateMock: func(mctx mcontext.Context, ac entity.Account) (entity.Account, error) {
				return entity.Account{}, nil
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewCreateAccountHTPP(tt.manager)

			body, _ := json.Marshal(tt.request)
			if tt.body != nil {
				body = tt.body
			}
			req, _ := http.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))

			w := httptest.NewRecorder()

			tt.want.ServeHTTP(w, req)

			g := httptest.NewRecorder()

			h.Handler()(g, req)

			assert.Equal(t, w.Code, g.Result().StatusCode, fmt.Sprintf("expected status code %v ", w.Code))

			assert.Equal(t, w.Body.String(), g.Body.String(), "body was not equal as expected")
		})
	}
}
