package loader

import (
	. "context"
	. "github.com/tilau2328/cql/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/package/shared/load"
)

func NewKeySpaceLoader(provider KeySpaceProvider) Dataloader[KeySpaceKey, KeySpace] {
	return NewLoader(func(ctx Context, keys []KeySpaceKey) []Res[KeySpace] {
		var results []Res[KeySpace]
		provider.List(ctx, KeySpace{})
		return results
	})
}
