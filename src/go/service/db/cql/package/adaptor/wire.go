package adaptor

import (
	"github.com/google/wire"
	"github.com/tilau2328/x/src/go/package/adaptor/data/cql/repo/ddl"
	"github.com/tilau2328/x/src/go/package/domain/provider"
	"github.com/tilau2328/x/src/go/package/shared/data/cql"
)

var Set = wire.NewSet(cql.Set,
	ddl.NewKeySpaceRepo, wire.Bind(new(provider.KeySpaceProvider), new(*ddl.KeySpaceRepo)),
	ddl.NewColumnRepo, wire.Bind(new(provider.ColumnProvider), new(*ddl.ColumnRepo)),
	ddl.NewTableRepo, wire.Bind(new(provider.TableProvider), new(*ddl.TableRepo)),
)
