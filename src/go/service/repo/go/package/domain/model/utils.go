package model

import (
	"github.com/samber/lo"
	"go/token"
)

func Names(names []string) []*Ident {
	return lo.Map(names, func(item string, index int) *Ident {
		return &Ident{Name: item}
	})
}
func SetRecv(recv *FieldList, fn []*Func) []Decl {
	return lo.Map(fn, func(item *Func, _ int) Decl {
		item = item.Clone()
		item.Recv = recv
		return item
	})
}
func RemoveRecv(fn []*Func) []Decl {
	return lo.Map(fn, func(item *Func, _ int) Decl {
		item = item.Clone()
		item.Recv = nil
		return item
	})
}

func NewType(name string, t Expr) *Gen {
	return &Gen{Tok: token.TYPE, Specs: []Spec{
		&Type{Name: &Ident{Name: name}, Type: t.AsExpr()},
	}}
}
func NewStruct(name string, fields []*Field) *Gen {
	return NewType(name, &Struct{Fields: &FieldList{List: fields}})
}
func NewInterface(name string, fields []*Field) *Gen {
	return NewType(name, &Interface{Methods: &FieldList{List: fields}})
}
func NewFunc(name string, t *FuncType, body *Block) *Func {
	return &Func{Name: &Ident{Name: name}, Type: t, Body: body}
}
func NewKeyValue(key, val Expr) *KeyValue {
	return &KeyValue{
		Key:   key.AsExpr(),
		Value: val.AsExpr(),
		Decs: KeyValueDecs{
			NodeDecs: NodeDecs{
				Before: NewLine,
				After:  NewLine,
			},
		},
	}
}
