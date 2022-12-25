package create

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
)

var (
	MapperCmd = cmd.New(
		cmd.Use("mapper"),
		cmd.Alias("m"),
		cmd.Run(runMapper),
	)
)

func runMapper(cmd *cobra.Command, strings []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(s.Create(cmd.Context(), model.CreateReq{
		FilePath: model.FilePath{
			Name: service.Name,
			Dir:  service.Dir,
		},
		File: nil,
	}))
}
