package proto

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type encoder interface{ Encode(int, io.Writer) error }

var Indent = "\t"

func multiLineComment(indent int, dst io.Writer, s string) error {
	scanner := bufio.NewScanner(strings.NewReader(s))
	i := strings.Repeat(Indent, indent)
	for scanner.Scan() {
		if _, err := fmt.Fprintf(dst, " %s// %s", i, strings.Replace(s, "\n", " ", -1)); err != nil {
			return err
		}
	}
	return nil
}
func singleLineComment(dst io.Writer, s string) error {
	_, err := fmt.Fprintf(dst, " // %s", strings.Replace(s, "\n", " ", -1))
	return err
}
func (e _enumElement) Encode(indent int, dst io.Writer) error {
	if _, err := fmt.Fprintf(dst, "\n%s%s = %d;", strings.Repeat(Indent, indent), e.Name, e.Value); err != nil {
		return err
	}
	if s := e.Comment; s != "" {
		if err := singleLineComment(dst, s); err != nil {
			return err
		}
	}
	return nil
}
func (e _enum) Encode(indent int, dst io.Writer) error {
	if s := e.Comment; s != "" {
		if err := multiLineComment(indent, dst, s); err != nil {
			return err
		}
	}
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%senum %s {", i, e.Name); err != nil {
		return err
	}
	for i, v := range e.Elements {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode enum declaration %d for enum %q: %w`, i, e.Name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (e _oneOf) Encode(indent int, dst io.Writer) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%soneof %s {", i, e.Name); err != nil {
		return err
	}
	for i, v := range e.Fields {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode field declaration %d for oneof %q: %w`, i, e.Name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (e _message) Encode(indent int, dst io.Writer) error {
	i := strings.Repeat(Indent, indent)
	if c := e.Comment; c != "" {
		if err := multiLineComment(indent, dst, c); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%smessage %s {", i, e.Name); err != nil {
		return err
	}
	for i, v := range e.OneOfs {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested oneof declaration %d for message %q: %w`, i, e.Name, err)
		}
	}
	for i, v := range e.Extensions {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested extension declaration %d for message %q: %w`, i, e.Name, err)
		}
	}
	for i, v := range e.Options {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested option declaration %d for message %q: %w`, i, e.Name, err)
		}
	}
	for i, v := range e.Enums {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested enum declaration %d for message %q: %w`, i, e.Name, err)
		}
	}
	for i, v := range e.Messages {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode nested message declaration %d for message %q: %w`, i, e.Name, err)
		} else if _, err := fmt.Fprint(dst, "\n"); err != nil {
			return err
		}
	}
	for i, v := range e.Fields {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode field declaration %d for message %q: %w`, i, e.Name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (e _literal) Encode(indent int, dst io.Writer) error {
	ind := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "%s{", ind); err != nil {
		return err
	}
	for i, field := range e.Fields {
		if !e.SingleLine {
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
	if !e.SingleLine {
		if _, err := fmt.Fprintf(dst, "\n%s", ind); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(dst, "}"); err != nil {
		return err
	}
	return nil
}
func (e _literalField) Encode(indent int, dst io.Writer) error {
	if _, err := fmt.Fprintf(dst, "%s: ", e.Name); err != nil {
		return err
	} else if f, ok := e.Value.(encoder); ok {
		if err := f.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode option value for message literal %q: %w`, e.Name, err)
		}
	} else if _, err := fmt.Fprintf(dst, "%#v", e.Value); err != nil {
		return err
	}
	return nil
}
func (e _option) Encode(indent int, dst io.Writer) error {
	if e.Compact {
		if _, err := fmt.Fprintf(dst, "%s = %s", e.Name, e.Value); err != nil {
			return err
		}
		return nil
	}
	if _, err := fmt.Fprintf(dst, "\n%soption %s = ", strings.Repeat(Indent, indent), e.Name); err != nil {
		return err
	}
	if v, ok := e.Value.(encoder); ok {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode option value for option %q: %w`, e.Name, err)
		}
	} else if _, err := fmt.Fprintf(dst, "%#v", e.Value); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(dst, ";"); err != nil {
		return err
	}
	return nil
}
func (e _extension) Encode(indent int, dst io.Writer) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%sextend %s {", i, e.Name); err != nil {
		return err
	}
	for i, v := range e.Fields {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode field declaration %d for extension %q: %w`, i, e.Name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (e _service) Encode(indent int, dst io.Writer) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%sservice %s {", i, e.Name); err != nil {
		return err
	}
	for i, v := range e.Methods {
		if err := v.Encode(indent+1, dst); err != nil {
			return fmt.Errorf(`failed to encode method %d for service %q: %w`, i, e.Name, err)
		}
	}
	if _, err := fmt.Fprintf(dst, "\n%s}", i); err != nil {
		return err
	}
	return nil
}
func (e _method) Encode(indent int, dst io.Writer) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "\n%srpc %s(%s) returns (%s)", i, e.Name, e.Input, e.Output); err != nil {
		return err
	}
	if options := e.Options; len(options) > 0 {
		if _, err := fmt.Fprintf(dst, " {"); err != nil {
			return err
		}
		for i, option := range options {
			if err := option.Encode(indent+1, dst); err != nil {
				return fmt.Errorf(`failed to encode option %d for method %q: %w`, i, e.Name, err)
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
func (e _import) Encode(indent int, dst io.Writer) error {
	if _, err := fmt.Fprintf(dst, "\n%simport", strings.Repeat(Indent, indent)); err != nil {
		return err
	}
	switch e.Type {
	case ImportPublic:
		if _, err := fmt.Fprintf(dst, " public"); err != nil {
			return err
		}
	case ImportWeak:
		if _, err := fmt.Fprintf(dst, " weak"); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(dst, " %q;", e.Path); err != nil {
		return err
	}
	return nil
}
func (e _file) Encode(indent int, dst io.Writer) error {
	i := strings.Repeat(Indent, indent)
	if _, err := fmt.Fprintf(dst, "%ssyntax = \"proto3\";", i); err != nil {
		return err
	} else if _, err := fmt.Fprintf(dst, "\n\n%spackage %s;", i, e.Package); err != nil {
		return err
	}
	if list := e.Imports; len(list) > 0 {
		if _, err := fmt.Fprint(dst, "\n"); err != nil {
			return err
		}
		for i, v := range list {
			if err := v.Encode(indent, dst); err != nil {
				return fmt.Errorf(`failed to encode import statement %d: %w`, i, err)
			}
		}
	}
	if list := e.Options; len(list) > 0 {
		if _, err := fmt.Fprint(dst, "\n"); err != nil {
			return err
		}
		for i, v := range list {
			if err := v.Encode(indent, dst); err != nil {
				return fmt.Errorf(`failed to encode option declaration %d: %w`, i, err)
			}
		}
	}
	for i, v := range e.Extensions {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode extension declaration %d: %w`, i, err)
		}
	}
	for i, v := range e.Messages {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode message declaration %d: %w`, i, err)
		}
	}
	for i, v := range e.Enums {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode enum declaration %d: %w`, i, err)
		}
	}
	for i, v := range e.Services {
		if _, err := fmt.Fprintf(dst, "\n"); err != nil {
			return err
		}
		if err := v.Encode(indent, dst); err != nil {
			return fmt.Errorf(`failed to encode service declaration %d: %w`, i, err)
		}
	}
	return nil
}
func (e _field) Encode(indent int, dst io.Writer) error {
	if _, err := fmt.Fprintf(dst, "\n%s", strings.Repeat(Indent, indent)); err != nil {
		return err
	}
	var err error
	switch e.Cardinality {
	case CardinalityRequired:
		_, err = fmt.Fprintf(dst, "required ")
	case CardinalityOptional:
		_, err = fmt.Fprintf(dst, "optional ")
	case CardinalityRepeated:
		_, err = fmt.Fprintf(dst, "repeated ")
	}
	if err != nil {
		return err
	} else if _, err := fmt.Fprintf(dst, "%s %s = %d", e.Type, e.Name, e.ID); err != nil {
		return err
	}
	if options := e.Options; len(options) > 0 {
		if _, err := fmt.Fprintf(dst, " ["); err != nil {
			return err
		}
		for i, option := range options {
			if err := option.Encode(indent+1, dst); err != nil {
				return fmt.Errorf(`failed to encode option %d for field %q: %w`, i, e.Name, err)
			}
		}
		if _, err := fmt.Fprintf(dst, "]"); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintf(dst, ";"); err != nil {
		return err
	}
	if c := e.Comment; c != "" {
		if err := singleLineComment(dst, c); err != nil {
			return err
		}
	}
	return nil
}
