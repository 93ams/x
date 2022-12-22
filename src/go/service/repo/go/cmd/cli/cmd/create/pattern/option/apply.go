package option

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
	applyFlags model.Apply
	ApplyCmd   = cmd.New(
		cmd.Use("apply"),
		cmd.Alias("a"),
		cmd.Run(runApply),
	)
)

func runApply(cmd *cobra.Command, strings []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(x.NewFile(service.File, func(file *os.File) error {
		return s.Create(file, model.CreateReq{
			Pkg:   service.Pkg,
			Props: applyFlags,
		})
	}))
}
