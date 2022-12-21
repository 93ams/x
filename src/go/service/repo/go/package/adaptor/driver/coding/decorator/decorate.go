package decorator

import (
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	"go/ast"
	"go/token"
)

func (f *fileDecorator) decorateNode(parent ast.Node, parentName, parentField, parentFieldType string, n ast.Node) (model2.Node, error) {
	if dn, ok := f.Dst.Nodes[n]; ok {
		return dn, nil
	}
	switch n := n.(type) {
	case *ast.ArrayType:
		out := &model2.Array{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n
		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]
		if n.Len != nil {
			child, err := f.decorateNode(n, "Array", "Len", "Expr", n.Len)
			if err != nil {
				return nil, err
			}
			out.Len = child.(model2.Expr)
		}
		if n.Elt != nil {
			child, err := f.decorateNode(n, "Array", "Elt", "Expr", n.Elt)
			if err != nil {
				return nil, err
			}
			out.Elt = child.(model2.Expr)
		}
		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Lbrack"]; ok {
				out.Decs.Lbrack = decs
			}
			if decs, ok := nd["Len"]; ok {
				out.Decs.Len = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}
		return out, nil
	case *ast.AssignStmt:
		out := &model2.Assign{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// List: Lhs
		for _, v := range n.Lhs {
			child, err := f.decorateNode(n, "Assign", "Lhs", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.Lhs = append(out.Lhs, child.(model2.Expr))
		}

		// Token: Tok
		out.Tok = n.Tok

		// List: Rhs
		for _, v := range n.Rhs {
			child, err := f.decorateNode(n, "Assign", "Rhs", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.Rhs = append(out.Rhs, child.(model2.Expr))
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Tok"]; ok {
				out.Decs.Tok = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.BadDecl:
		out := &model2.BadDecl{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Bad
		out.Length = int(n.To - n.From)

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.BadExpr:
		out := &model2.BadExpr{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Bad
		out.Length = int(n.To - n.From)

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.BadStmt:
		out := &model2.BadStmt{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Bad
		out.Length = int(n.To - n.From)

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.BasicLit:
		out := &model2.Lit{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// String: Value
		out.Value = n.Value

		// Value: Kind
		out.Kind = n.Kind

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.BinaryExpr:
		out := &model2.Binary{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Binary", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Op
		out.Op = n.Op

		// Node: Y
		if n.Y != nil {
			child, err := f.decorateNode(n, "Binary", "Y", "Expr", n.Y)
			if err != nil {
				return nil, err
			}
			out.Y = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["Op"]; ok {
				out.Decs.Op = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.BlockStmt:
		out := &model2.Block{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Lbrace

		// List: List
		for _, v := range n.List {
			child, err := f.decorateNode(n, "Block", "List", "Stmt", v)
			if err != nil {
				return nil, err
			}
			out.List = append(out.List, child.(model2.Stmt))
		}

		// Token: Rbrace
		if n.Rbrace == token.NoPos {
			out.RbraceHasNoPos = true
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Lbrace"]; ok {
				out.Decs.Lbrace = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.BranchStmt:
		out := &model2.Branch{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Tok
		out.Tok = n.Tok

		// Node: Label
		if n.Label != nil {
			child, err := f.decorateNode(n, "Branch", "Label", "Ident", n.Label)
			if err != nil {
				return nil, err
			}
			out.Label = child.(*model2.Ident)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Tok"]; ok {
				out.Decs.Tok = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.CallExpr:
		out := &model2.Call{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Fun
		if n.Fun != nil {
			child, err := f.decorateNode(n, "Call", "Fun", "Expr", n.Fun)
			if err != nil {
				return nil, err
			}
			out.Fun = child.(model2.Expr)
		}

		// Token: Lparen

		// List: Args
		for _, v := range n.Args {
			child, err := f.decorateNode(n, "Call", "Args", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.Args = append(out.Args, child.(model2.Expr))
		}

		// Token: Ellipsis
		out.Ellipsis = n.Ellipsis.IsValid()

		// Token: Rparen

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Fun"]; ok {
				out.Decs.Fun = decs
			}
			if decs, ok := nd["Lparen"]; ok {
				out.Decs.Lparen = decs
			}
			if decs, ok := nd["Ellipsis"]; ok {
				out.Decs.Ellipsis = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.CaseClause:
		out := &model2.CaseClause{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Case

		// List: List
		for _, v := range n.List {
			child, err := f.decorateNode(n, "CaseClause", "List", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.List = append(out.List, child.(model2.Expr))
		}

		// Token: Colon

		// List: Body
		for _, v := range n.Body {
			child, err := f.decorateNode(n, "CaseClause", "Body", "Stmt", v)
			if err != nil {
				return nil, err
			}
			out.Body = append(out.Body, child.(model2.Stmt))
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Case"]; ok {
				out.Decs.Case = decs
			}
			if decs, ok := nd["Colon"]; ok {
				out.Decs.Colon = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.ChanType:
		out := &model2.Chan{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Begin

		// Token: Chan

		// Token: Arrow

		// Node: Value
		if n.Value != nil {
			child, err := f.decorateNode(n, "Chan", "Value", "Expr", n.Value)
			if err != nil {
				return nil, err
			}
			out.Value = child.(model2.Expr)
		}

		// Value: Dir
		out.Dir = model2.ChanDir(n.Dir)

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Begin"]; ok {
				out.Decs.Begin = decs
			}
			if decs, ok := nd["Arrow"]; ok {
				out.Decs.Arrow = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.CommClause:
		out := &model2.CommClause{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Case

		// Node: Comm
		if n.Comm != nil {
			child, err := f.decorateNode(n, "CommClause", "Comm", "Stmt", n.Comm)
			if err != nil {
				return nil, err
			}
			out.Comm = child.(model2.Stmt)
		}

		// Token: Colon

		// List: Body
		for _, v := range n.Body {
			child, err := f.decorateNode(n, "CommClause", "Body", "Stmt", v)
			if err != nil {
				return nil, err
			}
			out.Body = append(out.Body, child.(model2.Stmt))
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Case"]; ok {
				out.Decs.Case = decs
			}
			if decs, ok := nd["Comm"]; ok {
				out.Decs.Comm = decs
			}
			if decs, ok := nd["Colon"]; ok {
				out.Decs.Colon = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.CompositeLit:
		out := &model2.CompositeLit{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Type
		if n.Type != nil {
			child, err := f.decorateNode(n, "CompositeLit", "Type", "Expr", n.Type)
			if err != nil {
				return nil, err
			}
			out.Type = child.(model2.Expr)
		}

		// Token: Lbrace

		// List: Elts
		for _, v := range n.Elts {
			child, err := f.decorateNode(n, "CompositeLit", "Elts", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.Elts = append(out.Elts, child.(model2.Expr))
		}

		// Token: Rbrace

		// Value: Ellipsis
		out.Incomplete = n.Incomplete

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Type"]; ok {
				out.Decs.Type = decs
			}
			if decs, ok := nd["Lbrace"]; ok {
				out.Decs.Lbrace = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.DeclStmt:
		out := &model2.DeclStmt{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Decl
		if n.Decl != nil {
			child, err := f.decorateNode(n, "DeclStmt", "Decl", "Decl", n.Decl)
			if err != nil {
				return nil, err
			}
			out.Decl = child.(model2.Decl)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.DeferStmt:
		out := &model2.Defer{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Defer

		// Node: Call
		if n.Call != nil {
			child, err := f.decorateNode(n, "Defer", "Call", "Call", n.Call)
			if err != nil {
				return nil, err
			}
			out.Call = child.(*model2.Call)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Defer"]; ok {
				out.Decs.Defer = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.Ellipsis:
		out := &model2.Ellipsis{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Ellipsis

		// Node: Elt
		if n.Elt != nil {
			child, err := f.decorateNode(n, "Ellipsis", "Elt", "Expr", n.Elt)
			if err != nil {
				return nil, err
			}
			out.Elt = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Ellipsis"]; ok {
				out.Decs.Ellipsis = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.EmptyStmt:
		out := &model2.Empty{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Semicolon

		// Value: Implicit
		out.Implicit = n.Implicit

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.ExprStmt:
		out := &model2.ExprStmt{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "ExprStmt", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.Field:
		out := &model2.Field{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// List: Names
		for _, v := range n.Names {
			child, err := f.decorateNode(n, "Field", "Names", "Ident", v)
			if err != nil {
				return nil, err
			}
			out.Names = append(out.Names, child.(*model2.Ident))
		}

		// Node: Type
		if n.Type != nil {
			child, err := f.decorateNode(n, "Field", "Type", "Expr", n.Type)
			if err != nil {
				return nil, err
			}
			out.Type = child.(model2.Expr)
		}

		// Node: Tag
		if n.Tag != nil {
			child, err := f.decorateNode(n, "Field", "Tag", "Lit", n.Tag)
			if err != nil {
				return nil, err
			}
			out.Tag = child.(*model2.Lit)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Type"]; ok {
				out.Decs.Type = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.FieldList:
		out := &model2.FieldList{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Opening
		out.Opening = n.Opening.IsValid()

		// List: List
		for _, v := range n.List {
			child, err := f.decorateNode(n, "FieldList", "List", "Field", v)
			if err != nil {
				return nil, err
			}
			out.List = append(out.List, child.(*model2.Field))
		}

		// Token: Closing
		out.Closing = n.Closing.IsValid()

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Opening"]; ok {
				out.Decs.Opening = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.File:
		out := &model2.File{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Package

		// Node: Name
		if n.Name != nil {
			child, err := f.decorateNode(n, "File", "Name", "Ident", n.Name)
			if err != nil {
				return nil, err
			}
			out.Name = child.(*model2.Ident)
		}

		// List: Decls
		for _, v := range n.Decls {
			child, err := f.decorateNode(n, "File", "Decls", "Decl", v)
			if err != nil {
				return nil, err
			}
			out.Decls = append(out.Decls, child.(model2.Decl))
		}

		// Scope: Scope
		scope, err := f.decorateScope(n.Scope)
		if err != nil {
			return nil, err
		}
		out.Scope = scope

		// List: Imports
		for _, v := range n.Imports {
			child, err := f.decorateNode(n, "File", "Imports", "Import", v)
			if err != nil {
				return nil, err
			}
			out.Imports = append(out.Imports, child.(*model2.Import))
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Package"]; ok {
				out.Decs.Package = decs
			}
			if decs, ok := nd["Name"]; ok {
				out.Decs.Name = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.ForStmt:
		out := &model2.For{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: For

		// Node: Init
		if n.Init != nil {
			child, err := f.decorateNode(n, "For", "Init", "Stmt", n.Init)
			if err != nil {
				return nil, err
			}
			out.Init = child.(model2.Stmt)
		}

		// Token: InitSemicolon

		// Node: Cond
		if n.Cond != nil {
			child, err := f.decorateNode(n, "For", "Cond", "Expr", n.Cond)
			if err != nil {
				return nil, err
			}
			out.Cond = child.(model2.Expr)
		}

		// Token: CondSemicolon

		// Node: Post
		if n.Post != nil {
			child, err := f.decorateNode(n, "For", "Post", "Stmt", n.Post)
			if err != nil {
				return nil, err
			}
			out.Post = child.(model2.Stmt)
		}

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "For", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["For"]; ok {
				out.Decs.For = decs
			}
			if decs, ok := nd["Init"]; ok {
				out.Decs.Init = decs
			}
			if decs, ok := nd["Cond"]; ok {
				out.Decs.Cond = decs
			}
			if decs, ok := nd["Post"]; ok {
				out.Decs.Post = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.FuncDecl:
		out := &model2.Func{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Init: Type
		out.Type = &model2.FuncType{}
		f.Dst.Nodes[n.Type] = out.Type
		f.Ast.Nodes[out.Type] = n.Type

		// Token: Func
		out.Type.Func = true

		// Node: Recv
		if n.Recv != nil {
			child, err := f.decorateNode(n, "Func", "Recv", "FieldList", n.Recv)
			if err != nil {
				return nil, err
			}
			out.Recv = child.(*model2.FieldList)
		}

		// Node: Name
		if n.Name != nil {
			child, err := f.decorateNode(n, "Func", "Name", "Ident", n.Name)
			if err != nil {
				return nil, err
			}
			out.Name = child.(*model2.Ident)
		}

		// Node: Params
		if n.Type.TypeParams != nil {
			child, err := f.decorateNode(n, "Func", "Params", "FieldList", n.Type.TypeParams)
			if err != nil {
				return nil, err
			}
			out.Type.TypeParams = child.(*model2.FieldList)
		}

		// Node: Params
		if n.Type.Params != nil {
			child, err := f.decorateNode(n, "Func", "Params", "FieldList", n.Type.Params)
			if err != nil {
				return nil, err
			}
			out.Type.Params = child.(*model2.FieldList)
		}

		// Node: Results
		if n.Type.Results != nil {
			child, err := f.decorateNode(n, "Func", "Results", "FieldList", n.Type.Results)
			if err != nil {
				return nil, err
			}
			out.Type.Results = child.(*model2.FieldList)
		}

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "Func", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Func"]; ok {
				out.Decs.Func = decs
			}
			if decs, ok := nd["Recv"]; ok {
				out.Decs.Recv = decs
			}
			if decs, ok := nd["Name"]; ok {
				out.Decs.Name = decs
			}
			if decs, ok := nd["Params"]; ok {
				out.Decs.TypeParams = decs
			}
			if decs, ok := nd["Params"]; ok {
				out.Decs.Params = decs
			}
			if decs, ok := nd["Results"]; ok {
				out.Decs.Results = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.FuncLit:
		out := &model2.FuncLit{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Type
		if n.Type != nil {
			child, err := f.decorateNode(n, "FuncLit", "Type", "FuncType", n.Type)
			if err != nil {
				return nil, err
			}
			out.Type = child.(*model2.FuncType)
		}

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "FuncLit", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Type"]; ok {
				out.Decs.Type = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.FuncType:
		out := &model2.FuncType{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Func
		out.Func = n.Func.IsValid()

		// Node: Params
		if n.TypeParams != nil {
			child, err := f.decorateNode(n, "FuncType", "Params", "FieldList", n.TypeParams)
			if err != nil {
				return nil, err
			}
			out.TypeParams = child.(*model2.FieldList)
		}

		// Node: Params
		if n.Params != nil {
			child, err := f.decorateNode(n, "FuncType", "Params", "FieldList", n.Params)
			if err != nil {
				return nil, err
			}
			out.Params = child.(*model2.FieldList)
		}

		// Node: Results
		if n.Results != nil {
			child, err := f.decorateNode(n, "FuncType", "Results", "FieldList", n.Results)
			if err != nil {
				return nil, err
			}
			out.Results = child.(*model2.FieldList)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Func"]; ok {
				out.Decs.Func = decs
			}
			if decs, ok := nd["Params"]; ok {
				out.Decs.TypeParams = decs
			}
			if decs, ok := nd["Params"]; ok {
				out.Decs.Params = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.GenDecl:
		out := &model2.Gen{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Tok
		out.Tok = n.Tok

		// Token: Lparen
		out.Lparen = n.Lparen.IsValid()

		// List: Specs
		for _, v := range n.Specs {
			child, err := f.decorateNode(n, "Gen", "Specs", "Spec", v)
			if err != nil {
				return nil, err
			}
			out.Specs = append(out.Specs, child.(model2.Spec))
		}

		// Token: Rparen
		out.Rparen = n.Rparen.IsValid()

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Tok"]; ok {
				out.Decs.Tok = decs
			}
			if decs, ok := nd["Lparen"]; ok {
				out.Decs.Lparen = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.GoStmt:
		out := &model2.Go{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Go

		// Node: Call
		if n.Call != nil {
			child, err := f.decorateNode(n, "Go", "Call", "Call", n.Call)
			if err != nil {
				return nil, err
			}
			out.Call = child.(*model2.Call)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Go"]; ok {
				out.Decs.Go = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.Ident:
		out := &model2.Ident{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// String: Name
		out.Name = n.Name

		// Object: Obj
		ob, err := f.decorateObject(n.Obj)
		if err != nil {
			return nil, err
		}
		out.Obj = ob

		// Path: Path
		if f.Resolver != nil {
			path, err := f.resolvePath(false, parent, parentName, parentField, parentFieldType, n)
			if err != nil {
				return nil, err
			}
			out.Path = path
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.IfStmt:
		out := &model2.If{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: If

		// Node: Init
		if n.Init != nil {
			child, err := f.decorateNode(n, "If", "Init", "Stmt", n.Init)
			if err != nil {
				return nil, err
			}
			out.Init = child.(model2.Stmt)
		}

		// Node: Cond
		if n.Cond != nil {
			child, err := f.decorateNode(n, "If", "Cond", "Expr", n.Cond)
			if err != nil {
				return nil, err
			}
			out.Cond = child.(model2.Expr)
		}

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "If", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		// Token: ElseTok

		// Node: Else
		if n.Else != nil {
			child, err := f.decorateNode(n, "If", "Else", "Stmt", n.Else)
			if err != nil {
				return nil, err
			}
			out.Else = child.(model2.Stmt)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["If"]; ok {
				out.Decs.If = decs
			}
			if decs, ok := nd["Init"]; ok {
				out.Decs.Init = decs
			}
			if decs, ok := nd["Cond"]; ok {
				out.Decs.Cond = decs
			}
			if decs, ok := nd["Else"]; ok {
				out.Decs.Else = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.ImportSpec:
		out := &model2.Import{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Name
		if n.Name != nil {
			child, err := f.decorateNode(n, "Import", "Name", "Ident", n.Name)
			if err != nil {
				return nil, err
			}
			out.Name = child.(*model2.Ident)
		}

		// Node: Path
		if n.Path != nil {
			child, err := f.decorateNode(n, "Import", "Path", "Lit", n.Path)
			if err != nil {
				return nil, err
			}
			out.Path = child.(*model2.Lit)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Name"]; ok {
				out.Decs.Name = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.IncDecStmt:
		out := &model2.IncDec{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "IncDec", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Tok
		out.Tok = n.Tok

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.IndexExpr:
		out := &model2.Index{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Index", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Lbrack

		// Node: Index
		if n.Index != nil {
			child, err := f.decorateNode(n, "Index", "Index", "Expr", n.Index)
			if err != nil {
				return nil, err
			}
			out.Index = child.(model2.Expr)
		}

		// Token: Rbrack

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["Lbrack"]; ok {
				out.Decs.Lbrack = decs
			}
			if decs, ok := nd["Index"]; ok {
				out.Decs.Index = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.IndexListExpr:
		out := &model2.IndexList{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "IndexList", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Lbrack

		// List: Indices
		for _, v := range n.Indices {
			child, err := f.decorateNode(n, "IndexList", "Indices", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.Indices = append(out.Indices, child.(model2.Expr))
		}

		// Token: Rbrack

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["Lbrack"]; ok {
				out.Decs.Lbrack = decs
			}
			if decs, ok := nd["Indices"]; ok {
				out.Decs.Indices = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.InterfaceType:
		out := &model2.Interface{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Interface

		// Node: Methods
		if n.Methods != nil {
			child, err := f.decorateNode(n, "Interface", "Methods", "FieldList", n.Methods)
			if err != nil {
				return nil, err
			}
			out.Methods = child.(*model2.FieldList)
		}

		// Value: Ellipsis
		out.Incomplete = n.Incomplete

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Interface"]; ok {
				out.Decs.Interface = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.KeyValueExpr:
		out := &model2.KeyValue{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Key
		if n.Key != nil {
			child, err := f.decorateNode(n, "KeyValue", "Key", "Expr", n.Key)
			if err != nil {
				return nil, err
			}
			out.Key = child.(model2.Expr)
		}

		// Token: Colon

		// Node: Value
		if n.Value != nil {
			child, err := f.decorateNode(n, "KeyValue", "Value", "Expr", n.Value)
			if err != nil {
				return nil, err
			}
			out.Value = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Key"]; ok {
				out.Decs.Key = decs
			}
			if decs, ok := nd["Colon"]; ok {
				out.Decs.Colon = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.LabeledStmt:
		out := &model2.Labeled{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Label
		if n.Label != nil {
			child, err := f.decorateNode(n, "Labeled", "Label", "Ident", n.Label)
			if err != nil {
				return nil, err
			}
			out.Label = child.(*model2.Ident)
		}

		// Token: Colon

		// Node: Stmt
		if n.Stmt != nil {
			child, err := f.decorateNode(n, "Labeled", "Stmt", "Stmt", n.Stmt)
			if err != nil {
				return nil, err
			}
			out.Stmt = child.(model2.Stmt)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Label"]; ok {
				out.Decs.Label = decs
			}
			if decs, ok := nd["Colon"]; ok {
				out.Decs.Colon = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.MapType:
		out := &model2.MapType{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Map

		// Token: Lbrack

		// Node: Key
		if n.Key != nil {
			child, err := f.decorateNode(n, "MapType", "Key", "Expr", n.Key)
			if err != nil {
				return nil, err
			}
			out.Key = child.(model2.Expr)
		}

		// Token: Rbrack

		// Node: Value
		if n.Value != nil {
			child, err := f.decorateNode(n, "MapType", "Value", "Expr", n.Value)
			if err != nil {
				return nil, err
			}
			out.Value = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Map"]; ok {
				out.Decs.Map = decs
			}
			if decs, ok := nd["Key"]; ok {
				out.Decs.Key = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.Package:
		out := &model2.Package{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		// Value: Name
		out.Name = n.Name

		// Scope: Scope
		scope, err := f.decorateScope(n.Scope)
		if err != nil {
			return nil, err
		}
		out.Scope = scope

		// Map: Imports
		out.Imports = map[string]*model2.Object{}
		for k, v := range n.Imports {
			ob, err := f.decorateObject(v)
			if err != nil {
				return nil, err
			}
			out.Imports[k] = ob
		}

		// Map: Files
		out.Files = map[string]*model2.File{}
		for k, v := range n.Files {
			child, err := f.decorateNode(n, "Package", "Files", "File", v)
			if err != nil {
				return nil, err
			}
			out.Files[k] = child.(*model2.File)
		}

		return out, nil
	case *ast.ParenExpr:
		out := &model2.Paren{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Lparen

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Paren", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Rparen

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Lparen"]; ok {
				out.Decs.Lparen = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.RangeStmt:
		out := &model2.Range{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: For

		// Node: Key
		if n.Key != nil {
			child, err := f.decorateNode(n, "Range", "Key", "Expr", n.Key)
			if err != nil {
				return nil, err
			}
			out.Key = child.(model2.Expr)
		}

		// Token: Comma

		// Node: Value
		if n.Value != nil {
			child, err := f.decorateNode(n, "Range", "Value", "Expr", n.Value)
			if err != nil {
				return nil, err
			}
			out.Value = child.(model2.Expr)
		}

		// Token: Tok
		out.Tok = n.Tok

		// Token: Range

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Range", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "Range", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["For"]; ok {
				out.Decs.For = decs
			}
			if decs, ok := nd["Key"]; ok {
				out.Decs.Key = decs
			}
			if decs, ok := nd["Value"]; ok {
				out.Decs.Value = decs
			}
			if decs, ok := nd["Range"]; ok {
				out.Decs.Range = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.ReturnStmt:
		out := &model2.Return{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Return

		// List: Results
		for _, v := range n.Results {
			child, err := f.decorateNode(n, "Return", "Results", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.Results = append(out.Results, child.(model2.Expr))
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Return"]; ok {
				out.Decs.Return = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.SelectStmt:
		out := &model2.Select{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Select

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "Select", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Select"]; ok {
				out.Decs.Select = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.SelectorExpr:

		// Special case for *ast.Selector - replace with Ident if needed
		id, err := f.decorateSelectorExpr(parent, parentName, parentField, parentFieldType, n)
		if err != nil {
			return nil, err
		}
		if id != nil {
			return id, nil
		}

		out := &model2.Selector{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Selector", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Period

		// Node: Sel
		if n.Sel != nil {
			child, err := f.decorateNode(n, "Selector", "Sel", "Ident", n.Sel)
			if err != nil {
				return nil, err
			}
			out.Sel = child.(*model2.Ident)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.SendStmt:
		out := &model2.Send{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Chan
		if n.Chan != nil {
			child, err := f.decorateNode(n, "Send", "Chan", "Expr", n.Chan)
			if err != nil {
				return nil, err
			}
			out.Chan = child.(model2.Expr)
		}

		// Token: Arrow

		// Node: Value
		if n.Value != nil {
			child, err := f.decorateNode(n, "Send", "Value", "Expr", n.Value)
			if err != nil {
				return nil, err
			}
			out.Value = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Chan"]; ok {
				out.Decs.Chan = decs
			}
			if decs, ok := nd["Arrow"]; ok {
				out.Decs.Arrow = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.SliceExpr:
		out := &model2.Slice{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Slice", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Lbrack

		// Node: Low
		if n.Low != nil {
			child, err := f.decorateNode(n, "Slice", "Low", "Expr", n.Low)
			if err != nil {
				return nil, err
			}
			out.Low = child.(model2.Expr)
		}

		// Token: Colon1

		// Node: High
		if n.High != nil {
			child, err := f.decorateNode(n, "Slice", "High", "Expr", n.High)
			if err != nil {
				return nil, err
			}
			out.High = child.(model2.Expr)
		}

		// Token: Colon2

		// Node: Max
		if n.Max != nil {
			child, err := f.decorateNode(n, "Slice", "Max", "Expr", n.Max)
			if err != nil {
				return nil, err
			}
			out.Max = child.(model2.Expr)
		}

		// Token: Rbrack

		// Value: Slice3
		out.Slice3 = n.Slice3

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["Lbrack"]; ok {
				out.Decs.Lbrack = decs
			}
			if decs, ok := nd["Low"]; ok {
				out.Decs.Low = decs
			}
			if decs, ok := nd["High"]; ok {
				out.Decs.High = decs
			}
			if decs, ok := nd["Max"]; ok {
				out.Decs.Max = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.StarExpr:
		out := &model2.Star{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Star

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Star", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Star"]; ok {
				out.Decs.Star = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.StructType:
		out := &model2.Struct{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Struct

		// Node: Methods
		if n.Fields != nil {
			child, err := f.decorateNode(n, "Struct", "Methods", "FieldList", n.Fields)
			if err != nil {
				return nil, err
			}
			out.Fields = child.(*model2.FieldList)
		}

		// Value: Ellipsis
		out.Incomplete = n.Incomplete

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Struct"]; ok {
				out.Decs.Struct = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.SwitchStmt:
		out := &model2.Switch{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Switch

		// Node: Init
		if n.Init != nil {
			child, err := f.decorateNode(n, "Switch", "Init", "Stmt", n.Init)
			if err != nil {
				return nil, err
			}
			out.Init = child.(model2.Stmt)
		}

		// Node: Tag
		if n.Tag != nil {
			child, err := f.decorateNode(n, "Switch", "Tag", "Expr", n.Tag)
			if err != nil {
				return nil, err
			}
			out.Tag = child.(model2.Expr)
		}

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "Switch", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Switch"]; ok {
				out.Decs.Switch = decs
			}
			if decs, ok := nd["Init"]; ok {
				out.Decs.Init = decs
			}
			if decs, ok := nd["Tag"]; ok {
				out.Decs.Tag = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.TypeAssertExpr:
		out := &model2.TypeAssert{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "TypeAssert", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		// Token: Period

		// Token: Lparen

		// Node: Type
		if n.Type != nil {
			child, err := f.decorateNode(n, "TypeAssert", "Type", "Expr", n.Type)
			if err != nil {
				return nil, err
			}
			out.Type = child.(model2.Expr)
		}

		// Token: TypeToken

		// Token: Rparen

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["X"]; ok {
				out.Decs.X = decs
			}
			if decs, ok := nd["Lparen"]; ok {
				out.Decs.Lparen = decs
			}
			if decs, ok := nd["Type"]; ok {
				out.Decs.Type = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.TypeSpec:
		out := &model2.Type{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Node: Name
		if n.Name != nil {
			child, err := f.decorateNode(n, "Type", "Name", "Ident", n.Name)
			if err != nil {
				return nil, err
			}
			out.Name = child.(*model2.Ident)
		}

		// Token: Assign
		out.Assign = n.Assign.IsValid()

		// Node: Params
		if n.TypeParams != nil {
			child, err := f.decorateNode(n, "Type", "Params", "FieldList", n.TypeParams)
			if err != nil {
				return nil, err
			}
			out.Params = child.(*model2.FieldList)
		}

		// Node: Type
		if n.Type != nil {
			child, err := f.decorateNode(n, "Type", "Type", "Expr", n.Type)
			if err != nil {
				return nil, err
			}
			out.Type = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Name"]; ok {
				out.Decs.Name = decs
			}
			if decs, ok := nd["Params"]; ok {
				out.Decs.TypeParams = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.TypeSwitchStmt:
		out := &model2.TypeSwitch{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Switch

		// Node: Init
		if n.Init != nil {
			child, err := f.decorateNode(n, "TypeSwitch", "Init", "Stmt", n.Init)
			if err != nil {
				return nil, err
			}
			out.Init = child.(model2.Stmt)
		}

		// Node: Assign
		if n.Assign != nil {
			child, err := f.decorateNode(n, "TypeSwitch", "Assign", "Stmt", n.Assign)
			if err != nil {
				return nil, err
			}
			out.Assign = child.(model2.Stmt)
		}

		// Node: Body
		if n.Body != nil {
			child, err := f.decorateNode(n, "TypeSwitch", "Body", "Block", n.Body)
			if err != nil {
				return nil, err
			}
			out.Body = child.(*model2.Block)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Switch"]; ok {
				out.Decs.Switch = decs
			}
			if decs, ok := nd["Init"]; ok {
				out.Decs.Init = decs
			}
			if decs, ok := nd["Assign"]; ok {
				out.Decs.Assign = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.UnaryExpr:
		out := &model2.Unary{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// Token: Op
		out.Op = n.Op

		// Node: X
		if n.X != nil {
			child, err := f.decorateNode(n, "Unary", "X", "Expr", n.X)
			if err != nil {
				return nil, err
			}
			out.X = child.(model2.Expr)
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Op"]; ok {
				out.Decs.Op = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	case *ast.ValueSpec:
		out := &model2.Value{}
		f.Dst.Nodes[n] = out
		f.Ast.Nodes[out] = n

		out.Decs.Before = f.before[n]
		out.Decs.After = f.after[n]

		// List: Names
		for _, v := range n.Names {
			child, err := f.decorateNode(n, "Value", "Names", "Ident", v)
			if err != nil {
				return nil, err
			}
			out.Names = append(out.Names, child.(*model2.Ident))
		}

		// Node: Type
		if n.Type != nil {
			child, err := f.decorateNode(n, "Value", "Type", "Expr", n.Type)
			if err != nil {
				return nil, err
			}
			out.Type = child.(model2.Expr)
		}

		// Token: Assign

		// List: Values
		for _, v := range n.Values {
			child, err := f.decorateNode(n, "Value", "Values", "Expr", v)
			if err != nil {
				return nil, err
			}
			out.Values = append(out.Values, child.(model2.Expr))
		}

		if nd, ok := f.decorations[n]; ok {
			if decs, ok := nd["Start"]; ok {
				out.Decs.Start = decs
			}
			if decs, ok := nd["Assign"]; ok {
				out.Decs.Assign = decs
			}
			if decs, ok := nd["End"]; ok {
				out.Decs.End = decs
			}
		}

		return out, nil
	}
	return nil, nil
}
