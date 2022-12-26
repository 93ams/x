package files

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/tilau2328/x/src/go/package/cmd"
	"github.com/tilau2328/x/src/go/package/cmd/flags"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/cmd/files/deps"
	"github.com/tilau2328/x/src/go/service/repo/go/cmd/cli/internal/service"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

var (
	RootCmd = cmd.New(
		cmd.Use("files"), cmd.Alias("f"),
		cmd.Add(FetchCmd, NewCmd, RemCmd, deps.RootCmd),
		cmd.Flags(flags.BoolP(&service.Recursive, "recursive", "r", "", false)),
		cmd.Run(listFiles),
	)
)

func listFiles(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		args = append(args, ".")
	}
	lock := sync.Mutex{}
	paths := map[string]struct{}{}
	lo.Must0(x.ParallelTry(func(t string) error {
		res, err := x.ListFiles(t, service.Recursive)
		if err != nil {
			return err
		}
		lock.Lock()
		for _, v := range res {
			paths[v] = struct{}{}
		}
		lock.Unlock()
		return nil
	}, args))
	res := lo.Keys(paths)
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	tree := x.Tree[string]{}
	for _, k := range res {
		tree = tree.Add(strings.Split(k, string(os.PathSeparator)))
	}
	tree.Walk(x.PathVisitor[string]{}.Visitor(func(path []string) {
		i := len(path) - 1
		log.Println(strings.Repeat(" * ", i), path[i])
	}))
}
