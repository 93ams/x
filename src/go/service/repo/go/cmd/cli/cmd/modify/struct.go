package modify

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/package/cmd/flags"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
)

var (
	StructCmd = cmd.New(
		cmd.Use("struct"),
		cmd.Flags(
			flags.StringP(&service.File, "file", "f", "", ""),
			flags.StringP(&service.Pkg, "package", "p", "", ""),
			flags.StringP(&service.Name, "name", "n", "", ""),
		),
		cmd.Run(createStruct),
	)
)

func createStruct(cmd *cobra.Command, _ []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(s.Create(cmd.Context(), model.CreateReq{
		FilePath: model.FilePath{
			Name: service.File,
			Dir:  service.Dir,
		},
		File: nil,
	}))
}
