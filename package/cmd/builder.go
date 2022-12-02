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
func Run(fn Fn) Opt[*Command]                { return func(c *Command) { c.Run = fn } }
func RunE(fn FnE) Opt[*Command]              { return func(c *Command) { c.RunE = fn } }
func Use(use string) Opt[*Command]           { return func(c *Command) { c.Use = use } }
func Alias(a ...string) Opt[*Command]        { return func(c *Command) { c.Aliases = a } }
func Add(cmd ...*Command) Opt[*Command]      { return func(c *Command) { c.AddCommand(cmd...) } }
func Flags(o ...Opt[*FlagSet]) Opt[*Command] { return func(c *Command) { Apply(c.Flags(), o) } }
func PersistentPreRun(fn Fn) Opt[*Command]   { return func(c *Command) { c.PersistentPreRun = fn } }
func PersistentPreRunE(fn FnE) Opt[*Command] { return func(c *Command) { c.PersistentPreRunE = fn } }
func PersistentFlags(o ...Opt[*FlagSet]) Opt[*Command] {
	return func(c *Command) { Apply(c.PersistentFlags(), o) }
}
func Required(f ...string) Opt[*Command] {
	return func(c *Command) { lo.Must0(ApplyE(f, c.MarkFlagRequired)) }
}
func PersistentRequired(f ...string) Opt[*Command] {
	return func(c *Command) { lo.Must0(ApplyE(f, c.MarkPersistentFlagRequired)) }
}
