package load

import (
	"context"
	. "github.com/graph-gophers/dataloader"
	"github.com/samber/lo"
)

type (
	Res[V any] struct {
		Data V
		Err  error
	}
	BatchFn[K Key, V any]    func(context.Context, []K) []Res[V]
	Dataloader[K Key, V any] struct{ *Loader }
	DataLoader[K Key, V any] interface {
		Load(context.Context, K) Res[V]
		LoadMany(context.Context, []K) []Res[V]
		Clear(context.Context, K) DataLoader[K, V]
		Prime(context.Context, K, V) DataLoader[K, V]
	}
)

func NewLoader[K Key, V any](fn BatchFn[K, V], opts ...Option) Dataloader[K, V] {
	return Dataloader[K, V]{NewBatchedLoader(func(ctx context.Context, keys Keys) []*Result {
		return FromRes(fn(ctx, FromKeys[K](keys)))
	}, opts...)}
}
func (d Dataloader[K, V]) Load(ctx context.Context, key K) Res[V] {
	thunk := d.Loader.Load(ctx, key)
	val, err := thunk()
	return Res[V]{Data: val, Err: err}
}
func (d Dataloader[K, V]) LoadMany(ctx context.Context, keys []K) []Res[V] {
	thunks := d.Loader.LoadMany(ctx, ToKeys(keys))
	return ToRes[V](thunks())
}
func (d Dataloader[K, V]) Clear(ctx context.Context, k K) DataLoader[K, V] {
	d.Loader.Clear(ctx, k)
	return d
}
func (d Dataloader[K, V]) Prime(ctx context.Context, k K, v V) DataLoader[K, V] {
	d.Loader.Prime(ctx, k, v)
	return d
}

func FromRes[V any](v []Res[V]) []*Result {
	return lo.Map(v, func(i Res[V], _ int) *Result {
		return &Result{
			Data:  i.Data,
			Error: i.Err,
		}
	})
}
func ToRes[V any](vals []any, errs []error) []Res[V] {
	return lo.Map(vals, func(v any, i int) Res[V] {
		return Res[V]{Data: v, Err: errs[i]}
	})
}
func ToKeys[K Key](k []K) Keys   { return lo.Map(k, func(i K, _ int) Key { return i }) }
func FromKeys[K Key](k Keys) []K { return lo.Map(k, func(i Key, _ int) K { return i.(K) }) }
