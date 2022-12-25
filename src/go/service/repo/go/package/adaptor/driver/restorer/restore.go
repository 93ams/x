package restorer

import (
	"fmt"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
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
	case *model.Array:
		out := &ast.ArrayType{}
		r.applyPrefix(n, out)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecs(out, "Lbrack", n.Decs.Lbrack, false)
		if n.Len != nil {
			out.Len = r.restoreNode(n.Len, "Array", "Len", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecs(out, "Len", n.Decs.Len, false)
		if n.Elt != nil {
			out.Elt = r.restoreNode(n.Elt, "Array", "Elt", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Assign:
		out := &ast.AssignStmt{}
		r.applyPrefix(n, out)
		for _, v := range n.Lhs {
			out.Lhs = append(out.Lhs, r.restoreNode(v, "Assign", "Lhs", "Expr", allowDuplicate).(ast.Expr))
		}
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applyDecs(out, "Tok", n.Decs.Tok, false)
		for _, v := range n.Rhs {
			out.Rhs = append(out.Rhs, r.restoreNode(v, "Assign", "Rhs", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applySuffix(n, out)
		return out
	case *model.BadDecl:
		out := &ast.BadDecl{}
		r.applyPrefix(n, out)
		out.From = r.cursor
		r.cursor += token.Pos(n.Length)
		out.To = r.cursor
		r.applySuffix(n, out)
		return out
	case *model.BadExpr:
		out := &ast.BadExpr{}
		r.applyPrefix(n, out)
		out.From = r.cursor
		r.cursor += token.Pos(n.Length)
		out.To = r.cursor
		r.applySuffix(n, out)
		return out
	case *model.BadStmt:
		out := &ast.BadStmt{}
		r.applyPrefix(n, out)
		out.From = r.cursor
		r.cursor += token.Pos(n.Length)
		out.To = r.cursor
		r.applySuffix(n, out)
		return out
	case *model.Lit:
		out := &ast.BasicLit{}
		r.applyPrefix(n, out)
		r.applyLiteral(n.Value)
		out.ValuePos = r.cursor
		out.Value = n.Value
		r.cursor += token.Pos(len(n.Value))
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Kind = n.Kind
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.Binary:
		out := &ast.BinaryExpr{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Binary", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "X", n.Decs.X, false)
		out.Op = n.Op
		out.OpPos = r.cursor
		r.cursor += token.Pos(len(n.Op.String()))
		r.applyDecs(out, "Op", n.Decs.Op, false)
		if n.Y != nil {
			out.Y = r.restoreNode(n.Y, "Binary", "Y", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Block:
		out := &ast.BlockStmt{}
		r.applyPrefix(n, out)
		out.Lbrace = r.cursor
		r.cursor += token.Pos(len(token.LBRACE.String()))
		r.applyDecs(out, "Lbrace", n.Decs.Lbrace, false)
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v, "Block", "List", "Stmt", allowDuplicate).(ast.Stmt))
		}
		if n.RbraceHasNoPos {
			out.Rbrace = token.NoPos
		} else {
			out.Rbrace = r.cursor
		}
		r.cursor += token.Pos(len(token.RBRACE.String()))
		r.applySuffix(n, out)
		return out
	case *model.Branch:
		out := &ast.BranchStmt{}
		r.applyPrefix(n, out)
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applyDecs(out, "Tok", n.Decs.Tok, false)
		if n.Label != nil {
			out.Label = r.restoreNode(n.Label, "Branch", "Label", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applySuffix(n, out)
		return out
	case *model.Call:
		out := &ast.CallExpr{}
		r.applyPrefix(n, out)
		if n.Fun != nil {
			out.Fun = r.restoreNode(n.Fun, "Call", "Fun", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Fun", n.Decs.Fun, false)
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))
		r.applyDecs(out, "Lparen", n.Decs.Lparen, false)
		for _, v := range n.Args {
			out.Args = append(out.Args, r.restoreNode(v, "Call", "Args", "Expr", allowDuplicate).(ast.Expr))
		}
		if n.Ellipsis {
			out.Ellipsis = r.cursor
			r.cursor += token.Pos(len(token.ELLIPSIS.String()))
		}
		r.applyDecs(out, "Ellipsis", n.Decs.Ellipsis, false)
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))
		r.applySuffix(n, out)
		return out
	case *model.CaseClause:
		out := &ast.CaseClause{}
		r.applyPrefix(n, out)
		out.Case = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.List == nil {
				return token.DEFAULT
			}
			return token.CASE
		}().String()))
		r.applyDecs(out, "Case", n.Decs.Case, false)
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v, "CaseClause", "List", "Expr", allowDuplicate).(ast.Expr))
		}
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecs(out, "Colon", n.Decs.Colon, false)
		for _, v := range n.Body {
			out.Body = append(out.Body, r.restoreNode(v, "CaseClause", "Body", "Stmt", allowDuplicate).(ast.Stmt))
		}
		r.applySuffix(n, out)
		return out
	case *model.Chan:
		out := &ast.ChanType{}
		r.applyPrefix(n, out)
		out.Begin = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.Dir == model.RecvChan {
				return token.ARROW
			}
			return token.CHAN
		}().String()))
		if n.Dir == model.RecvChan {
			r.cursor += token.Pos(len(token.CHAN.String()))
		}
		r.applyDecs(out, "Begin", n.Decs.Begin, false)
		if n.Dir == model.SendChan {
			out.Arrow = r.cursor
			r.cursor += token.Pos(len(token.ARROW.String()))
		}
		r.applyDecs(out, "Arrow", n.Decs.Arrow, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "Chan", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Dir = ast.ChanDir(n.Dir)
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.CommClause:
		out := &ast.CommClause{}
		r.applyPrefix(n, out)
		out.Case = r.cursor
		r.cursor += token.Pos(len(func() token.Token {
			if n.Comm == nil {
				return token.DEFAULT
			}
			return token.CASE
		}().String()))
		r.applyDecs(out, "Case", n.Decs.Case, false)
		if n.Comm != nil {
			out.Comm = r.restoreNode(n.Comm, "CommClause", "Comm", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecs(out, "Comm", n.Decs.Comm, false)
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecs(out, "Colon", n.Decs.Colon, false)
		for _, v := range n.Body {
			out.Body = append(out.Body, r.restoreNode(v, "CommClause", "Body", "Stmt", allowDuplicate).(ast.Stmt))
		}
		r.applySuffix(n, out)
		return out
	case *model.CompositeLit:
		out := &ast.CompositeLit{}
		r.applyPrefix(n, out)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "CompositeLit", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Type", n.Decs.Type, false)
		out.Lbrace = r.cursor
		r.cursor += token.Pos(len(token.LBRACE.String()))
		r.applyDecs(out, "Lbrace", n.Decs.Lbrace, false)
		for _, v := range n.Elts {
			out.Elts = append(out.Elts, r.restoreNode(v, "CompositeLit", "Elts", "Expr", allowDuplicate).(ast.Expr))
		}
		out.Rbrace = r.cursor
		r.cursor += token.Pos(len(token.RBRACE.String()))
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Incomplete = n.Incomplete
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.DeclStmt:
		out := &ast.DeclStmt{}
		r.applyPrefix(n, out)
		if n.Decl != nil {
			out.Decl = r.restoreNode(n.Decl, "DeclStmt", "Decl", "Decl", allowDuplicate).(ast.Decl)
		}
		r.applySuffix(n, out)
		return out
	case *model.Defer:
		out := &ast.DeferStmt{}
		r.applyPrefix(n, out)
		out.Defer = r.cursor
		r.cursor += token.Pos(len(token.DEFER.String()))
		r.applyDecs(out, "Defer", n.Decs.Defer, false)
		if n.Call != nil {
			out.Call = r.restoreNode(n.Call, "Defer", "Call", "Call", allowDuplicate).(*ast.CallExpr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Ellipsis:
		out := &ast.Ellipsis{}
		r.applyPrefix(n, out)
		out.Ellipsis = r.cursor
		r.cursor += token.Pos(len(token.ELLIPSIS.String()))
		r.applyDecs(out, "Ellipsis", n.Decs.Ellipsis, false)
		if n.Elt != nil {
			out.Elt = r.restoreNode(n.Elt, "Ellipsis", "Elt", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Empty:
		out := &ast.EmptyStmt{}
		r.applyPrefix(n, out)
		if !n.Implicit {
			out.Semicolon = r.cursor
			r.cursor += token.Pos(len(token.ARROW.String()))
		}
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Implicit = n.Implicit
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.ExprStmt:
		out := &ast.ExprStmt{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "ExprStmt", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Field:
		out := &ast.Field{}
		r.applyPrefix(n, out)
		for _, v := range n.Names {
			out.Names = append(out.Names, r.restoreNode(v, "Field", "Names", "Ident", allowDuplicate).(*ast.Ident))
		}
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "Field", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Type", n.Decs.Type, false)
		if n.Tag != nil {
			out.Tag = r.restoreNode(n.Tag, "Field", "Tag", "Lit", allowDuplicate).(*ast.BasicLit)
		}
		r.applySuffix(n, out)
		return out
	case *model.FieldList:
		out := &ast.FieldList{}
		r.applyPrefix(n, out)
		if n.Opening {
			out.Opening = r.cursor
			r.cursor += token.Pos(len(token.LPAREN.String()))
		}
		r.applyDecs(out, "Opening", n.Decs.Opening, false)
		for _, v := range n.List {
			out.List = append(out.List, r.restoreNode(v, "FieldList", "List", "Field", allowDuplicate).(*ast.Field))
		}
		if n.Closing {
			out.Closing = r.cursor
			r.cursor += token.Pos(len(token.RPAREN.String()))
		}
		r.applySuffix(n, out)
		return out
	case *model.File:
		out := &ast.File{}
		r.applyPrefix(n, out)
		out.Package = r.cursor
		r.cursor += token.Pos(len(token.PACKAGE.String()))
		r.applyDecs(out, "Package", n.Decs.Package, false)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "File", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecs(out, "Name", n.Decs.Name, false)
		for _, v := range n.Decls {
			out.Decls = append(out.Decls, r.restoreNode(v, "File", "Decls", "Decl", allowDuplicate).(ast.Decl))
		}
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Scope = r.restoreScope(n.Scope)
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.For:
		out := &ast.ForStmt{}
		r.applyPrefix(n, out)
		out.For = r.cursor
		r.cursor += token.Pos(len(token.FOR.String()))
		r.applyDecs(out, "For", n.Decs.For, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "For", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		if n.Init != nil {
			r.cursor += token.Pos(len(token.SEMICOLON.String()))
		}
		r.applyDecs(out, "Init", n.Decs.Init, false)
		if n.Cond != nil {
			out.Cond = r.restoreNode(n.Cond, "For", "Cond", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Post != nil {
			r.cursor += token.Pos(len(token.SEMICOLON.String()))
		}
		r.applyDecs(out, "Cond", n.Decs.Cond, false)
		if n.Post != nil {
			out.Post = r.restoreNode(n.Post, "For", "Post", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecs(out, "Post", n.Decs.Post, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "For", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.Func:
		out := &ast.FuncDecl{}
		r.Ast.Nodes[n], r.Dst.Nodes[out] = out, n
		r.applySpace(n, "Before", n.Decs.Before)
		out.Type = &ast.FuncType{}
		r.applyDecs(out, "Start", n.Decs.Start, false)
		r.applyDecs(out, "Start", n.Type.Decs.Start, false)
		if true {
			out.Type.Func = r.cursor
			r.cursor += token.Pos(len(token.FUNC.String()))
		}
		r.applyDecs(out, "Func", n.Decs.Func, false)
		r.applyDecs(out, "Func", n.Type.Decs.Func, false)
		if n.Recv != nil {
			out.Recv = r.restoreNode(n.Recv, "Func", "RecvChan", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "RecvChan", n.Decs.Recv, false)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "Func", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecs(out, "Name", n.Decs.Name, false)
		if n.Type.TypeParams != nil {
			out.Type.TypeParams = r.restoreNode(n.Type.TypeParams, "Func", "Params", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "Params", n.Decs.TypeParams, false)
		r.applyDecs(out, "Params", n.Type.Decs.TypeParams, false)
		if n.Type.Params != nil {
			out.Type.Params = r.restoreNode(n.Type.Params, "Func", "Params", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "Params", n.Decs.Params, false)
		r.applyDecs(out, "Params", n.Type.Decs.Params, false)
		if n.Type.Results != nil {
			out.Type.Results = r.restoreNode(n.Type.Results, "Func", "Results", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "Results", n.Decs.Results, false)
		r.applyDecs(out, "End", n.Type.Decs.End, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "Func", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.FuncLit:
		out := &ast.FuncLit{}
		r.applyPrefix(n, out)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "FuncLit", "Type", "FuncType", allowDuplicate).(*ast.FuncType)
		}
		r.applyDecs(out, "Type", n.Decs.Type, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "FuncLit", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.FuncType:
		out := &ast.FuncType{}
		r.applyPrefix(n, out)
		if n.Func {
			out.Func = r.cursor
			r.cursor += token.Pos(len(token.FUNC.String()))
		}
		r.applyDecs(out, "Func", n.Decs.Func, false)
		if n.TypeParams != nil {
			out.TypeParams = r.restoreNode(n.TypeParams, "FuncType", "Params", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "Params", n.Decs.TypeParams, false)
		if n.Params != nil {
			out.Params = r.restoreNode(n.Params, "FuncType", "Params", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "Params", n.Decs.Params, false)
		if n.Results != nil {
			out.Results = r.restoreNode(n.Results, "FuncType", "Results", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applySuffix(n, out)
		return out
	case *model.Gen:
		out := &ast.GenDecl{}
		r.applyPrefix(n, out)
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applyDecs(out, "Tok", n.Decs.Tok, false)
		if n.Lparen {
			out.Lparen = r.cursor
			r.cursor += token.Pos(len(token.LPAREN.String()))
		}
		r.applyDecs(out, "Lparen", n.Decs.Lparen, false)
		for _, v := range n.Specs {
			out.Specs = append(out.Specs, r.restoreNode(v, "Gen", "Specs", "Spec", allowDuplicate).(ast.Spec))
		}
		if n.Rparen {
			out.Rparen = r.cursor
			r.cursor += token.Pos(len(token.RPAREN.String()))
		}
		r.applySuffix(n, out)
		return out
	case *model.Go:
		out := &ast.GoStmt{}
		r.applyPrefix(n, out)
		out.Go = r.cursor
		r.cursor += token.Pos(len(token.GO.String()))
		r.applyDecs(out, "Go", n.Decs.Go, false)
		if n.Call != nil {
			out.Call = r.restoreNode(n.Call, "Go", "Call", "Call", allowDuplicate).(*ast.CallExpr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Ident:
		sel := r.restoreIdent(n, parentName, parentField, parentFieldType, allowDuplicate)
		if sel != nil {
			return sel
		}
		out := &ast.Ident{}
		r.applyPrefix(n, out)
		r.applyDecs(out, "X", n.Decs.X, false)
		out.NamePos = r.cursor
		out.Name = n.Name
		r.cursor += token.Pos(len(n.Name))
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Obj = r.restoreObject(n.Obj)
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.If:
		out := &ast.IfStmt{}
		r.applyPrefix(n, out)
		out.If = r.cursor
		r.cursor += token.Pos(len(token.IF.String()))
		r.applyDecs(out, "If", n.Decs.If, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "If", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecs(out, "Init", n.Decs.Init, false)
		if n.Cond != nil {
			out.Cond = r.restoreNode(n.Cond, "If", "Cond", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Cond", n.Decs.Cond, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "If", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		if n.Else != nil {
			r.cursor += token.Pos(len(token.ELSE.String()))
		}
		r.applyDecs(out, "Else", n.Decs.Else, false)
		if n.Else != nil {
			out.Else = r.restoreNode(n.Else, "If", "Else", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.Import:
		out := &ast.ImportSpec{}
		r.applyPrefix(n, out)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "Import", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecs(out, "Name", n.Decs.Name, false)
		if n.Path != nil {
			out.Path = r.restoreNode(n.Path, "Import", "Path", "Lit", allowDuplicate).(*ast.BasicLit)
		}
		r.applySuffix(n, out)
		return out
	case *model.IncDec:
		out := &ast.IncDecStmt{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "IncDec", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "X", n.Decs.X, false)
		out.Tok = n.Tok
		out.TokPos = r.cursor
		r.cursor += token.Pos(len(n.Tok.String()))
		r.applySuffix(n, out)
		return out
	case *model.Index:
		out := &ast.IndexExpr{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Index", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "X", n.Decs.X, false)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecs(out, "Lbrack", n.Decs.Lbrack, false)
		if n.Index != nil {
			out.Index = r.restoreNode(n.Index, "Index", "Index", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Index", n.Decs.Index, false)
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applySuffix(n, out)
		return out
	case *model.IndexList:
		out := &ast.IndexListExpr{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "IndexList", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "X", n.Decs.X, false)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecs(out, "Lbrack", n.Decs.Lbrack, false)
		for _, v := range n.Indices {
			out.Indices = append(out.Indices, r.restoreNode(v, "IndexList", "Indices", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applyDecs(out, "Indices", n.Decs.Indices, false)
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applySuffix(n, out)
		return out
	case *model.Interface:
		out := &ast.InterfaceType{}
		r.applyPrefix(n, out)
		out.Interface = r.cursor
		r.cursor += token.Pos(len(token.INTERFACE.String()))
		r.applyDecs(out, "Interface", n.Decs.Interface, false)
		if n.Methods != nil {
			out.Methods = r.restoreNode(n.Methods, "Interface", "Methods", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Incomplete = n.Incomplete
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.KeyValue:
		out := &ast.KeyValueExpr{}
		r.applyPrefix(n, out)
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key, "KeyValue", "Key", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Key", n.Decs.Key, false)
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecs(out, "Colon", n.Decs.Colon, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "KeyValue", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Labeled:
		out := &ast.LabeledStmt{}
		r.applyPrefix(n, out)
		if n.Label != nil {
			out.Label = r.restoreNode(n.Label, "Labeled", "Label", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applyDecs(out, "Label", n.Decs.Label, false)
		out.Colon = r.cursor
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecs(out, "Colon", n.Decs.Colon, false)
		if n.Stmt != nil {
			out.Stmt = r.restoreNode(n.Stmt, "Labeled", "Stmt", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.MapType:
		out := &ast.MapType{}
		r.applyPrefix(n, out)
		out.Map = r.cursor
		r.cursor += token.Pos(len(token.MAP.String()))
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecs(out, "Map", n.Decs.Map, false)
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key, "MapType", "Key", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecs(out, "Key", n.Decs.Key, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "MapType", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Package:
		out := &ast.Package{}
		r.Ast.Nodes[n], r.Dst.Nodes[out] = out, n
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
	case *model.Paren:
		out := &ast.ParenExpr{}
		r.applyPrefix(n, out)
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))
		r.applyDecs(out, "Lparen", n.Decs.Lparen, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Paren", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "X", n.Decs.X, false)
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))
		r.applySuffix(n, out)
		return out
	case *model.Range:
		out := &ast.RangeStmt{}
		r.applyPrefix(n, out)
		out.For = r.cursor
		r.cursor += token.Pos(len(token.FOR.String()))
		r.applyDecs(out, "For", n.Decs.For, false)
		if n.Key != nil {
			out.Key = r.restoreNode(n.Key, "Range", "Key", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Value != nil {
			r.cursor += token.Pos(len(token.COMMA.String()))
		}
		r.applyDecs(out, "Key", n.Decs.Key, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "Range", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Value", n.Decs.Value, false)
		if n.Tok != token.ILLEGAL {
			out.Tok = n.Tok
			out.TokPos = r.cursor
			r.cursor += token.Pos(len(n.Tok.String()))
		}
		r.cursor += token.Pos(len(token.RANGE.String()))
		r.applyDecs(out, "Range", n.Decs.Range, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Range", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "X", n.Decs.X, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "Range", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.Return:
		out := &ast.ReturnStmt{}
		r.applyPrefix(n, out)
		out.Return = r.cursor
		r.cursor += token.Pos(len(token.RETURN.String()))
		r.applyDecs(out, "Return", n.Decs.Return, false)
		for _, v := range n.Results {
			out.Results = append(out.Results, r.restoreNode(v, "Return", "Results", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applySuffix(n, out)
		return out
	case *model.Select:
		out := &ast.SelectStmt{}
		r.applyPrefix(n, out)
		out.Select = r.cursor
		r.cursor += token.Pos(len(token.SELECT.String()))
		r.applyDecs(out, "Select", n.Decs.Select, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "Select", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.Selector:
		out := &ast.SelectorExpr{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Selector", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.PERIOD.String()))
		r.applyDecs(out, "X", n.Decs.X, false)
		if n.Sel != nil {
			out.Sel = r.restoreNode(n.Sel, "Selector", "Sel", "Ident", allowDuplicate).(*ast.Ident)
		}
		r.applySuffix(n, out)
		return out
	case *model.Send:
		out := &ast.SendStmt{}
		r.applyPrefix(n, out)
		if n.Chan != nil {
			out.Chan = r.restoreNode(n.Chan, "Send", "Chan", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Chan", n.Decs.Chan, false)
		out.Arrow = r.cursor
		r.cursor += token.Pos(len(token.ARROW.String()))
		r.applyDecs(out, "Arrow", n.Decs.Arrow, false)
		if n.Value != nil {
			out.Value = r.restoreNode(n.Value, "Send", "Value", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Slice:
		out := &ast.SliceExpr{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Slice", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "X", n.Decs.X, false)
		out.Lbrack = r.cursor
		r.cursor += token.Pos(len(token.LBRACK.String()))
		r.applyDecs(out, "Lbrack", n.Decs.Lbrack, false)
		if n.Low != nil {
			out.Low = r.restoreNode(n.Low, "Slice", "Low", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.COLON.String()))
		r.applyDecs(out, "Low", n.Decs.Low, false)
		if n.High != nil {
			out.High = r.restoreNode(n.High, "Slice", "High", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Slice3 {
			r.cursor += token.Pos(len(token.COLON.String()))
		}
		r.applyDecs(out, "High", n.Decs.High, false)
		if n.Max != nil {
			out.Max = r.restoreNode(n.Max, "Slice", "Max", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Max", n.Decs.Max, false)
		out.Rbrack = r.cursor
		r.cursor += token.Pos(len(token.RBRACK.String()))
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Slice3 = n.Slice3
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.Star:
		out := &ast.StarExpr{}
		r.applyPrefix(n, out)
		out.Star = r.cursor
		r.cursor += token.Pos(len(token.MUL.String()))
		r.applyDecs(out, "Star", n.Decs.Star, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Star", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Struct:
		out := &ast.StructType{}
		r.applyPrefix(n, out)
		out.Struct = r.cursor
		r.cursor += token.Pos(len(token.STRUCT.String()))
		r.applyDecs(out, "Struct", n.Decs.Struct, false)
		if n.Fields != nil {
			out.Fields = r.restoreNode(n.Fields, "Struct", "Methods", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "End", n.Decs.End, true)
		out.Incomplete = n.Incomplete
		r.applySpace(n, "After", n.Decs.After)
		return out
	case *model.Switch:
		out := &ast.SwitchStmt{}
		r.applyPrefix(n, out)
		out.Switch = r.cursor
		r.cursor += token.Pos(len(token.SWITCH.String()))
		r.applyDecs(out, "Switch", n.Decs.Switch, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "Switch", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecs(out, "Init", n.Decs.Init, false)
		if n.Tag != nil {
			out.Tag = r.restoreNode(n.Tag, "Switch", "Tag", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applyDecs(out, "Tag", n.Decs.Tag, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "Switch", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.TypeAssert:
		out := &ast.TypeAssertExpr{}
		r.applyPrefix(n, out)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "TypeAssert", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.cursor += token.Pos(len(token.PERIOD.String()))
		r.applyDecs(out, "X", n.Decs.X, false)
		out.Lparen = r.cursor
		r.cursor += token.Pos(len(token.LPAREN.String()))
		r.applyDecs(out, "Lparen", n.Decs.Lparen, false)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "TypeAssert", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Type == nil {
			r.cursor += token.Pos(len(token.TYPE.String()))
		}
		r.applyDecs(out, "Type", n.Decs.Type, false)
		out.Rparen = r.cursor
		r.cursor += token.Pos(len(token.RPAREN.String()))
		r.applySuffix(n, out)
		return out
	case *model.Type:
		out := &ast.TypeSpec{}
		r.applyPrefix(n, out)
		if n.Name != nil {
			out.Name = r.restoreNode(n.Name, "Type", "Name", "Ident", allowDuplicate).(*ast.Ident)
		}
		if n.Assign {
			out.Assign = r.cursor
			r.cursor += token.Pos(len(token.ASSIGN.String()))
		}
		r.applyDecs(out, "Name", n.Decs.Name, false)
		if n.Params != nil {
			out.TypeParams = r.restoreNode(n.Params, "Type", "Params", "FieldList", allowDuplicate).(*ast.FieldList)
		}
		r.applyDecs(out, "Params", n.Decs.TypeParams, false)
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "Type", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.TypeSwitch:
		out := &ast.TypeSwitchStmt{}
		r.applyPrefix(n, out)
		out.Switch = r.cursor
		r.cursor += token.Pos(len(token.SWITCH.String()))
		r.applyDecs(out, "Switch", n.Decs.Switch, false)
		if n.Init != nil {
			out.Init = r.restoreNode(n.Init, "TypeSwitch", "Init", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecs(out, "Init", n.Decs.Init, false)
		if n.Assign != nil {
			out.Assign = r.restoreNode(n.Assign, "TypeSwitch", "Assign", "Stmt", allowDuplicate).(ast.Stmt)
		}
		r.applyDecs(out, "Assign", n.Decs.Assign, false)
		if n.Body != nil {
			out.Body = r.restoreNode(n.Body, "TypeSwitch", "Body", "Block", allowDuplicate).(*ast.BlockStmt)
		}
		r.applySuffix(n, out)
		return out
	case *model.Unary:
		out := &ast.UnaryExpr{}
		r.applyPrefix(n, out)
		out.Op = n.Op
		out.OpPos = r.cursor
		r.cursor += token.Pos(len(n.Op.String()))
		r.applyDecs(out, "Op", n.Decs.Op, false)
		if n.X != nil {
			out.X = r.restoreNode(n.X, "Unary", "X", "Expr", allowDuplicate).(ast.Expr)
		}
		r.applySuffix(n, out)
		return out
	case *model.Value:
		out := &ast.ValueSpec{}
		r.applyPrefix(n, out)
		r.applyPrefix(n, out)
		for _, v := range n.Names {
			out.Names = append(out.Names, r.restoreNode(v, "Value", "Names", "Ident", allowDuplicate).(*ast.Ident))
		}
		if n.Type != nil {
			out.Type = r.restoreNode(n.Type, "Value", "Type", "Expr", allowDuplicate).(ast.Expr)
		}
		if n.Values != nil {
			r.cursor += token.Pos(len(token.ASSIGN.String()))
		}
		r.applyDecs(out, "Assign", n.Decs.Assign, false)
		for _, v := range n.Values {
			out.Values = append(out.Values, r.restoreNode(v, "Value", "Values", "Expr", allowDuplicate).(ast.Expr))
		}
		r.applySuffix(n, out)
		return out
	default:
		panic(fmt.Sprintf("%T", n))
	}
}
func (r *FileRestorer) applySuffix(n model.Node, out ast.Node) {
	d := n.Decors()
	r.applyDecs(out, "End", d.End, true)
	r.applySpace(n, "After", d.After)
}
func (r *FileRestorer) applyPrefix(n model.Node, out ast.Node) {
	r.Ast.Nodes[n], r.Dst.Nodes[out] = out, n
	d := n.Decors()
	r.applySpace(n, "Before", d.Before)
	r.applyDecs(out, "Start", d.Start, false)
}
