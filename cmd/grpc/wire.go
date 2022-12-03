//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/adaptor/data/cql"
	cql2 "github.com/tilau2328/cql/package/shared/cql"
	"google.golang.org/grpc"
	. "grpc/package/handler"
)

var Set = wire.NewSet(NewServer, NewDML, NewDDL)

func Init(cql2.Options) (*grpc.Server, func(), error) {
	panic(wire.Build(cql.Set, domain.Set, Set))
}
