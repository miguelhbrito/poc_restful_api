package mcontext

import (
	"context"

	"github.com/stone_assignment/pkg/api"
)

type Context interface {
	context.Context
	Cpf() api.Cpf
}

type myContext struct {
	context.Context
}

func NewContext() Context {
	return myContext{Context: context.Background()}
}

func NewFrom(ctx context.Context) Context {
	return myContext{ctx}
}

func WithValue(ctx Context, key interface{}, val interface{}) Context {
	return NewFrom(context.WithValue(ctx, key, val))
}

func (ctx myContext) Cpf() api.Cpf {
	user, ok := ctx.Value(api.CpfCtxKey).(api.Cpf)
	if !ok && user.String() == "" {
		return ""
	}
	return api.Cpf(user)
}
