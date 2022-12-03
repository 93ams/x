//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gocql/gocql"
	"github.com/google/wire"
	cql2 "github.com/tilau2328/cql/package/adaptor/data/cql"
)

func NewServer(cluster *gocql.ClusterConfig) (*handler.Server, error, func()) {
	panic(wire.Build(cql2.Set, Set))
}
