package service

import (
	"context"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/provider"
)

type ctxKey struct{}

func FromCtx(ctx context.Context) (ret provider.GolangProvider) {
	ret, _ = ctx.Value(ctxKey{}).(provider.GolangProvider)
	return
}

func ToCtx(ctx context.Context, val provider.GolangProvider) context.Context {
	return context.WithValue(ctx, ctxKey{}, val)
}
