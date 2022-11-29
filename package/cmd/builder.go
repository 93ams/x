package cmd

import (
	"github.com/samber/lo"
	. "github.com/spf13/cobra"
	. "github.com/spf13/pflag"
	. "github.com/tilau2328/cql/internal/generic"
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
