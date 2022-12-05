package cmd

import (
	. "github.com/tilau2328/cql/src/go/package/shared/cmd"
)

var RootCmd = New(
	Use("grpc"),
	Add(ServeCmd),
)

func Execute() error { return RootCmd.Execute() }
