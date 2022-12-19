//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/tilau2328/x/src/go/cmd/gql/package/exec"
	"github.com/tilau2328/x/src/go/cmd/gql/package/resolver"
	"github.com/tilau2328/x/src/go/package/domain"
	"github.com/tilau2328/x/src/go/package/shared/api/gql"
	cql2 "github.com/tilau2328/x/src/go/package/shared/data/cql"
)

var Set = wire.NewSet(wire.Bind(new(exec.ResolverRoot), new(*resolver.Resolver)), resolver.NewResolver,
	wire.Struct(new(exec.Config), "Resolvers"), exec.NewExecutableSchema)

func Init(cql2.Options, gql.Endpoint) (*gql.Server, func(), error) {
	panic(wire.Build(cql.Set, domain.Set, Set, gql.Set))
}
