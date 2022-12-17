package domain

import (
	"github.com/google/wire"
	"github.com/tilau2328/x/src/go/package/domain/provider"
	"github.com/tilau2328/x/src/go/package/domain/service"
)

var Set = wire.NewSet(
	wire.Struct(new(service.DDLServiceOptions), "*"), service.NewDDL,
	wire.Bind(new(provider.DDL), new(*service.DDLService)),
	wire.Struct(new(service.DMLServiceOptions), "*"), service.NewDML,
	wire.Bind(new(provider.DML), new(*service.DMLService)),
)
