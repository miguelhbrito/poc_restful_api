package transfers

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stone_assignment/pkg/accounts"
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/storage"
)

func Test_manager_Create(t *testing.T) {
	type args struct {
		mctx mcontext.Context
		tr   entity.Transfer
	}
	tests := []struct {
		name            string
		m               Transfer
		transferStorage storage.Transfer
		accountManager  accounts.Account
		args            args
		want            entity.Transfer
		wantErr         bool
	}{
		{
			name: "Success",
			transferStorage: storage.TransferCustomMock{
				SaveTransferMock: func(mctx mcontext.Context, tr entity.Transfer) error {
					return nil
				},
			},
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{
						Id:        "any_id_1",
						Balance:   100,
						CreatedAt: time.Time{},
					}, nil
				},
				GetByIdMock: func(mctx mcontext.Context, id string) (entity.Account, error) {
					return entity.Account{
						Id:        "any_id_2",
						Balance:   100,
						CreatedAt: time.Time{},
					}, nil
				},
				UpdateMock: func(mctx mcontext.Context, ac entity.Account) error {
					return nil
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				tr: entity.Transfer{
					Id:              "any_id_transfer",
					AccountOriginId: "any_id_1",
					AccountDestId:   "any_id_2",
					Ammount:         25.50,
					CreatedAt:       time.Time{},
				},
			},
			want: entity.Transfer{
				Id:              "any_id_transfer",
				AccountOriginId: "any_id_1",
				AccountDestId:   "any_id_2",
				Ammount:         25.50,
				CreatedAt:       time.Time{},
			},
			wantErr: false,
		},
		{
			name: "Error on get account by cpf",
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{}, errors.New("some error")
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
			},
			wantErr: true,
		},
		{
			name: "Error on get account by id",
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{
						Id:        "any_id_1",
						Balance:   100,
						CreatedAt: time.Time{},
					}, nil
				},
				GetByIdMock: func(mctx mcontext.Context, id string) (entity.Account, error) {
					return entity.Account{}, errors.New("some error")
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
			},
			wantErr: true,
		},
		{
			name: "Error on update account",
			transferStorage: storage.TransferCustomMock{
				SaveTransferMock: func(mctx mcontext.Context, tr entity.Transfer) error {
					return nil
				},
			},
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{
						Id:        "any_id_1",
						Balance:   100,
						CreatedAt: time.Time{},
					}, nil
				},
				GetByIdMock: func(mctx mcontext.Context, id string) (entity.Account, error) {
					return entity.Account{
						Id:        "any_id_2",
						Balance:   100,
						CreatedAt: time.Time{},
					}, nil
				},
				UpdateMock: func(mctx mcontext.Context, ac entity.Account) error {
					return errors.New("some error")
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				tr: entity.Transfer{
					Id:              "any_id_transfer",
					AccountOriginId: "any_id_1",
					AccountDestId:   "any_id_2",
					Ammount:         25.50,
					CreatedAt:       time.Time{},
				},
			},
			want:    entity.Transfer{},
			wantErr: true,
		},
		{
			name: "Error on save transfer",
			transferStorage: storage.TransferCustomMock{
				SaveTransferMock: func(mctx mcontext.Context, tr entity.Transfer) error {
					return errors.New("some error")
				},
			},
			accountManager: accounts.AccountCustomMock{
				GetByCpfMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
					return entity.Account{
						Id:        "any_id_1",
						Balance:   100,
						CreatedAt: time.Time{},
					}, nil
				},
				GetByIdMock: func(mctx mcontext.Context, id string) (entity.Account, error) {
					return entity.Account{
						Id:        "any_id_2",
						Balance:   100,
						CreatedAt: time.Time{},
					}, nil
				},
				UpdateMock: func(mctx mcontext.Context, ac entity.Account) error {
					return nil
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				tr: entity.Transfer{
					Id:              "any_id_transfer",
					AccountOriginId: "any_id_1",
					AccountDestId:   "any_id_2",
					Ammount:         25.50,
					CreatedAt:       time.Time{},
				},
			},
			want:    entity.Transfer{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.transferStorage, tt.accountManager)
			got, err := m.Create(tt.args.mctx, tt.args.tr)
			if (err != nil) != tt.wantErr {
				t.Errorf("manager.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_List(t *testing.T) {
	type args struct {
		mctx mcontext.Context
	}
	tests := []struct {
		name            string
		transferStorage storage.Transfer
		accountManager  accounts.Account
		args            args
		want            entity.Transfers
		wantErr         bool
	}{
		{
			name: "Success",
			transferStorage: storage.TransferCustomMock{
				ListTransfersMock: func(mctx mcontext.Context) ([]entity.Transfer, error) {
					return []entity.Transfer{{
						Id:        "any_id",
						CreatedAt: time.Time{},
					}}, nil
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
			},
			want: []entity.Transfer{{
				Id:        "any_id",
				CreatedAt: time.Time{},
			}},
			wantErr: false,
		},
		{
			name: "Error on list transfers from database",
			transferStorage: storage.TransferCustomMock{
				ListTransfersMock: func(mctx mcontext.Context) ([]entity.Transfer, error) {
					return nil, errors.New("some error")
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.transferStorage, tt.accountManager)
			got, err := m.List(tt.args.mctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
