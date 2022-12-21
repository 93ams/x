package cmd

import (
	"context"
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/create"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/modify"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/search"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
)

var RootCmd = cmd.New(
	cmd.Use("xgo"),
	cmd.Add(create.RootCmd, search.RootCmd, modify.RootCmd),
)

func Execute() error {
	p, cleanup := lo.Must2(Init())
	defer cleanup()
	return RootCmd.ExecuteContext(
		service.ToCtx(context.Background(), p),
	)
}
