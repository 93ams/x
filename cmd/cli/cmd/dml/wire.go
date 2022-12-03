//go:build wireinject
// +build wireinject

package dml

import (
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/adaptor/data/cql"
	"github.com/tilau2328/cql/package/domain"
	"github.com/tilau2328/cql/package/domain/provider"
	cql2 "github.com/tilau2328/cql/package/shared/data/cql"
)

func Init(cql2.Options) (provider.DML, func(), error) {
	panic(wire.Build(cql.Set, domain.Set))
}
