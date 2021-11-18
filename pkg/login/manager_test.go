package login

import (
	"errors"
	"testing"

	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/mcontext"
)

func Test_manager_LoginIntoSystem(t *testing.T) {
	type args struct {
		mctx mcontext.Context
		l    entity.LoginEntity
	}
	tests := []struct {
		name           string
		accountManager accounts.Account
		auth           auth.Auth
		args           args
		wantErr        bool
	}{
		{
			name: "Success",
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{
						Id:     "any_id",
						Cpf:    "any_cpf",
						Secret: "password",
					}, nil
				},
			},
			auth: auth.AuthCustomMock{
				CheckPasswordHashMock: func(password, hash string) bool {
					return true
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				l: entity.LoginEntity{
					Cpf:    "any_cpf",
					Secret: "password",
				},
			},
			wantErr: false,
		},
		{
			name: "Error on get by cpf",
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{}, errors.New("some error")
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				l: entity.LoginEntity{
					Cpf:    "any_cpf",
					Secret: "password",
				},
			},
			wantErr: true,
		},
		{
			name: "Error to check hashedPassword",
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{
						Id:     "any_id",
						Cpf:    "any_cpf",
						Secret: "password",
					}, nil
				},
			},
			auth: auth.AuthCustomMock{
				CheckPasswordHashMock: func(password, hash string) bool {
					return false
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				l: entity.LoginEntity{
					Cpf:    "any_cpf",
					Secret: "password",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.accountManager, tt.auth)
			_, err := m.LoginIntoSystem(tt.args.mctx, tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("manager.LoginIntoSystem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
