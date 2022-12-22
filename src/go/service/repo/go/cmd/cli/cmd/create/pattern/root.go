package pattern

import (
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/create/pattern/option"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/create/pattern/service"
)

var RootCmd = cmd.New(
	cmd.Use("pattern"),
	cmd.Alias("p"),
	cmd.Add(
		EnumCmd,
		MapperCmd,
		option.RootCmd,
		service.RootCmd,
	),
)
