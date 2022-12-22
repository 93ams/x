package printer

import (
	"bufio"
	"fmt"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"io"
	"strings"
)

type encoder interface{ Encode(int, io.Writer) error }

var Indent = "\t"

func (p *Printer) multiLineComment(indent int, dst io.Writer, s string) error {
	scanner := bufio.NewScanner(strings.NewReader(s))
	i := strings.Repeat(Indent, indent)
	for scanner.Scan() {
		if _, err := fmt.Fprintf(dst, " %s// %s", i, strings.Replace(s, "\n", " ", -1)); err != nil {
			return err
		}
	}
	return nil
}
func (p *Printer) comment(dst io.Writer, s string) error {
	_, err := fmt.Fprintf(dst, " // %s", strings.Replace(s, "\n", " ", -1))
	return err
}
func (p *Printer) LiteralField(dst io.Writer, e model.LiteralField) error {
	if _, err := fmt.Fprintf(dst, "%s: ", e.name); err != nil {
		return err
	} else if f, ok := e.value.(encoder); ok {
		if err := f.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode option value for message literal %q: %w`, e.name, err)
		}
	} else if _, err := fmt.Fprintf(dst, "%#v", e.value); err != nil {
		return err
	}
	return nil
}
func (p *Printer) EnumElement(dst io.Writer, e model.EnumElement) error {
	if _, err := fmt.Fprintf(dst, "\n%s%s = %d;", strings.Repeat(Indent, indent), e.name, e.value); err != nil {
		return err
	}
	if s := e.comment; s != "" {
		if err := singleLineComment(dst, s); err != nil {
			return err
		}
	}
	return nil
}
func (p *Printer) Extension(dst io.Writer, e model.Extension) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%sextend %s {", i, e.name); err != nil {
		return err
	}
	for i, v := range e.fields {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode field declaration %d for extension %q: %w`, i, e.name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (p *Printer) Message(dst io.Writer, e model2.Message) error {
	i := strings.Repeat(Indent, indent)
	if c := e.comment; c != "" {
		if err := multiLineComment(indent, dst, c); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%smessage %s {", i, e.name); err != nil {
		return err
	}
	for i, v := range e.oneOfs {
		if err := p.OneOf(dst, v); err != nil {
			return fmt.Errorf(`failed to encode nested oneof declaration %d for message %q: %w`, i, e.name, err)
		}
	}
	for i, v := range e.extensions {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested extension declaration %d for message %q: %w`, i, e.name, err)
		}
	}
	for i, v := range e.options {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested option declaration %d for message %q: %w`, i, e.name, err)
		}
	}
	for i, v := range e.enums {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested enum declaration %d for message %q: %w`, i, e.name, err)
		}
	}
	for i, v := range e.messages {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested message declaration %d for message %q: %w`, i, e.name, err)
		} else if _, err := fmt.Fprint(dst, "\n"); err != nil {
			return err
		}
	}
	for i, v := range e.fields {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode field declaration %d for message %q: %w`, i, e.name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (p *Printer) Literal(dst io.Writer, e model.Literal) error {
	ind := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "%s{", ind); err != nil {
		return err
	}
	for i, field := range e.fields {
		if !e.singleLine {
			if _, err := fmt.Fprintf(dst, "\n%s", ind); err != nil {
				return err
			}
		} else if i > 0 {
			if _, err := fmt.Fprintf(dst, " "); err != nil {
				return err
			}
		}
		if err := field.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode field %d for message literal: %w`, i, err)
		}
	}
	if !e.singleLine {
		if _, err := fmt.Fprintf(dst, "\n%s", ind); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(dst, "}"); err != nil {
		return err
	}
	return nil
}
func (p *Printer) Service(dst io.Writer, e model2.Service) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%sservice %s {", i, e.name); err != nil {
		return err
	}
	for i, v := range e.methods {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode method %d for service %q: %w`, i, e.name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (p *Printer) Option(dst io.Writer, e model2.Option) error {
	if e.compact {
		if _, err := fmt.Fprintf(dst, "%s = %s", e.name, e.value); err != nil {
			return err
		}
		return nil
	}
	if _, err := fmt.Fprintf(dst, "\n%soption %s = ", strings.Repeat(Indent, indent), e.name); err != nil {
		return err
	}
	if v, ok := e.value.(encoder); ok {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode option value for option %q: %w`, e.name, err)
		}
	} else if _, err := fmt.Fprintf(dst, "%#v", e.value); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(dst, ";"); err != nil {
		return err
	}
	return nil
}
func (p *Printer) Method(dst io.Writer, e model.Method) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%srpc %s(%s) returns (%s)", i, e.name, e.input, e.output); err != nil {
		return err
	}
	if options := e.options; len(options) > 0 {
		if _, err := fmt.Fprintf(dst, " {"); err != nil {
			return err
		}
		for i, option := range options {
			if err := option.Encode(indent+1, dst); err != nil {
				return fmt.Errorf(`failed to encode option %d for method %q: %w`, i, e.name, err)
			}
		}
		if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(dst, ";"); err != nil {
		return err
	}
	return nil
}
func (p *Printer) Import(dst io.Writer, e model2.Import) error {
	if _, err := fmt.Fprintf(dst, "\n%simport", strings.Repeat(Indent, indent)); err != nil {
		return err
	}
	switch e.typ {
	case ImportPublic:
		if _, err := fmt.Fprintf(dst, " public"); err != nil {
			return err
		}
	case ImportWeak:
		if _, err := fmt.Fprintf(dst, " weak"); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(dst, " %q;", e.path); err != nil {
		return err
	}
	return nil
}
func (p *Printer) OneOf(dst io.Writer, e model2.OneOf) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%soneof %s {", i, e.name); err != nil {
		return err
	}
	for i, v := range e.fields {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode field declaration %d for oneof %q: %w`, i, e.name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (p *Printer) Field(dst io.Writer, e model2.Field) error {
	if _, err := fmt.Fprintf(dst, "\n%s", strings.Repeat(Indent, indent)); err != nil {
		return err
	}
	var err error
	switch e.cardinality {
	case CardinalityRequired:
		_, err = fmt.Fprintf(dst, "required ")
	case CardinalityOptional:
		_, err = fmt.Fprintf(dst, "optional ")
	case CardinalityRepeated:
		_, err = fmt.Fprintf(dst, "repeated ")
	}
	if err != nil {
		return err
	} else if _, err := fmt.Fprintf(dst, "%s %s = %d", e.typ, e.name, e.id); err != nil {
		return err
	}
	if options := e.options; len(options) > 0 {
		if _, err := fmt.Fprintf(dst, " ["); err != nil {
			return err
		}
		for i, option := range options {
			if err := option.Encode(indent+1, dst); err != nil {
				return fmt.Errorf(`failed to encode option %d for field %q: %w`, i, e.name, err)
			}
		}
		if _, err := fmt.Fprintf(dst, "]"); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(dst, ";"); err != nil {
		return err
	}
	if c := e.comment; c != "" {
		if err := singleLineComment(dst, c); err != nil {
			return err
		}
	}
	return nil
}
func (p *Printer) Enum(dst io.Writer, e model2.Enum) error {
	if s := e.comment; s != "" {
		if err := multiLineComment(indent, dst, s); err != nil {
			return err
		}
	}
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%senum %s {", i, e.name); err != nil {
		return err
	}
	for i, v := range e.elements {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode enum declaration %d for enum %q: %w`, i, e.name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (p *Printer) File(dst io.Writer, e model.File) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "%ssyntax = \"proto3\";", i); err != nil {
		return err
	} else if _, err := fmt.Fprintf(dst, "\n\n%spackage %s;", i, e.pkg); err != nil {
		return err
	}
	if list := e.imports; len(list) > 0 {
		if _, err := fmt.Fprint(dst, "\n"); err != nil {
			return err
		}
		for i, v := range list {
			if err := v.Encode(indent, dst); err != nil {
				return fmt.Errorf(`failed to encode import statement %d: %w`, i, err)
			}
		}
	}
	if list := e.options; len(list) > 0 {
		if _, err := fmt.Fprint(dst, "\n"); err != nil {
			return err
		}
		for i, v := range list {
			if err := v.Encode(indent, dst); err != nil {
				return fmt.Errorf(`failed to encode option declaration %d: %w`, i, err)
			}
		}
	}
	for i, v := range e.extensions {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode extension declaration %d: %w`, i, err)
		}
	}
	for i, v := range e.messages {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode message declaration %d: %w`, i, err)
		}
	}
	for i, v := range e.enums {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode enum declaration %d: %w`, i, err)
		}
	}
	for i, v := range e.services {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode service declaration %d: %w`, i, err)
		}
	}
	return nil
}
