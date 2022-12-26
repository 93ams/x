package cmd

import (
	"context"
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/code"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/files"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/mods"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/pkgs"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
)

type Something struct {
	Foo string
	Bar int
}
type SomethingElse interface {
	Foo(string) error
	Bar() int
}

var RootCmd = cmd.New(
	cmd.Use("xgo"),
	cmd.Add(code.RootCmd, files.RootCmd, mods.RootCmd, pkgs.RootCmd),
)

func Execute() error {
	p, cleanup := lo.Must2(Init())
	defer cleanup()
	return RootCmd.ExecuteContext(service.ToCtx(context.Background(), p))
}
