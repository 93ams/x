package printer

import "io"

type Printer struct {
	Ident     string
	identSize int
}

func NewPrinter(ident string) *Printer {
	return &Printer{Ident: ident}
}
func (p *Printer) ident(at io.Writer) {

}
