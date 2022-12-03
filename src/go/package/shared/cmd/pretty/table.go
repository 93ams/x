package pretty

import (
	"github.com/jedib0t/go-pretty/v6/table"
	. "github.com/tilau2328/cql/src/go/package/shared/x"
	"io"
)

type Table struct {
	Header []any
	Footer []any
	Rows   [][]any
}

func (t Table) Write(w io.Writer) {
	tw := table.NewWriter()
	tw.SetOutputMirror(w)
	tw.AppendHeader(t.Header)
	for _, v := range t.Rows {
		tw.AppendRow(v)
	}
	tw.AppendFooter(t.Footer)
	tw.Render()
}

func NewTable(opt ...Opt[*Table]) *Table { return Apply(&Table{}, opt) }

func Header(v ...any) Opt[*Table] { return func(t *Table) { t.Header = v } }
func Footer(v ...any) Opt[*Table] { return func(t *Table) { t.Footer = v } }
func Rows(r ...[]any) Opt[*Table] { return func(t *Table) { t.Rows = r } }
