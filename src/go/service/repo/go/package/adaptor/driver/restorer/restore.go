package restorer

import (
	"fmt"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/adaptor/model"
	"go/ast"
	"go/token"
)

func (r *FileRestorer) restoreNode(n model.Node, parentName, parentField, parentFieldType string, allowDuplicate bool) ast.Node {
	if an, ok := r.Ast.Nodes[n]; ok {
		if allowDuplicate {
			return an
		} else {
			panic(fmt.Sprintf("duplicate node: %#v", n))
		}
	}
	switch n := n.(type) {
	case *model.ArrayType:
		out := &ast.ArrayType{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecorations(out, "Lbrack", n.Decs.Lbrack, false)
		if n.Len != nil {
			out.Len = r.restoreNode(n.Len, "ArrayType", "Len", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecorations(out, "Len", n.Decs.Len, false)
		if n.Elt != nil {
			out.Elt = r.restoreNode(n.Elt, "ArrayType", "Elt", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.AssignStmt:
		out := &ast.AssignStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		for _, v := range n.Lhs {
			out.Lhs = append(out.Lhs, r.restoreNode(v, "AssignStmt", "Lhs", "Expr", allowDuplicate).(ast.Expr))
		}
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applyDecorations(out, "Tok", n.Decs.Tok, false)
		for _, v := range n.Rhs {
			out.Rhs = append(out.Rhs, r.restoreNode(v, "AssignStmt", "Rhs", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.BadDecl:
		out := &ast.BadDecl{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.From = r.cursor
		r.cursor += token.Pos(n.Length)
		out.To = r.cursor
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.BadExpr:
		out := &ast.BadExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.From = r.cursor
		r.cursor += token.Pos(n.Length)
		out.To = r.cursor
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.BadStmt:
		out := &ast.BadStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.From = r.cursor
		r.cursor += token.Pos(n.Length)
		out.To = r.cursor
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.BasicLit:
		out := &ast.BasicLit{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		r.applyLiteral(n.Value)
		out.ValuePos = r.cursor
		out.Value = n.Value
		r.cursor += token.Pos(len(n.Value))
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Kind = n.Kind
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.BinaryExpr:
		out := &ast.BinaryExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "BinaryExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.Op = n.Op
		out.OpPos = r.cursor
		r.cursor += token.Pos(len(n.Op.String()))
		r.applyDecorations(out, "Op", n.Decs.Op, false)
		if n.Y != nil {
			out.Y = r.restoreNode(n.Y, "BinaryExpr", "Y", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.BlockStmt:
		out := &ast.BlockStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Lbrace = r.cursor
		r.cursor += token.Pos(len(token.LBRACE.String()))
		r.applyDecorations(out, "Lbrace", n.Decs.Lbrace, false)
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v, "BlockStmt", "List", "Stmt", allowDuplicate).(ast.Stmt))
		}
		if n.RbraceHasNoPos {
			out.Rbrace = token.NoPos
		} else {
			out.Rbrace = r.cursor
		}
		r.cursor += token.Pos(len(token.RBRACE.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.BranchStmt:
		out := &ast.BranchStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applyDecorations(out, "Tok", n.Decs.Tok, false)
		if n.Label != nil {
			out.Label = r.restoreNode(n.Label, "BranchStmt", "Label", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.CallExpr:
		out := &ast.CallExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Fun != nil {
			out.Fun = r.restoreNode(n.Fun, "CallExpr", "Fun", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Fun", n.Decs.Fun, false)
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))
		r.applyDecorations(out, "Lparen", n.Decs.Lparen, false)
		for _, v := range n.Args {
			out.Args = append(out.Args, r.restoreNode(v, "CallExpr", "Args", "Expr", allowDuplicate).(ast.Expr))
		}
		if n.Ellipsis {
			out.Ellipsis = r.cursor
			r.cursor += token.Pos(len(token.ELLIPSIS.String()))
		}
		r.applyDecorations(out, "Ellipsis", n.Decs.Ellipsis, false)
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.CaseClause:
		out := &ast.CaseClause{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Case = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.List == nil {
				return token.DEFAULT
			}
			return token.CASE
		}().String()))
		r.applyDecorations(out, "Case", n.Decs.Case, false)
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v, "CaseClause", "List", "Expr", allowDuplicate).(ast.Expr))
		}
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecorations(out, "Colon", n.Decs.Colon, false)
		for _, v := range n.Body {
			out.Body = append(out.Body, r.restoreNode(v, "CaseClause", "Body", "Stmt", allowDuplicate).(ast.Stmt))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.ChanType:
		out := &ast.ChanType{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Begin = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.Dir == model.RECV {
				return token.ARROW
			}
			return token.CHAN
		}().String()))
		if n.Dir == model.RECV {
			r.cursor += token.Pos(len(token.CHAN.String()))
		}
		r.applyDecorations(out, "Begin", n.Decs.Begin, false)
		if n.Dir == model.SEND {
			out.Arrow = r.cursor
			r.cursor += token.Pos(len(token.ARROW.String()))
		}
		r.applyDecorations(out, "Arrow", n.Decs.Arrow, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "ChanType", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Dir = ast.ChanDir(n.Dir)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.CommClause:
		out := &ast.CommClause{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Case = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.Comm == nil {
				return token.DEFAULT
			}
			return token.CASE
		}().String()))
		r.applyDecorations(out, "Case", n.Decs.Case, false)
		if n.Comm != nil {
			out.Comm = r.restoreNode(n.Comm, "CommClause", "Comm", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "Comm", n.Decs.Comm, false)
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecorations(out, "Colon", n.Decs.Colon, false)
		for _, v := range n.Body {
			out.Body = append(out.Body, r.restoreNode(v, "CommClause", "Body", "Stmt", allowDuplicate).(ast.Stmt))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.CompositeLit:
		out := &ast.CompositeLit{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "CompositeLit", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Type", n.Decs.Type, false)
		out.Lbrace = r.cursor
		r.cursor += token.Pos(len(token.LBRACE.String()))
		r.applyDecorations(out, "Lbrace", n.Decs.Lbrace, false)
		for _, v := range n.Elts {
			out.Elts = append(out.Elts, r.restoreNode(v, "CompositeLit", "Elts", "Expr", allowDuplicate).(ast.Expr))
		}
		out.Rbrace = r.cursor
		r.cursor += token.Pos(len(token.RBRACE.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Incomplete = n.Incomplete
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.DeclStmt:
		out := &ast.DeclStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Decl != nil {
			out.Decl = r.restoreNode(n.Decl, "DeclStmt", "Decl", "Decl", allowDuplicate).(ast.Decl)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.DeferStmt:
		out := &ast.DeferStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Defer = r.cursor
		r.cursor += token.Pos(len(token.DEFER.String()))
		r.applyDecorations(out, "Defer", n.Decs.Defer, false)
		if n.Call != nil {
			out.Call = r.restoreNode(n.Call, "DeferStmt", "Call", "CallExpr", allowDuplicate).(*ast.CallExpr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.Ellipsis:
		out := &ast.Ellipsis{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Ellipsis = r.cursor
		r.cursor += token.Pos(len(token.ELLIPSIS.String()))
		r.applyDecorations(out, "Ellipsis", n.Decs.Ellipsis, false)
		if n.Elt != nil {
			out.Elt = r.restoreNode(n.Elt, "Ellipsis", "Elt", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.EmptyStmt:
		out := &ast.EmptyStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if !n.Implicit {
			out.Semicolon = r.cursor
			r.cursor += token.Pos(len(token.ARROW.String()))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Implicit = n.Implicit
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.ExprStmt:
		out := &ast.ExprStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "ExprStmt", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.Field:
		out := &ast.Field{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		for _, v := range n.Names {
			out.Names = append(out.Names, r.restoreNode(v, "Field", "Names", "Ident", allowDuplicate).(*ast.Ident))
		}
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "Field", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Type", n.Decs.Type, false)
		if n.Tag != nil {
			out.Tag = r.restoreNode(n.Tag, "Field", "Tag", "BasicLit", allowDuplicate).(*ast.BasicLit)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.FieldList:
		out := &ast.FieldList{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Opening {
			out.Opening = r.cursor
			r.cursor += token.Pos(len(token.LPAREN.String()))
		}
		r.applyDecorations(out, "Opening", n.Decs.Opening, false)
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v, "FieldList", "List", "Field", allowDuplicate).(*ast.Field))
		}
		if n.Closing {
			out.Closing = r.cursor
			r.cursor += token.Pos(len(token.RPAREN.String()))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.File:
		out := &ast.File{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Package = r.cursor
		r.cursor += token.Pos(len(token.PACKAGE.String()))
		r.applyDecorations(out, "Package", n.Decs.Package, false)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "File", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecorations(out, "Name", n.Decs.Name, false)
		for _, v := range n.Decls {
			out.Decls = append(out.Decls, r.restoreNode(v, "File", "Decls", "Decl", allowDuplicate).(ast.Decl))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Scope = r.restoreScope(n.Scope)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.ForStmt:
		out := &ast.ForStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.For = r.cursor
		r.cursor += token.Pos(len(token.FOR.String()))
		r.applyDecorations(out, "For", n.Decs.For, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "ForStmt", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		if n.Init != nil {
			r.cursor += token.Pos(len(token.SEMICOLON.String()))
		}
		r.applyDecorations(out, "Init", n.Decs.Init, false)
		if n.Cond != nil {
			out.Cond = r.restoreNode(n.Cond, "ForStmt", "Cond", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Post != nil {
			r.cursor += token.Pos(len(token.SEMICOLON.String()))
		}
		r.applyDecorations(out, "Cond", n.Decs.Cond, false)
		if n.Post != nil {
			out.Post = r.restoreNode(n.Post, "ForStmt", "Post", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "Post", n.Decs.Post, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "ForStmt", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.FuncDecl:
		out := &ast.FuncDecl{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		out.Type = &ast.FuncType{}
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		r.applyDecorations(out, "Start", n.Type.Decs.Start, false)
		if true {
			out.Type.Func = r.cursor
			r.cursor += token.Pos(len(token.FUNC.String()))
		}
		r.applyDecorations(out, "Func", n.Decs.Func, false)
		r.applyDecorations(out, "Func", n.Type.Decs.Func, false)
		if n.Recv != nil {
			out.Recv = r.restoreNode(n.Recv, "FuncDecl", "Recv", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "Recv", n.Decs.Recv, false)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "FuncDecl", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecorations(out, "Name", n.Decs.Name, false)
		if n.Type.TypeParams != nil {
			out.Type.TypeParams = r.restoreNode(n.Type.TypeParams, "FuncDecl", "TypeParams", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "TypeParams", n.Decs.TypeParams, false)
		r.applyDecorations(out, "TypeParams", n.Type.Decs.TypeParams, false)
		if n.Type.Params != nil {
			out.Type.Params = r.restoreNode(n.Type.Params, "FuncDecl", "Params", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "Params", n.Decs.Params, false)
		r.applyDecorations(out, "Params", n.Type.Decs.Params, false)
		if n.Type.Results != nil {
			out.Type.Results = r.restoreNode(n.Type.Results, "FuncDecl", "Results", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "Results", n.Decs.Results, false)
		r.applyDecorations(out, "End", n.Type.Decs.End, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "FuncDecl", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.FuncLit:
		out := &ast.FuncLit{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "FuncLit", "Type", "FuncType", allowDuplicate).(*ast.FuncType)
		}
		r.applyDecorations(out, "Type", n.Decs.Type, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "FuncLit", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.FuncType:
		out := &ast.FuncType{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Func {
			out.Func = r.cursor
			r.cursor += token.Pos(len(token.FUNC.String()))
		}
		r.applyDecorations(out, "Func", n.Decs.Func, false)
		if n.TypeParams != nil {
			out.TypeParams = r.restoreNode(n.TypeParams, "FuncType", "TypeParams", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "TypeParams", n.Decs.TypeParams, false)
		if n.Params != nil {
			out.Params = r.restoreNode(n.Params, "FuncType", "Params", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "Params", n.Decs.Params, false)
		if n.Results != nil {
			out.Results = r.restoreNode(n.Results, "FuncType", "Results", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.GenDecl:
		out := &ast.GenDecl{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applyDecorations(out, "Tok", n.Decs.Tok, false)
		if n.Lparen {
			out.Lparen = r.cursor
			r.cursor += token.Pos(len(token.LPAREN.String()))
		}
		r.applyDecorations(out, "Lparen", n.Decs.Lparen, false)
		for _, v := range n.Specs {
			out.Specs = append(out.Specs, r.restoreNode(v, "GenDecl", "Specs", "Spec", allowDuplicate).(ast.Spec))
		}
		if n.Rparen {
			out.Rparen = r.cursor
			r.cursor += token.Pos(len(token.RPAREN.String()))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.GoStmt:
		out := &ast.GoStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Go = r.cursor
		r.cursor += token.Pos(len(token.GO.String()))
		r.applyDecorations(out, "Go", n.Decs.Go, false)
		if n.Call != nil {
			out.Call = r.restoreNode(n.Call, "GoStmt", "Call", "CallExpr", allowDuplicate).(*ast.CallExpr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.Ident:
		sel := r.restoreIdent(n, parentName, parentField, parentFieldType, allowDuplicate)
		if sel != nil {
			return sel
		}

		out := &ast.Ident{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.NamePos = r.cursor
		out.Name = n.Name
		r.cursor += token.Pos(len(n.Name))
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Obj = r.restoreObject(n.Obj)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.IfStmt:
		out := &ast.IfStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.If = r.cursor
		r.cursor += token.Pos(len(token.IF.String()))
		r.applyDecorations(out, "If", n.Decs.If, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "IfStmt", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "Init", n.Decs.Init, false)
		if n.Cond != nil {
			out.Cond = r.restoreNode(n.Cond, "IfStmt", "Cond", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Cond", n.Decs.Cond, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "IfStmt", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		if n.Else != nil {
			r.cursor += token.Pos(len(token.ELSE.String()))
		}
		r.applyDecorations(out, "Else", n.Decs.Else, false)
		if n.Else != nil {
			out.Else = r.restoreNode(n.Else, "IfStmt", "Else", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.ImportSpec:
		out := &ast.ImportSpec{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "ImportSpec", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecorations(out, "Name", n.Decs.Name, false)
		if n.Path != nil {
			out.Path = r.restoreNode(n.Path, "ImportSpec", "Path", "BasicLit", allowDuplicate).(*ast.BasicLit)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.IncDecStmt:
		out := &ast.IncDecStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "IncDecStmt", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.IndexExpr:
		out := &ast.IndexExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "IndexExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecorations(out, "Lbrack", n.Decs.Lbrack, false)
		if n.Index != nil {
			out.Index = r.restoreNode(n.Index, "IndexExpr", "Index", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Index", n.Decs.Index, false)
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.IndexListExpr:
		out := &ast.IndexListExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "IndexListExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecorations(out, "Lbrack", n.Decs.Lbrack, false)
		for _, v := range n.Indices {
			out.Indices = append(out.Indices, r.restoreNode(v, "IndexListExpr", "Indices", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applyDecorations(out, "Indices", n.Decs.Indices, false)
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.InterfaceType:
		out := &ast.InterfaceType{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Interface = r.cursor
		r.cursor += token.Pos(len(token.INTERFACE.String()))
		r.applyDecorations(out, "Interface", n.Decs.Interface, false)
		if n.Methods != nil {
			out.Methods = r.restoreNode(n.Methods, "InterfaceType", "Methods", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Incomplete = n.Incomplete
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.KeyValueExpr:
		out := &ast.KeyValueExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key, "KeyValueExpr", "Key", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Key", n.Decs.Key, false)
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecorations(out, "Colon", n.Decs.Colon, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "KeyValueExpr", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.LabeledStmt:
		out := &ast.LabeledStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Label != nil {
			out.Label = r.restoreNode(n.Label, "LabeledStmt", "Label", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecorations(out, "Label", n.Decs.Label, false)
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecorations(out, "Colon", n.Decs.Colon, false)
		if n.Stmt != nil {
			out.Stmt = r.restoreNode(n.Stmt, "LabeledStmt", "Stmt", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.MapType:
		out := &ast.MapType{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Map = r.cursor
		r.cursor += token.Pos(len(token.MAP.String()))
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecorations(out, "Map", n.Decs.Map, false)
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key, "MapType", "Key", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecorations(out, "Key", n.Decs.Key, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "MapType", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.Package:
		out := &ast.Package{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		out.Name = n.Name
		out.Scope = r.restoreScope(n.Scope)
		out.Imports = map[string]*ast.Object{}
		for k, v := range n.Imports {
			out.Imports[k] = r.restoreObject(v)
		}
		out.Files = map[string]*ast.File{}
		for k, v := range n.Files {
			out.Files[k] = r.restoreNode(v, "Package", "Files", "File", allowDuplicate).(*ast.File)
		}

		return out
	case *model.ParenExpr:
		out := &ast.ParenExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))
		r.applyDecorations(out, "Lparen", n.Decs.Lparen, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "ParenExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.RangeStmt:
		out := &ast.RangeStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.For = r.cursor
		r.cursor += token.Pos(len(token.FOR.String()))
		r.applyDecorations(out, "For", n.Decs.For, false)
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key, "RangeStmt", "Key", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Value != nil {
			r.cursor += token.Pos(len(token.COMMA.String()))
		}
		r.applyDecorations(out, "Key", n.Decs.Key, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "RangeStmt", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Value", n.Decs.Value, false)
		if n.Tok != token.ILLEGAL {
			out.Tok = n.Tok
			out.TokPos = r.cursor
			r.cursor += token.Pos(len(n.Tok.String()))
		}
		r.cursor += token.Pos(len(token.RANGE.String()))
		r.applyDecorations(out, "Range", n.Decs.Range, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "RangeStmt", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "X", n.Decs.X, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "RangeStmt", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.ReturnStmt:
		out := &ast.ReturnStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Return = r.cursor
		r.cursor += token.Pos(len(token.RETURN.String()))
		r.applyDecorations(out, "Return", n.Decs.Return, false)
		for _, v := range n.Results {
			out.Results = append(out.Results, r.restoreNode(v, "ReturnStmt", "Results", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.SelectStmt:
		out := &ast.SelectStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Select = r.cursor
		r.cursor += token.Pos(len(token.SELECT.String()))
		r.applyDecorations(out, "Select", n.Decs.Select, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "SelectStmt", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.SelectorExpr:
		out := &ast.SelectorExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "SelectorExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.PERIOD.String()))
		r.applyDecorations(out, "X", n.Decs.X, false)
		if n.Sel != nil {
			out.Sel = r.restoreNode(n.Sel, "SelectorExpr", "Sel", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.SendStmt:
		out := &ast.SendStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Chan != nil {
			out.Chan = r.restoreNode(n.Chan, "SendStmt", "Chan", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Chan", n.Decs.Chan, false)
		out.Arrow = r.cursor
		r.cursor += token.Pos(len(token.ARROW.String()))
		r.applyDecorations(out, "Arrow", n.Decs.Arrow, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "SendStmt", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.SliceExpr:
		out := &ast.SliceExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "SliceExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecorations(out, "Lbrack", n.Decs.Lbrack, false)
		if n.Low != nil {
			out.Low = r.restoreNode(n.Low, "SliceExpr", "Low", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecorations(out, "Low", n.Decs.Low, false)
		if n.High != nil {
			out.High = r.restoreNode(n.High, "SliceExpr", "High", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Slice3 {
			r.cursor += token.Pos(len(token.COLON.String()))
		}
		r.applyDecorations(out, "High", n.Decs.High, false)
		if n.Max != nil {
			out.Max = r.restoreNode(n.Max, "SliceExpr", "Max", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Max", n.Decs.Max, false)
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Slice3 = n.Slice3
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.StarExpr:
		out := &ast.StarExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Star = r.cursor
		r.cursor += token.Pos(len(token.MUL.String()))
		r.applyDecorations(out, "Star", n.Decs.Star, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "StarExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.StructType:
		out := &ast.StructType{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Struct = r.cursor
		r.cursor += token.Pos(len(token.STRUCT.String()))
		r.applyDecorations(out, "Struct", n.Decs.Struct, false)
		if n.Fields != nil {
			out.Fields = r.restoreNode(n.Fields, "StructType", "Fields", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		out.Incomplete = n.Incomplete
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.SwitchStmt:
		out := &ast.SwitchStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Switch = r.cursor
		r.cursor += token.Pos(len(token.SWITCH.String()))
		r.applyDecorations(out, "Switch", n.Decs.Switch, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "SwitchStmt", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "Init", n.Decs.Init, false)
		if n.Tag != nil {
			out.Tag = r.restoreNode(n.Tag, "SwitchStmt", "Tag", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "Tag", n.Decs.Tag, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "SwitchStmt", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.TypeAssertExpr:
		out := &ast.TypeAssertExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "TypeAssertExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.PERIOD.String()))
		r.applyDecorations(out, "X", n.Decs.X, false)
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))
		r.applyDecorations(out, "Lparen", n.Decs.Lparen, false)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "TypeAssertExpr", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Type == nil {
			r.cursor += token.Pos(len(token.TYPE.String()))
		}
		r.applyDecorations(out, "Type", n.Decs.Type, false)
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.TypeSpec:
		out := &ast.TypeSpec{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "TypeSpec", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		if n.Assign {
			out.Assign = r.cursor
			r.cursor += token.Pos(len(token.ASSIGN.String()))
		}
		r.applyDecorations(out, "Name", n.Decs.Name, false)
		if n.TypeParams != nil {
			out.TypeParams = r.restoreNode(n.TypeParams, "TypeSpec", "TypeParams", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecorations(out, "TypeParams", n.Decs.TypeParams, false)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "TypeSpec", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.TypeSwitchStmt:
		out := &ast.TypeSwitchStmt{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Switch = r.cursor
		r.cursor += token.Pos(len(token.SWITCH.String()))
		r.applyDecorations(out, "Switch", n.Decs.Switch, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "TypeSwitchStmt", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "Init", n.Decs.Init, false)
		if n.Assign != nil {
			out.Assign = r.restoreNode(n.Assign, "TypeSwitchStmt", "Assign", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecorations(out, "Assign", n.Decs.Assign, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "TypeSwitchStmt", "Body", "BlockStmt", allowDuplicate).(*ast.BlockStmt)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.UnaryExpr:
		out := &ast.UnaryExpr{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		out.Op = n.Op
		out.OpPos = r.cursor
		r.cursor += token.Pos(len(n.Op.String()))
		r.applyDecorations(out, "Op", n.Decs.Op, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "UnaryExpr", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	case *model.ValueSpec:
		out := &ast.ValueSpec{}
		r.Ast.Nodes[n] = out
		r.Dst.Nodes[out] = n
		r.applySpace(n, "Before", n.Decs.Before)
		r.applyDecorations(out, "Start", n.Decs.Start, false)
		for _, v := range n.Names {
			out.Names = append(out.Names, r.restoreNode(v, "ValueSpec", "Names", "Ident", allowDuplicate).(*ast.Ident))
		}
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "ValueSpec", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Values != nil {
			r.cursor += token.Pos(len(token.ASSIGN.String()))
		}
		r.applyDecorations(out, "Assign", n.Decs.Assign, false)
		for _, v := range n.Values {
			out.Values = append(out.Values, r.restoreNode(v, "ValueSpec", "Values", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applyDecorations(out, "End", n.Decs.End, true)
		r.applySpace(n, "After", n.Decs.After)

		return out
	default:
		panic(fmt.Sprintf("%T", n))
	}
}
