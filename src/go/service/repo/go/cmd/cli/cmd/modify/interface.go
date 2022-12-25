package modify

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/package/cmd/flags"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"strings"
)

var (
	InterfaceCmd = cmd.New(
		cmd.Use("interface"),
		cmd.Flags(
			flags.StringP(&service.File, "file", "f", "", ""),
			flags.StringP(&service.Pkg, "package", "p", "", ""),
			flags.StringP(&service.Name, "name", "n", "", ""),
		),
		cmd.Run(createInterface),
	)
)

func createInterface(cmd *cobra.Command, args []string) {
	s := service.FromCtx(cmd.Context())
	lo.Must0(s.Create(cmd.Context(), model.CreateReq{
		FilePath: model.FilePath{
			Name: service.File,
			Dir:  service.Dir,
		},
		File: &model.File{},
	}))
}

func mapInterfaceMethods(args []string) *model.FieldList {
	return &model.FieldList{List: lo.Map(args, func(item string, index int) *model.Field {
		lb, rb := strings.IndexRune(item, '('), strings.IndexRune(item, ')')
		return &model.Field{
			Names: model.Names([]string{item[:lb]}),
			Type: &model.FuncType{
				Params:  mapMethodTypes(item[lb+1 : rb]),
				Results: mapMethodTypes(item[rb+1:]),
			},
		}
	})}
}

func mapMethodTypes(s string) *model.FieldList {
	return &model.FieldList{List: lo.Map(strings.Split(
		strings.TrimSuffix(strings.TrimPrefix(
			strings.TrimSpace(s),
			"("), ")"), ","), func(item string, _ int) *model.Field {
		parts := strings.Split(strings.TrimSpace(item), " ")
		var typ string
		var names []string
		switch len(parts) {
		case 1:
			typ = parts[0]
		case 2:
			names = append(names, parts[0])
			typ = parts[1]
		}
		parts = strings.Split(typ, ".")
		var path string
		if len(parts) == 2 {
			path, typ = parts[0], parts[1]
		}
		return &model.Field{
			Names: model.Names(names),
			Type: &model.Ident{
				Name: typ,
				Path: path,
			},
		}
	})}
}
