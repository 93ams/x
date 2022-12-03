//go:build wireinject
// +build wireinject

package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gocql/gocql"
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/adaptor/data/cql"
)

func NewServer(cluster *gocql.ClusterConfig) (*handler.Server, error, func()) {
	panic(wire.Build(cql.Set, Set))
}
