package create

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/package/cmd/flags"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"os"
)

var (
	structFlags model.Struct
	StructCmd   = cmd.New(
		cmd.Use("struct"),
		cmd.Flags(
			flags.StringP(&service.File, "file", "f", "", ""),
			flags.StringP(&service.Pkg, "package", "p", "", ""),
			flags.StringP(&structFlags.Name, "name", "n", "", ""),
		),
		cmd.Run(createStruct),
	)
)

func createStruct(cmd *cobra.Command, _ []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(x.NewFile(service.File, func(file *os.File) error {
		return s.Create(file, model.CreateReq{
			Pkg:   service.Pkg,
			Props: structFlags,
		})
	}))
}
