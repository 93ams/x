package cql

import (
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/adaptor/data/cql/repo/ddl"
	"github.com/tilau2328/cql/package/shared/cql"
)

var Set = wire.NewSet(cql.Set, ddl.NewKeySpaceRepo, ddl.NewTableRepo)
