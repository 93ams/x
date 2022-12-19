//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/tilau2328/x/src/go/cmd/grpc/package/handler"
	"github.com/tilau2328/x/src/go/package/domain"
	cql2 "github.com/tilau2328/x/src/go/package/shared/data/cql"
)

var Set = wire.NewSet(
	wire.Struct(new(handler.DMLOptions), "*"), handler.NewDML,
	wire.Struct(new(handler.DDLOptions), "*"), handler.NewDDLHandler,
	handler.NewServer,
)

func Init(cql2.Options) (handler.Server, func(), error) {
	panic(wire.Build(cql.Set, domain.Set, Set))
}
