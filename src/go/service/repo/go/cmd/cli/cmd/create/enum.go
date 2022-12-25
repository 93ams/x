package create

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/package/cmd/flags"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"path/filepath"
	"strings"
)

var (
	EnumCmd = cmd.New(
		cmd.Use("enum"),
		cmd.Alias("e"),
		cmd.Flags(
			flags.String(&service.Dir, "dir", "", "."),
			flags.String(&service.From, "from", "", ""),
			flags.String(&service.File, "file", "", ""),
			flags.String(&service.Pkg, "package", "", "main"),
		),
		cmd.Run(runEnum),
	)
)

func runEnum(cmd *cobra.Command, args []string) {
	s := service.FromCtx(cmd.Context())
	li := strings.LastIndex(service.From, ".")
	path, name := service.From[:li], service.From[li+1:]
	search := lo.Must(s.Search(cmd.Context(), model.SearchReq{
		Type: model.TypeFilter{Name: name},
		FilePath: model.FilePath{
			Dir:  service.Dir,
			Name: path,
		},
	}))
	path = filepath.Join(service.Dir, path)
	if len(search[path]) != 1 {
		return
	}
	node, ok := search[path][0].(*model.Type)
	if !ok {
		return
	}
	lo.Must0(s.Create(cmd.Context(), model.CreateReq{
		FilePath: model.FilePath{
			Name: service.Name,
			Dir:  service.Dir,
		},
		File: &model.File{
			Name: &model.Ident{Name: service.Pkg},
			Decls: model.Enum{
				Type: node.Name.Clone(),
				Values: lo.SliceToMap(args, func(item string) (string, string) {
					parts := strings.Split(item, ":")
					switch len(parts) {
					case 0:
						return "", ""
					case 1:
						return item, ""
					default:
						return parts[0], parts[1]
					}
				}),
			}.Decl(),
		},
	}))
}
