package service

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"os"
)

var (
	builderFlags model.Builder
	BuilderCmd   = cmd.New(
		cmd.Use("builder"),
		cmd.Alias("b"),
		cmd.Run(runBuilder),
	)
)

func runBuilder(cmd *cobra.Command, args []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(x.NewFile(service.File, func(file *os.File) error {
		builderFlags.From = args[0]
		builderFlags.Filter = args[0:]
		return s.Create(file, model.CreateReq{
			Pkg:     service.Pkg,
			Props:   builderFlags,
			Pattern: true,
		})
	}))
}
