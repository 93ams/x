package proto

import "github.com/tilau2328/x/src/go/package/cmd"

var RootCmd = cmd.New(
	cmd.Use("proto"),
	cmd.Add(CreateCmd, SearchCmd, ModifyCmd),
)
