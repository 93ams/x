package driver

import (
	"fmt"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/adaptor/driver/restorer"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"os"
	"reflect"
)

type (
	FieldFilter func(name string, value reflect.Value) bool
	printer     struct {
		output io.Writer
		filter FieldFilter
		ptrmap map[interface{}]int
		indent int
		last   byte
		line   int
	}
	localError struct{ err error }
)

var indent = []byte(".  ")

func NotNilFilter(_ string, v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return !v.IsNil()
	}
	return true
}
func Write(w io.Writer, f *model.File) error {
	fset, af, err := RestoreFile(f)
	if err != nil {
		return err
	}
	return format.Node(w, fset, af)
}

func RestoreFile(file *model.File) (*token.FileSet, *ast.File, error) {
	r := restorer.NewRestorer()
	f, err := r.RestoreFile(file)
	if err != nil {
		return nil, nil, err
	}
	return r.Fset, f, nil
}
func Print(x interface{}) error                              { return Fprint(os.Stdout, x, NotNilFilter) }
func Fprint(w io.Writer, x interface{}, f FieldFilter) error { return fprint(w, x, f) }
func fprint(w io.Writer, x interface{}, f FieldFilter) (err error) {
	p := printer{
		output: w,
		filter: f,
		ptrmap: make(map[interface{}]int),
		last:   '\n',
	}
	defer func() {
		if e := recover(); e != nil {
			err = e.(localError).err
		}
	}()
	if x == nil {
		p.printf("nil\n")
		return
	}
	p.print(reflect.ValueOf(x))
	p.printf("\n")
	return
}
func (p *printer) Write(data []byte) (n int, err error) {
	var m int
	for i, b := range data {
		if b == '\n' {
			m, err = p.output.Write(data[n : i+1])
			n += m
			if err != nil {
				return
			}
			p.line++
		} else if p.last == '\n' {
			_, err = fmt.Fprintf(p.output, "%6d  ", p.line)
			if err != nil {
				return
			}
			for j := p.indent; j > 0; j-- {
				_, err = p.output.Write(indent)
				if err != nil {
					return
				}
			}
		}
		p.last = b
	}
	if len(data) > n {
		m, err = p.output.Write(data[n:])
		n += m
	}
	return
}
func (p *printer) printf(format string, args ...interface{}) {
	if _, err := fmt.Fprintf(p, format, args...); err != nil {
		panic(localError{err})
	}
}
func (p *printer) print(x reflect.Value) {
	if !NotNilFilter("", x) {
		p.printf("nil")
		return
	}
	switch x.Kind() {
	case reflect.Interface:
		p.print(x.Elem())
	case reflect.Map:
		p.printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for _, key := range x.MapKeys() {
				p.print(key)
				p.printf(": ")
				p.print(x.MapIndex(key))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")
	case reflect.Ptr:
		p.printf("*")
		ptr := x.Interface()
		if line, exists := p.ptrmap[ptr]; exists {
			p.printf("(obj @ %d)", line)
		} else {
			p.ptrmap[ptr] = p.line
			p.print(x.Elem())
		}
	case reflect.Array:
		p.printf("%s {", x.Type())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.printf("%d: ", i)
				p.print(x.Index(i))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")
	case reflect.Slice:
		if s, ok := x.Interface().([]byte); ok {
			p.printf("%#q", s)
			return
		}
		p.printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.printf("%d: ", i)
				p.print(x.Index(i))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")
	case reflect.Struct:
		t := x.Type()
		p.printf("%s {", t)
		p.indent++
		first := true
		for i, n := 0, t.NumField(); i < n; i++ {
			if name := t.Field(i).Name; model.IsExported(name) {
				value := x.Field(i)
				if p.filter == nil || p.filter(name, value) {
					if first {
						p.printf("\n")
						first = false
					}
					p.printf("%s: ", name)
					p.print(value)
					p.printf("\n")
				}
			}
		}
		p.indent--
		p.printf("}")
	default:
		v := x.Interface()
		switch v := v.(type) {
		case string:
			p.printf("%q", v)
			return
		}
		p.printf("%v", v)
	}
}
