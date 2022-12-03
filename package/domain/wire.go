package domain

import (
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/domain/provider"
	"github.com/tilau2328/cql/package/domain/service"
)

var (
	Set = wire.NewSet(
		wire.Bind(new(provider.DDL), new(service.Ddl)),
		wire.Bind(new(provider.DML), new(service.Dml)),
	)
	MockSet = wire.NewSet(
		wire.Bind(new(provider.DDL), new(service.Ddl)),
		wire.Bind(new(provider.DML), new(service.Dml)),
	)
)
