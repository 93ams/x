package flags

import (
	"github.com/spf13/pflag"
	. "github.com/tilau2328/cql/src/go/package/shared/x"
)

func String(v *string, name string, usage string, value string) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.StringVar(v, name, value, usage) }
}
func StringP(v *string, name string, sh string, usage string, value string) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.StringVarP(v, name, sh, value, usage) }
}
func StringSlice(v *[]string, name string, usage string, value ...string) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.StringSliceVar(v, name, value, usage) }
}
func StringSliceP(v *[]string, name string, sh string, usage string, value ...string) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.StringSliceVarP(v, name, sh, value, usage) }
}
func Uint16(v *uint16, name string, usage string, value uint16) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.Uint16Var(v, name, value, usage) }
}
func Uint16P(v *uint16, name string, sh string, usage string, value uint16) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.Uint16VarP(v, name, sh, value, usage) }
}
func Bool(v *bool, name string, usage string, value bool) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.BoolVar(v, name, value, usage) }
}
func BoolP(v *bool, name string, sh string, usage string, value bool) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.BoolVarP(v, name, sh, value, usage) }
}
func MapStringString(v *map[string]string, name string, usage string, value map[string]string) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.StringToStringVar(v, name, value, usage) }
}
func MapStringStringP(v *map[string]string, name string, sh string, usage string, value map[string]string) Opt[*pflag.FlagSet] {
	return func(set *pflag.FlagSet) { set.StringToStringVarP(v, name, sh, value, usage) }
}
