package login

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

func TestLoginHTPP_Handler(t *testing.T) {
	tests := []struct {
		name    string
		manager Login
		h       LoginHTPP
		body    []byte
		request request.LoginRequest
		want    http.HandlerFunc
	}{
		{
			name: "Success",
			manager: LoginCustomMock{
				LoginIntoSystemMock: func(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error) {
					return response.LoginToken{
						Token:   "any_token",
						ExpTime: time.Time{}.Unix(),
					}, nil
				},
			},
			request: request.LoginRequest{
				Cpf:    "12345678910",
				Secret: "any_secret",
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(response.LoginToken{
					Token:   "any_token",
					ExpTime: time.Time{}.Unix(),
				})
			},
		},
		{
			name: "Error to decode json",
			manager: LoginCustomMock{
				LoginIntoSystemMock: func(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error) {
					return response.LoginToken{}, nil
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
			name: "Error to login into system",
			manager: LoginCustomMock{
				LoginIntoSystemMock: func(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error) {
					return response.LoginToken{}, errors.New("some error")
				},
			},
			request: request.LoginRequest{
				Cpf:    "12345678910",
				Secret: "any_secret",
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				data, _ := json.Marshal(merrors.HTTPError{Msg: "some error"})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error on validate fields",
			manager: LoginCustomMock{
				LoginIntoSystemMock: func(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error) {
					return response.LoginToken{}, nil
				},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				data, _ := json.Marshal(merrors.HTTPError{Msg: "cpf is required,secret is required"})
				_, _ = w.Write(data)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewLoginHTPP(tt.manager)

			body, _ := json.Marshal(tt.request)
			if tt.body != nil {
				body = tt.body
			}
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))

			w := httptest.NewRecorder()

			tt.want.ServeHTTP(w, req)

			g := httptest.NewRecorder()

			h.Handler()(g, req)

			assert.Equal(t, w.Code, g.Result().StatusCode, fmt.Sprintf("expected status code %v ", w.Code))

			assert.Equal(t, w.Body.String(), g.Body.String(), "body was not equal as expected")
		})
	}
}
