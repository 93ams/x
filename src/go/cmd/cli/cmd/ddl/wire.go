//go:build wireinject
// +build wireinject

package ddl

import (
	"github.com/google/wire"
	cql3 "github.com/tilau2328/cql/package/adaptor/data/cql"
	domain2 "github.com/tilau2328/cql/package/domain"
	provider2 "github.com/tilau2328/cql/package/domain/provider"
	cql2 "github.com/tilau2328/cql/src/go/package/shared/data/cql"
)

func Init(cql2.Options) (provider2.DDL, func(), error) {
	panic(wire.Build(cql3.Set, domain2.Set))
}
