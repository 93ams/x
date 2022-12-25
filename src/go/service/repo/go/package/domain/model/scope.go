package model

import (
	"bytes"
	"fmt"
	"go/ast"
)

type (
	Map struct {
		Ast AstMap
		Dst DstMap
	}
	AstMap struct {
		Nodes   map[Node]ast.Node
		Objects map[*Object]*ast.Object
		Scopes  map[*Scope]*ast.Scope
	}
	DstMap struct {
		Nodes   map[ast.Node]Node
		Objects map[*ast.Object]*Object
		Scopes  map[*ast.Scope]*Scope
	}
	Scope struct {
		Outer   *Scope
		Objects map[string]*Object
	}
	ObjKind int
	Object  struct {
		Kind ObjKind
		Name string
		Decl interface{}
		Data interface{}
		Type interface{}
	}
)

const (
	Bad ObjKind = iota
	Pkg
	Con
	Typ
	Var
	Fun
	Lbl
)

var objKindStrings = [...]string{
	Bad: "bad",
	Pkg: "package",
	Con: "const",
	Typ: "type",
	Var: "var",
	Fun: "func",
	Lbl: "label",
}

func NewObj(kind ObjKind, name string) *Object { return &Object{Kind: kind, Name: name} }
func (kind ObjKind) String() string            { return objKindStrings[kind] }
func NewScope(outer *Scope) *Scope             { return &Scope{outer, make(map[string]*Object, 4)} }
func (s *Scope) Lookup(name string) *Object    { return s.Objects[name] }
func (s *Scope) Insert(obj *Object) (alt *Object) {
	if alt = s.Objects[obj.Name]; alt == nil {
		s.Objects[obj.Name] = obj
	}
	return
}
func (s *Scope) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "scope %p {", s)
	if s != nil && len(s.Objects) > 0 {
		for _, obj := range s.Objects {
			fmt.Fprintf(&buf, "\t%s %s\n", obj.Kind, obj.Name)
		}
	}
	fmt.Fprintf(&buf, "}\n")
	return buf.String()
}
func NewMap() Map {
	return Map{
		Ast: AstMap{
			Nodes:   map[Node]ast.Node{},
			Scopes:  map[*Scope]*ast.Scope{},
			Objects: map[*Object]*ast.Object{},
		},
		Dst: DstMap{
			Nodes:   map[ast.Node]Node{},
			Scopes:  map[*ast.Scope]*Scope{},
			Objects: map[*ast.Object]*Object{},
		},
	}
}
