package _package

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/google/wire"
	"github.com/tilau2328/cql/package/domain"
	"gql/package/exec"
	"gql/package/resolver"
)

var Set = wire.NewSet(
	domain.Set,
	handler.NewDefaultServer,
	exec.NewExecutableSchema,
	resolver.NewResolver,
)
