package loader

import (
	. "context"
	. "github.com/tilau2328/cql/package/adaptor/data/cql/model"
	. "github.com/tilau2328/cql/package/shared/load"
)

func NewTableLoader(provider TableProvider) Dataloader[TableKey, Table] {
	return NewLoader(func(ctx Context, keys []TableKey) []Res[Table] {
		var results []Res[Table]
		// SELECT * FROM system_schema.table WHERE table_name IN (...)
		return results
	})
}
