package model

import (
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
)

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
