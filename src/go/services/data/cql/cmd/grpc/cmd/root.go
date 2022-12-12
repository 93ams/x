package cmd

var RootCmd = New(
	Use("grpc"),
	Add(ServeCmd),
)

func Execute() error { return RootCmd.Execute() }
