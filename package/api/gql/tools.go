//go:build tools
// +build tools

package gql

//go:generate gqlgen generate --config config.yaml

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
