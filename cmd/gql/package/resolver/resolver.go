package resolver

import (
	"github.com/tilau2328/cql/package/domain/provider"
)

type Resolver struct {
	ddl provider.DDL
	dml provider.DML
}

func NewResolver(
	ddl provider.DDL,
	dml provider.DML,
) *Resolver {
	return &Resolver{
		ddl: ddl,
		dml: dml,
	}
}
