package loader

import (
	. "context"
	. "github.com/tilau2328/cql/internal/adaptor/repo"
	. "github.com/tilau2328/cql/internal/adaptor/repo/model"
	. "github.com/tilau2328/cql/package/load"
)

func NewTableLoader(provider TableProvider) Dataloader[TableKey, Table] {
	return NewLoader(func(ctx Context, keys []TableKey) []Res[Table] {
		var results []Res[Table]
		// SELECT * FROM system_schema.table WHERE table_name IN (...)
		return results
	})
}
