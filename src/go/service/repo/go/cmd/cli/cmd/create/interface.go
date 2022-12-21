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
	"strings"
)

var (
	interfaceFlags model.Interface
	InterfaceCmd   = cmd.New(
		cmd.Use("interface"),
		cmd.Flags(
			flags.StringP(&service.File, "file", "f", "", ""),
			flags.StringP(&service.Pkg, "package", "p", "", ""),
			flags.StringP(&interfaceFlags.Name, "name", "n", "", ""),
		),
		cmd.Run(createInterface),
	)
)

func createInterface(cmd *cobra.Command, args []string) {
	s := service.FromCtx(cmd.Context())
	interfaceFlags.Methods = mapInterfaceArgs(args)
	lo.Must0(x.NewFile(service.File, func(file *os.File) error {
		return s.Create(file, model.CreateReq{
			Pkg:   service.Pkg,
			Props: interfaceFlags,
		})
	}))
}

func mapInterfaceArgs(args []string) []model.FuncType {
	return lo.Map(args, func(item string, index int) model.FuncType {
		lb, rb := strings.IndexRune(item, '('), strings.IndexRune(item, ')')
		return model.FuncType{
			Name: item[:lb],
			In:   mapMethodTypes(item[lb+1 : rb]),
			Out:  mapMethodTypes(item[rb+1:]),
		}
	})
}

func mapMethodTypes(s string) []model.MethodType {
	return lo.Map(strings.Split(
		strings.TrimSuffix(strings.TrimPrefix(
			strings.TrimSpace(s),
			"("), ")"), ","), func(item string, _ int) model.MethodType {
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
		return model.MethodType{
			Names: names,
			Type: model.Type{
				Name: typ,
				Path: path,
			},
		}
	})
}
