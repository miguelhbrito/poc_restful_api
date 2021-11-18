package transfers

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

type TransferCustomMock struct {
	CreateMock func(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error)
	ListMock   func(mctx mcontext.Context) (entity.Transfers, error)
}

func (t TransferCustomMock) Create(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error) {
	return t.CreateMock(mctx, tr)
}

func (t TransferCustomMock) List(mctx mcontext.Context) (entity.Transfers, error) {
	return t.ListMock(mctx)
}
