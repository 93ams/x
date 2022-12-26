package funcs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"os"
)

var RootCmd = cmd.New(
	cmd.Use("funcs"), cmd.Alias("f"),
	cmd.Add(EditCmd, FetchCmd, NewCmd, RemCmd),
	cmd.Run(listFuncs),
)

func listFuncs(cmd *cobra.Command, args []string) {
	s := service.FromCtx(cmd.Context())
	idMap := map[string]string{}
	res := lo.Must(s.Search(cmd.Context(), lo.Map(args, func(item string, index int) model.SearchReq {
		id := uuid.NewString()
		idMap[id] = item
		return model.SearchReq{
			Id: id,
			FilePath: model.FilePath{
				Name: item,
				Dir:  service.Dir,
			},
			Func:    &model.FuncFilter{},
			Package: true,
		}
	})...))
	for k, v := range res {
		fmt.Println("// ****************  " + idMap[k] + "  ****************")
		lo.Must0(driver.Write(os.Stdout, &model.File{
			Name: v[0].(*model.Ident),
			Decls: lo.Map(v[1:], func(item model.Node, _ int) model.Decl {
				f := item.(*model.Func)
				f.Body = nil
				return f
			}),
		}))
	}
}
