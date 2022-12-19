package adaptor

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(cql.Set,
	ddl.NewKeySpaceRepo, wire.Bind(new(provider.KeySpaceProvider), new(*ddl.KeySpaceRepo)),
	ddl.NewColumnRepo, wire.Bind(new(provider.ColumnProvider), new(*ddl.ColumnRepo)),
	ddl.NewTableRepo, wire.Bind(new(provider.TableProvider), new(*ddl.TableRepo)),
)
