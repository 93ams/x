package cmd

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	. "github.com/spf13/pflag"
	. "github.com/tilau2328/x/src/go/package/x"
)

type Fn func(*cobra.Command, []string)
type FnE func(*cobra.Command, []string) error

func New(opts ...Opt[*cobra.Command]) *cobra.Command { return Apply(&cobra.Command{}, opts) }
func Run(fn Fn) Opt[*cobra.Command]                  { return func(c *cobra.Command) { c.Run = fn } }
func RunE(fn FnE) Opt[*cobra.Command]                { return func(c *cobra.Command) { c.RunE = fn } }
func Use(use string) Opt[*cobra.Command]             { return func(c *cobra.Command) { c.Use = use } }
func Alias(a ...string) Opt[*cobra.Command]          { return func(c *cobra.Command) { c.Aliases = a } }
func Add(cmd ...*cobra.Command) Opt[*cobra.Command] {
	return func(c *cobra.Command) { c.AddCommand(cmd...) }
}
func Flags(o ...Opt[*FlagSet]) Opt[*cobra.Command] {
	return func(c *cobra.Command) { Apply(c.Flags(), o) }
}
func PersistentPreRun(fn Fn) Opt[*cobra.Command] {
	return func(c *cobra.Command) { c.PersistentPreRun = fn }
}
func PersistentPreRunE(fn FnE) Opt[*cobra.Command] {
	return func(c *cobra.Command) { c.PersistentPreRunE = fn }
}
func PersistentFlags(o ...Opt[*FlagSet]) Opt[*cobra.Command] {
	return func(c *cobra.Command) { Apply(c.PersistentFlags(), o) }
}
func Required(f ...string) Opt[*cobra.Command] {
	return func(c *cobra.Command) { lo.Must0(ApplyE(f, c.MarkFlagRequired)) }
}
func PersistentRequired(f ...string) Opt[*cobra.Command] {
	return func(c *cobra.Command) { lo.Must0(ApplyE(f, c.MarkPersistentFlagRequired)) }
}
