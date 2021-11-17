package login

import (
	"github.com/stone_assignment/pkg/api/entity"
	"github.com/stone_assignment/pkg/api/response"
	"github.com/stone_assignment/pkg/mcontext"
)

type LoginCustomMock struct {
	LoginIntoSystemMock func(mctx mcontext.Context, l entity.LoginEntity) (response.LoginToken, error)
}

func (l LoginCustomMock) LoginIntoSystem(mctx mcontext.Context, le entity.LoginEntity) (response.LoginToken, error) {
	return l.LoginIntoSystemMock(mctx, le)
}
