package x

import "context"

func FromCtx[K any, V any](ctx context.Context, key K) (ret V) {
	ret, _ = ctx.Value(key).(V)
	return
}
