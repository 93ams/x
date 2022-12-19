package cmd

import . "github.com/tilau2328/x/src/go/package/cmd"

var RootCmd = New(
	Use("xgo"),
	Add(NewCmd, GrepCmd),
)

func Execute() error { return RootCmd.Execute() }
