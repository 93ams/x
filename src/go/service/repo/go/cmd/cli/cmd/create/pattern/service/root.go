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
	strategyFlags model.Service
	RootCmd       = cmd.New(
		cmd.Use("service"),
		cmd.Run(runStrategy),
		cmd.Alias("s"),
		cmd.Add(BuilderCmd),
	)
)

func runStrategy(cmd *cobra.Command, strings []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(x.NewFile(service.File, func(file *os.File) error {
		return s.Create(file, model.CreateReq{
			Pkg:   service.Pkg,
			Props: strategyFlags,
		})
	}))
}
