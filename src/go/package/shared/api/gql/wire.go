package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/google/wire"
)

var Set = wire.NewSet(handler.NewDefaultServer)
