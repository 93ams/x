//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/provider"
	"github.com/tilau2328/x/src/go/service/repo/go/package/service"
)

var Set = wire.NewSet(
	service.NewService,
	wire.Bind(new(provider.GolangProvider), new(*service.Service)),
)

func Init() (provider.GolangProvider, func(), error) {
	panic(wire.Build(Set))
}
