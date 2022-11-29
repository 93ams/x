package cmd

import (
	"github.com/samber/lo"
	. "github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	. "github.com/spf13/pflag"
	. "github.com/tilau2328/cql/internal/generic"
	"io"
	"os"
	"path/filepath"
)

type Fn func(*Command, []string)
type FnE func(*Command, []string) error

func New(opts ...Opt[*Command]) *Command     { return Apply(&Command{}, opts) }
func Run(fn Fn) Opt[*Command]                { return func(cmd *Command) { cmd.Run = fn } }
func RunE(fn FnE) Opt[*Command]              { return func(cmd *Command) { cmd.RunE = fn } }
func Use(use string) Opt[*Command]           { return func(cmd *Command) { cmd.Use = use } }
func Alias(a ...string) Opt[*Command]        { return func(cmd *Command) { cmd.Aliases = a } }
func Add(c ...*Command) Opt[*Command]        { return func(cmd *Command) { cmd.AddCommand(c...) } }
func Flags(o ...Opt[*FlagSet]) Opt[*Command] { return func(cmd *Command) { Apply(cmd.Flags(), o) } }
func PersistentPreRun(fn Fn) Opt[*Command]   { return func(cmd *Command) { cmd.PersistentPreRun = fn } }
func PersistentFlags(o ...Opt[*FlagSet]) Opt[*Command] {
	return func(cmd *Command) { Apply(cmd.PersistentFlags(), o) }
}
func Required(f ...string) Opt[*Command] {
	return func(cmd *Command) { lo.Must0(ApplyE(f, cmd.MarkFlagRequired)) }
}
func PersistentRequired(f ...string) Opt[*Command] {
	return func(cmd *Command) { lo.Must0(ApplyE(f, cmd.MarkPersistentFlagRequired)) }
}
func GenDocument(cmd *Command, dir string) error {
	identity := func(s string) string { return "" }
	emptyStr := func(s string) string { return "" }
	return GenDocumentCustom(cmd, dir, identity, emptyStr)
}
func GenDocumentCustom(cmd *Command, dir string, prepender, linkHandler func(string) string) error {
	f, err := os.Create(filepath.Join(dir, "README.md"))
	if err != nil {
		return err
	}
	defer f.Close()
	return genDocument(cmd, f, prepender, linkHandler)
}
func genDocument(cmd *Command, f io.Writer, prepender, linkHandler func(string) string) error {
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
			continue
		}
		if err := genDocument(c, f, prepender, linkHandler); err != nil {
			return err
		}
	}

	if _, err := io.WriteString(f, prepender(cmd.Name())); err != nil {
		return err
	} else if err := doc.GenMarkdownCustom(cmd, f, linkHandler); err != nil {
		return err
	}
	return nil
}
