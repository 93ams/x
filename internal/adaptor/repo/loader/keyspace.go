package loader

import (
	. "context"
	. "github.com/tilau2328/cql/internal/adaptor/repo"
	. "github.com/tilau2328/cql/internal/adaptor/repo/model"
	. "github.com/tilau2328/cql/package/load"
)

func NewKeySpaceLoader(provider KeySpaceProvider) Dataloader[KeySpaceKey, KeySpace] {
	return NewLoader(func(ctx Context, keys []KeySpaceKey) []Res[KeySpace] {
		var results []Res[KeySpace]
		// SELECT * FROM system_schema.keyspace WHERE keyspace_name IN (...)
		provider.List(ctx, KeySpace{})
		return results
	})
}
