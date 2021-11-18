package transfers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/api/response"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/merrors"
	"github.com/stretchr/testify/assert"
)

func TestListTransferHTPP_Handler(t *testing.T) {
	tests := []struct {
		name    string
		manager Transfer
		h       ListTransferHTPP
		want    http.HandlerFunc
	}{
		{
			name: "Success",
			manager: TransferCustomMock{
				ListMock: func(mctx mcontext.Context) (entity.Transfers, error) {
					return entity.Transfers{{
						Id:        "any_id",
						CreatedAt: time.Time{},
					}}, nil
				},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode([]response.Transfer{{
					Id:        "any_id",
					CreatedAt: time.Time{},
				}})
			},
		},
		{
			name: "Error on list transfers from database",
			manager: TransferCustomMock{
				ListMock: func(mctx mcontext.Context) (entity.Transfers, error) {
					return nil, errors.New("some error")
				},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(merrors.HTTPError{Msg: "some error"})
				_, _ = w.Write(data)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewListTransferHTPP(tt.manager)

			req, _ := http.NewRequest(http.MethodGet, "/transfers", nil)

			w := httptest.NewRecorder()

			tt.want.ServeHTTP(w, req)

			g := httptest.NewRecorder()

			h.Handler()(g, req)

			assert.Equal(t, w.Code, g.Result().StatusCode, fmt.Sprintf("expected status code %v ", w.Code))

			assert.Equal(t, w.Body.String(), g.Body.String(), "body was not equal as expected")
		})
	}
}
