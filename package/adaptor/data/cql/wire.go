package cql

import (
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
	"github.com/tilau2328/cql/package/domain/provider"
	"github.com/tilau2328/cql/package/shared/data/cql"
)

var Set = wire.NewSet(cql.Set,
	ddl.NewKeySpaceRepo, wire.Bind(new(provider.KeySpaceProvider), new(*ddl.KeySpaceRepo)),
	ddl.NewTableRepo, wire.Bind(new(provider.TableProvider), new(*ddl.TableRepo)),
)