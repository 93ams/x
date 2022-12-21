package cmd

import (
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/api/grpc/cmd/cli/cmd/gen"
	"github.com/tilau2328/x/src/go/service/api/grpc/cmd/cli/cmd/proto"
)

var RootCmd = cmd.New(
	cmd.Use("xgrpc"),
	cmd.Add(gen.RootCmd, proto.RootCmd),
)

func Execute() error {
	return RootCmd.Execute()
}
