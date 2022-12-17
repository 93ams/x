package cmd

var RootCmd = New(
	Use("xgo"),
	Add(NewCmd, GrepCmd),
)

func Execute() error { return RootCmd.Execute() }
