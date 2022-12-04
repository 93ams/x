package pretty

import (
	. "github.com/jedib0t/go-pretty/v6/table"
	. "github.com/tilau2328/cql/package/shared/x"
	"io"
)

type TablePretty struct {
	Header []any
	Footer []any
	Rows   [][]any
}

func (t TablePretty) Write(w io.Writer) {
	tw := NewWriter()
	tw.SetOutputMirror(w)
	tw.AppendHeader(t.Header)
	for _, v := range t.Rows {
		tw.AppendRow(v)
	}
	tw.AppendFooter(t.Footer)
	tw.Render()
}

func NewTable(opt ...Opt[*TablePretty]) *TablePretty { return Apply(&TablePretty{}, opt) }

func Header(v ...any) Opt[*TablePretty] { return func(t *TablePretty) { t.Header = v } }
func Footer(v ...any) Opt[*TablePretty] { return func(t *TablePretty) { t.Footer = v } }
func Rows(r ...[]any) Opt[*TablePretty] { return func(t *TablePretty) { t.Rows = r } }
