package model

import "fmt"

type (
	Visitor   interface{ Visit(node Node) (w Visitor) }
	inspector func(Node) bool
)

func walk[T Node](v Visitor, list []T) {
	for _, x := range list {
		Walk(v, x)
	}
}

func Walk(v Visitor, node Node) {
	if v == nil || node == nil {
		return
	}
	if v = v.Visit(node); v == nil {
		return
	}
	if v == nil || node == nil {
		return
	}
	switch n := node.(type) {
	case *Field:
		walk(v, n.Names)
		Walk(v, n.Type)
		if n.Tag != nil {
			Walk(v, n.Tag)
		}
	case *FieldList:
		for _, f := range n.List {
			Walk(v, f)
		}
	case *BadExpr, *Ident, *Lit:
	case *Ellipsis:
		if n.Elt != nil {
			Walk(v, n.Elt)
		}
	case *FuncLit:
		Walk(v, n.Type)
		Walk(v, n.Body)
	case *CompositeLit:
		if n.Type != nil {
			Walk(v, n.Type)
		}
		walk(v, n.Elts)
	case *Paren:
		Walk(v, n.X)
	case *Selector:
		Walk(v, n.X)
		Walk(v, n.Sel)
	case *Index:
		Walk(v, n.X)
		Walk(v, n.Index)
	case *IndexList:
		Walk(v, n.X)
		walk(v, n.Indices)
	case *Slice:
		Walk(v, n.X)
		if n.Low != nil {
			Walk(v, n.Low)
		}
		if n.High != nil {
			Walk(v, n.High)
		}
		if n.Max != nil {
			Walk(v, n.Max)
		}
	case *TypeAssert:
		Walk(v, n.X)
		if n.Type != nil {
			Walk(v, n.Type)
		}
	case *Call:
		Walk(v, n.Fun)
		walk(v, n.Args)
	case *Star:
		Walk(v, n.X)
	case *Unary:
		Walk(v, n.X)
	case *Binary:
		Walk(v, n.X)
		Walk(v, n.Y)
	case *KeyValue:
		Walk(v, n.Key)
		Walk(v, n.Value)
	case *Array:
		if n.Len != nil {
			Walk(v, n.Len)
		}
		Walk(v, n.Elt)
	case *Struct:
		Walk(v, n.Fields)
	case *FuncType:
		if n.TypeParams != nil {
			Walk(v, n.TypeParams)
		}
		if n.Params != nil {
			Walk(v, n.Params)
		}
		if n.Results != nil {
			Walk(v, n.Results)
		}
	case *Interface:
		Walk(v, n.Methods)
	case *MapType:
		Walk(v, n.Key)
		Walk(v, n.Value)
	case *Chan:
		Walk(v, n.Value)
	case *BadStmt:
	case *DeclStmt:
		Walk(v, n.Decl)
	case *Empty:
	case *Labeled:
		Walk(v, n.Label)
		Walk(v, n.Stmt)
	case *ExprStmt:
		Walk(v, n.X)
	case *Send:
		Walk(v, n.Chan)
		Walk(v, n.Value)
	case *IncDec:
		Walk(v, n.X)
	case *Assign:
		walk(v, n.Lhs)
		walk(v, n.Rhs)
	case *Go:
		Walk(v, n.Call)
	case *Defer:
		Walk(v, n.Call)
	case *Return:
		walk(v, n.Results)
	case *Branch:
		if n.Label != nil {
			Walk(v, n.Label)
		}
	case *Block:
		walk(v, n.List)
	case *If:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		Walk(v, n.Cond)
		Walk(v, n.Body)
		if n.Else != nil {
			Walk(v, n.Else)
		}
	case *CaseClause:
		walk(v, n.List)
		walk(v, n.Body)
	case *Switch:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		if n.Tag != nil {
			Walk(v, n.Tag)
		}
		Walk(v, n.Body)
	case *TypeSwitch:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		Walk(v, n.Assign)
		Walk(v, n.Body)
	case *CommClause:
		if n.Comm != nil {
			Walk(v, n.Comm)
		}
		walk(v, n.Body)
	case *Select:
		Walk(v, n.Body)
	case *For:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		if n.Cond != nil {
			Walk(v, n.Cond)
		}
		if n.Post != nil {
			Walk(v, n.Post)
		}
		Walk(v, n.Body)
	case *Range:
		if n.Key != nil {
			Walk(v, n.Key)
		}
		if n.Value != nil {
			Walk(v, n.Value)
		}
		Walk(v, n.X)
		Walk(v, n.Body)
	case *Import:
		if n.Name != nil {
			Walk(v, n.Name)
		}
		Walk(v, n.Path)
	case *Value:
		walk(v, n.Names)
		if n.Type != nil {
			Walk(v, n.Type)
		}
		walk(v, n.Values)
	case *Type:
		Walk(v, n.Name)
		if n.Params != nil {
			Walk(v, n.Params)
		}
		Walk(v, n.Type)
	case *BadDecl:
	case *Gen:
		for _, s := range n.Specs {
			Walk(v, s)
		}
	case *Func:
		if n.Recv != nil {
			Walk(v, n.Recv)
		}
		Walk(v, n.Name)
		Walk(v, n.Type)
		if n.Body != nil {
			Walk(v, n.Body)
		}
	case *File:
		Walk(v, n.Name)
		walk(v, n.Decls)
	case *Package:
		for _, f := range n.Files {
			Walk(v, f)
		}
	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}
}

func (f inspector) Visit(node Node) Visitor {
	if f(node) {
		return f
	}
	return nil
}
func Inspect(node Node, f func(Node) bool) { Walk(inspector(f), node) }
