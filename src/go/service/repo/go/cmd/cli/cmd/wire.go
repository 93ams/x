//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/tilau2328/x/src/go/service/repo/go/package/services"
	"github.com/tilau2328/x/src/go/service/repo/go/package/services/provider"
)

var Set = wire.NewSet(
	services.NewService,
	wire.Bind(new(provider.GolangProvider), new(*services.Service)),
)

func Init() (provider.GolangProvider, func(), error) {
	panic(wire.Build(Set))
}
