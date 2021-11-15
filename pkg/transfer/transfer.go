package transfer

import (
	"errors"

	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/mcontext"
)

var (
	errBalanceLowerThan0 = errors.New("Origin balance is equal or lower 0")
	errBalanceCantAffort = errors.New("Origin balance can not affort this ammount")
)

type TransferManager interface {
	Create(mctx mcontext.Context, tr entity.Transfer) (entity.Transfer, error)
	List(mctx mcontext.Context) (entity.Transfer, error)
}
