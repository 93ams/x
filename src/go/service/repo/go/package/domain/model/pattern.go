package model

import (
	"github.com/samber/lo"
	"go/token"
)

type (
	Constructor struct {
		Name         string
		Result       Expr
		Dependencies []*Field
		Assignments  []*KeyValue
	}
	Mapper struct {
		Name        string
		From, To    Expr
		Assignments []*KeyValue
	}
	Builder struct {
		Fields  []*Field
		Methods []*Func
		Name    string
		Type    *Ident
	}
	Service struct {
		Name         string
		Methods      []*Func
		Dependencies []*Field
	}
	Enum struct {
		Type   *Ident
		Values map[string]string
	}
)

func (c Constructor) Func() *Func {
	return NewFunc(c.Name, &FuncType{
		Results: &FieldList{
			List: []*Field{{
				Type: &Star{
					X: c.Result.AsExpr(),
				},
			}},
		},
	},
		&Block{List: []Stmt{
			&Return{
				Results: []Expr{
					&Unary{Op: token.AND, X: &CompositeLit{
						Type: c.Result.AsExpr(),
						Elts: lo.Map(c.Assignments, func(item *KeyValue, _ int) Expr {
							return item.Clone()
						}),
					}},
				},
				Decs: ReturnDecs{
					NodeDecs: NodeDecs{
						Before: NewLine,
					},
				},
			},
		},
		})
}
func (m Mapper) Func() *Func {
	return NewFunc(m.Name, &FuncType{
		Params: &FieldList{
			List: []*Field{{
				Names: []*Ident{{Name: "in"}},
				Type:  m.From.AsExpr(),
			}},
		},
		Results: &FieldList{
			List: []*Field{{
				Type: m.To.AsExpr(),
			}},
		},
	}, &Block{
		List: []Stmt{
			&Return{
				Results: []Expr{
					&CompositeLit{
						Type: m.To.AsExpr(),
						Elts: lo.Map(m.Assignments, func(item *KeyValue, _ int) Expr {
							return NewKeyValue(item.Key, &Selector{X: &Ident{Name: "in"}, Sel: item.Value.(*Ident)})
						}),
					},
				},
				Decs: ReturnDecs{
					NodeDecs: NodeDecs{
						Before: NewLine,
						After:  NewLine,
					},
				},
			},
		},
	})
}
func (b Builder) Decls() []Decl {
	return append(append([]Decl{
		NewType(b.Name, &Index{
			X: &Selector{
				Sel: &Ident{Name: "Builder"},
				X:   &Ident{Name: "x"},
			},
			Index: &Selector{
				Sel: &Ident{Name: b.Type.Name},
				X:   &Ident{Name: b.Type.Path},
			},
		}),
		Constructor{Name: "New" + b.Name, Result: b.Type}.Func(),
	}, lo.Map(b.Fields, func(item *Field, _ int) Decl {
		return &Func{}
	})...), lo.Map(b.Methods, func(item *Func, _ int) Decl {
		return &Func{}
	})...)
}
func (s Service) Decls() []Decl {
	var ret []Decl
	if len(s.Dependencies) > 0 {
		deps := &Field{Names: []*Ident{{Name: "opts"}}, Type: &Ident{Name: s.Name + "Options"}}
		ret = []Decl{
			NewStruct(s.Name+"Options", s.Dependencies),
			NewStruct(s.Name, []*Field{deps}),
			Constructor{Name: "New" + s.Name, Result: &Ident{Name: s.Name}, Dependencies: []*Field{deps}}.Func(),
		}
	} else {
		ret = []Decl{
			NewStruct(s.Name, nil),
			Constructor{Name: "New" + s.Name, Result: &Ident{Name: s.Name}}.Func(),
		}
	}
	return append(ret, SetRecv(&FieldList{List: []*Field{{Type: &Star{X: &Ident{Name: s.Name}}}}}, s.Methods)...)
}
func (e Enum) Decl() []Decl {
	return nil
}
