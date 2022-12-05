//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gocql/gocql"
	"github.com/google/wire"
	"github.com/tilau2328/cql/src/go/cmd/gql/package/exec"
	"github.com/tilau2328/cql/src/go/cmd/gql/package/resolver"
	"github.com/tilau2328/cql/src/go/package/adaptor/data/cql"
	"github.com/tilau2328/cql/src/go/package/shared/api/gql"
)

var Set = wire.NewSet(resolver.NewResolver,
	wire.Struct(new(exec.Config), "Resolvers"), exec.NewExecutableSchema)

func NewServer(cluster *gocql.ClusterConfig) (*handler.Server, error, func()) {
	panic(wire.Build(cql.Set, Set, gql.Set))
}
