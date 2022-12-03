package cql

import (
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
	provider2 "github.com/tilau2328/cql/package/domain/provider"
	"github.com/tilau2328/cql/src/go/package/shared/data/cql"
)

var Set = wire.NewSet(cql.Set,
	ddl.NewKeySpaceRepo, wire.Bind(new(provider2.KeySpaceProvider), new(*ddl.KeySpaceRepo)),
	ddl.NewTableRepo, wire.Bind(new(provider2.TableProvider), new(*ddl.TableRepo)),
)
