package main

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/cmd/gql/cmd"
)

//go:generate gqlgen generate --config gencfg.yaml

func main() { lo.Must0(cmd.Execute()) }
