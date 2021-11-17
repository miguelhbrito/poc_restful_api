package accounts

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/storage"
	tests "github.com/stone_assignment/pkg/storage/tests"
)

func TestManager_Create(t *testing.T) {
	type args struct {
		mctx mcontext.Context
		ac   entity.Account
	}
	tests := []struct {
		name           string
		storageAccount storage.Account
		auth           auth.Auth
		args           args
		want           entity.Account
		wantErr        bool
		err            error
	}{
		{
			name: "Success",
			storageAccount: tests.AccountCustomMock{
				SaveAccountMock: func(mctx mcontext.Context, ac entity.Account) error {
					return nil
				},
			},
			auth: auth.AuthCustomMock{
				GenerateHashPasswordMock: func(password string) (string, error) {
					return "hashedPassword", nil
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				ac: entity.Account{
					Id:        "any_id",
					Name:      "any_name",
					Cpf:       "96483478593",
					Secret:    "password",
					Balance:   100,
					CreatedAt: time.Time{},
				},
			},
			want: entity.Account{
				Id:        "any_id",
				Name:      "any_name",
				Cpf:       "96483478593",
				Secret:    "hashedPassword",
				Balance:   100,
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "Error on generate hashedPassword",
			auth: auth.AuthCustomMock{
				GenerateHashPasswordMock: func(password string) (string, error) {
					return "", errors.New("some error")
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				ac:   entity.Account{Cpf: "96483478593"},
			},
			want:    entity.Account{},
			wantErr: true,
			err:     errors.New("some error"),
		},
		{
			name: "Error on save account",
			storageAccount: tests.AccountCustomMock{
				SaveAccountMock: func(mctx mcontext.Context, ac entity.Account) error {
					return errors.New("some error")
				},
			},
			auth: auth.AuthCustomMock{
				GenerateHashPasswordMock: func(password string) (string, error) {
					return "hashedPassword", nil
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				ac: entity.Account{
					Cpf:    "96483478593",
					Secret: "password",
				},
			},
			wantErr: true,
			err:     errors.New("some error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := accounts.NewManager(tt.storageAccount, tt.auth)
			got, err := m.Create(tt.args.mctx, tt.args.ac)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
