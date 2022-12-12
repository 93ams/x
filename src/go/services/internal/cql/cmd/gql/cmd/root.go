package cmd

var RootCmd = New(
	Use("gql"),
	Add(ServeCmd),
)

func Execute() error { return RootCmd.Execute() }
