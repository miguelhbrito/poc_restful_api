package accounts

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/auth"
	"github.com/stone_assignment/pkg/mcontext"
	"github.com/stone_assignment/pkg/storage"
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
			storageAccount: storage.AccountCustomMock{
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
			storageAccount: storage.AccountCustomMock{
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
			m := NewManager(tt.storageAccount, tt.auth)
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

func TestManager_GetById(t *testing.T) {
	type fields struct {
		accountStorage storage.Account
		auth           auth.Auth
	}
	type args struct {
		mctx mcontext.Context
		id   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Account
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				accountStorage: storage.AccountCustomMock{
					GetByIdAccountMock: func(mctx mcontext.Context, id string) (entity.Account, error) {
						return entity.Account{
							Id:        "any_id",
							Name:      "any_name",
							Cpf:       "96483478593",
							Secret:    "password",
							Balance:   100,
							CreatedAt: time.Time{},
						}, nil
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				id:   "any_id",
			},
			want: entity.Account{
				Id:        "any_id",
				Name:      "any_name",
				Cpf:       "96483478593",
				Secret:    "password",
				Balance:   100,
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "Handler error",
			fields: fields{
				accountStorage: storage.AccountCustomMock{
					GetByIdAccountMock: func(mctx mcontext.Context, id string) (entity.Account, error) {
						return entity.Account{}, errors.New("some error")
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				id:   "any_id",
			},
			want:    entity.Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.fields.accountStorage, tt.fields.auth)
			got, err := m.GetById(tt.args.mctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_GetByCpf(t *testing.T) {
	type fields struct {
		AccountStorage storage.Account
		Auth           auth.Auth
	}
	type args struct {
		mctx mcontext.Context
		cpf  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Account
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					GetByCpfAccountMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
						return entity.Account{
							Id:        "any_id",
							Name:      "any_name",
							Cpf:       "96483478593",
							Secret:    "password",
							Balance:   100,
							CreatedAt: time.Time{},
						}, nil
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				cpf:  "96483478593",
			},
			want: entity.Account{
				Id:        "any_id",
				Name:      "any_name",
				Cpf:       "96483478593",
				Secret:    "password",
				Balance:   100,
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "Handler error",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					GetByCpfAccountMock: func(mctx mcontext.Context, cpf string) (entity.Account, error) {
						return entity.Account{}, errors.New("some error")
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				cpf:  "wrong_cpf",
			},
			want:    entity.Account{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.fields.AccountStorage, tt.fields.Auth)
			got, err := m.GetByCpf(tt.args.mctx, tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByCpf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByCpf() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_List(t *testing.T) {
	type fields struct {
		AccountStorage storage.Account
		Auth           auth.Auth
	}
	type args struct {
		mctx mcontext.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Accounts
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					ListAccountMock: func(mctx mcontext.Context) ([]entity.Account, error) {
						return []entity.Account{
							{
								Id:        "any_id",
								Name:      "any_name",
								Cpf:       "96483478593",
								Secret:    "password",
								Balance:   100,
								CreatedAt: time.Time{},
							},
							{
								Id:        "another_id",
								Name:      "another_name",
								Cpf:       "85372367482",
								Secret:    "another_password",
								Balance:   200,
								CreatedAt: time.Time{},
							},
						}, nil
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
			},
			want: []entity.Account{
				{
					Id:        "any_id",
					Name:      "any_name",
					Cpf:       "96483478593",
					Secret:    "password",
					Balance:   100,
					CreatedAt: time.Time{},
				},
				{
					Id:        "another_id",
					Name:      "another_name",
					Cpf:       "85372367482",
					Secret:    "another_password",
					Balance:   200,
					CreatedAt: time.Time{},
				},
			},
			wantErr: false,
		},
		{
			name: "Handler error",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					ListAccountMock: func(mctx mcontext.Context) ([]entity.Account, error) {
						return nil, errors.New("some error")
					},
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
			m := NewManager(tt.fields.AccountStorage, tt.fields.Auth)
			got, err := m.List(tt.args.mctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_Delete(t *testing.T) {
	type fields struct {
		AccountStorage storage.Account
		Auth           auth.Auth
	}
	type args struct {
		mctx mcontext.Context
		id   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					DeleteAccountMock: func(mctx mcontext.Context, id string) error {
						return nil
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				id:   "any_id",
			},
			wantErr: false,
		},
		{
			name: "Handler error",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					DeleteAccountMock: func(mctx mcontext.Context, id string) error {
						return errors.New("some error")
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				id:   "any_id",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.fields.AccountStorage, tt.fields.Auth)
			if err := m.Delete(tt.args.mctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_Update(t *testing.T) {
	type fields struct {
		AccountStorage storage.Account
		Auth           auth.Auth
	}
	type args struct {
		mctx mcontext.Context
		ac   entity.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					UpdateAccountMock: func(mctx mcontext.Context, ac entity.Account) error {
						return nil
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				ac: entity.Account{
					Id:      "any_id",
					Balance: 150,
				},
			},
			wantErr: false,
		},
		{
			name: "Handler error",
			fields: fields{
				AccountStorage: storage.AccountCustomMock{
					UpdateAccountMock: func(mctx mcontext.Context, ac entity.Account) error {
						return errors.New("some error")
					},
				},
			},
			args: args{
				mctx: mcontext.NewContext(),
				ac: entity.Account{
					Id:      "any_id",
					Balance: 150,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager(tt.fields.AccountStorage, tt.fields.Auth)
			if err := m.Update(tt.args.mctx, tt.args.ac); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
