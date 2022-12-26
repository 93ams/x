package files

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"log"
	"os"
)

var FetchCmd = cmd.New(
	cmd.Use("fetch"),
	cmd.Alias("f"),
	cmd.Run(fetchFiles),
)

func fetchFiles(cmd *cobra.Command, args []string) {
	s := service.FromCtx(cmd.Context())
	files := lo.Must(s.Read(cmd.Context(), lo.Map(args, func(item string, index int) model.ReadReq {
		return model.ReadReq{FilePath: model.FilePath{Name: item, Dir: service.Dir}}
	})...))
	for k, v := range files {
		log.Println("****************  " + k + "  ****************")
		lo.Must0(driver.Write(os.Stdout, v))
	}
}
