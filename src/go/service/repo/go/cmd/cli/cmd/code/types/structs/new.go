package structs

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/package/cmd/flags"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
)

var (
	NewCmd = cmd.New(
		cmd.Use("new"), cmd.Alias("n"),
		cmd.Flags(
			flags.StringP(&service.Dest, "file", "f", "", ""),
			flags.StringP(&service.Pkg, "package", "p", "", ""),
			flags.StringP(&service.Name, "name", "n", "", ""),
		),
		cmd.Run(newStruct),
	)
)

func newStruct(cmd *cobra.Command, _ []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(s.Create(cmd.Context(), model.CreateReq{
		FilePath: model.FilePath{
			Name: service.Dest,
			Dir:  service.Dir,
		},
		File: nil,
	}))
}
