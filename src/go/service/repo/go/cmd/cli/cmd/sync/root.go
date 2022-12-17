package sync

var RootCmd = New(
	Use("sync"),
	Add(ImportCmd, ExportCmd),
)
