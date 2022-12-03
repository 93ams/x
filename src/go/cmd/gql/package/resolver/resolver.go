package resolver

import (
	provider2 "github.com/tilau2328/cql/package/domain/provider"
)

type Resolver struct {
	ddl provider2.DDL
	dml provider2.DML
}

func NewResolver(
	ddl provider2.DDL,
	dml provider2.DML,
) *Resolver {
	return &Resolver{
		ddl: ddl,
		dml: dml,
	}
}
