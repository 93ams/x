package decorator

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper/model"
	"go/ast"
	"go/token"
	"sort"
	"strings"
)

func (f *fileDecorator) addDecorationFragment(n ast.Node, name string, pos token.Pos) {
	f.fragments = append(f.fragments, &decorationFragment{Node: n, Name: name, Pos: token.Pos(f.cursor)})
}
func (f *fileDecorator) addTokenFragment(n ast.Node, t token.Token, pos token.Pos) {
	if pos.IsValid() {
		f.cursor = int(pos)
	}
	f.fragments = append(f.fragments, &tokenFragment{Node: n, Token: t, Pos: token.Pos(f.cursor)})
	f.cursor += len(t.String())
}
func (f *fileDecorator) addStringFragment(n ast.Node, s string, pos token.Pos) {
	if pos.IsValid() {
		f.cursor = int(pos)
	}
	f.fragments = append(f.fragments, &stringFragment{Node: n, String: s, Pos: token.Pos(f.cursor)})
	f.cursor += len(s)
}
func (f *fileDecorator) addBadFragment(n ast.Node, pos token.Pos, length int) {
	if pos.IsValid() {
		f.cursor = int(pos)
	}
	f.fragments = append(f.fragments, &badFragment{Node: n, Pos: token.Pos(f.cursor), Length: length})
	f.cursor += length
}
func (f *fileDecorator) addCommentFragment(text string, pos token.Pos) {
	f.fragments = append(f.fragments, &commentFragment{Text: text, Pos: pos})
}
func (f *fileDecorator) addNewlineFragment(pos token.Pos, empty bool) {
	f.fragments = append(f.fragments, &newlineFragment{Pos: pos, Empty: empty})
}
func (f *fileDecorator) fragment(node ast.Node) {
	f.addNodeFragments(node)
	if f.Fset != nil {
		processFile := func(astf *ast.File) {
			avoid := map[int]bool{}
			for _, cg := range astf.Comments {
				for _, c := range cg.List {
					f.addCommentFragment(c.Text, c.Slash)
					if strings.HasPrefix(c.Text, "/*") {
						startLine := f.Fset.Position(c.Pos()).Line
						endLine := f.Fset.Position(c.End()).Line
						if endLine > startLine {
							for i := startLine; i < endLine; i++ {
								avoid[i+1] = true
							}
						}
					}
				}
			}
			for _, frag := range f.fragments {
				switch frag := frag.(type) {
				case *stringFragment:
					if !strings.HasPrefix(frag.String, "`") {
						continue
					}
					startLine := f.Fset.Position(frag.Pos).Line
					endLine := f.Fset.Position(frag.Pos + token.Pos(len(frag.String))).Line
					if endLine > startLine {
						for i := startLine; i < endLine; i++ {
							avoid[i+1] = true
						}
					}
				case *badFragment:
					startLine := f.Fset.Position(frag.Pos).Line
					endLine := f.Fset.Position(frag.Pos + token.Pos(frag.Length)).Line
					if endLine > startLine {
						for i := startLine; i < endLine; i++ {
							avoid[i+1] = true
						}
					}
				}
			}
			line := 1
			tokenf := f.Fset.File(astf.Pos())
			max := tokenf.Base() + tokenf.Size()
			for i := tokenf.Base(); i < max; i++ {
				pos := f.Fset.Position(token.Pos(i))
				if pos.Line != line {
					line = pos.Line
					if avoid[line] {
						continue
					}
					nextLine := line
					if i < max-1 {
						nextLine = f.Fset.Position(token.Pos(i + 1)).Line
					}
					if nextLine != line {
						f.addNewlineFragment(token.Pos(i-1), true)
						line = nextLine
						i++
					} else {
						f.addNewlineFragment(token.Pos(i-1), false)
					}
				}
			}
		}
		switch val := node.(type) {
		case *ast.File:
			processFile(val)
		case *ast.Package:
			for _, file := range val.Files {
				processFile(file)
			}
		}
	}
	sort.SliceStable(f.fragments, func(i, j int) bool {
		return f.fragments[i].Position() < f.fragments[j].Position()
	})
	currentIndent := 0
	for i, frag := range f.fragments {
		if i == 0 || f.fragments[i-1].Newline() {
			currentIndent = f.Fset.Position(frag.Position()).Column
		}
		switch frag := frag.(type) {
		case *decorationFragment:
			switch frag.Name {
			case "Start":
				f.startIndents[frag.Node] = currentIndent
			case "End":
				f.endIndents[frag.Node] = currentIndent
			}
		case *commentFragment:
			frag.Indent = currentIndent
		}
	}
}
func (f *fileDecorator) link() {
	for i, frag := range f.fragments {
		switch frag := frag.(type) {
		case *decorationFragment:
			if frag.Name != "End" {
				continue
			}
			_, stmt := frag.Node.(ast.Stmt)
			_, decl := frag.Node.(ast.Decl)
			if !stmt && !decl {
				continue
			}
			if _, labeledStmt := frag.Node.(*ast.LabeledStmt); labeledStmt {
				continue
			}
			start := f.startIndents[frag.Node]
			end := f.endIndents[frag.Node]
			_, caseClause := frag.Node.(*ast.CaseClause)
			_, commClause := frag.Node.(*ast.CommClause)
			if start == end && (caseClause || commClause) {
				end++
			}
			if end != start+1 {
				continue
			}
			frags, next := f.findIndentedComments(i+1, [2]int{end, start})
			endFrags := frags[0]
			nextFrags := frags[1]
			if len(endFrags) > 0 {
				_, nl := endFrags[len(endFrags)-1].(*newlineFragment)
				if nl {
					f.attachToDecoration(endFrags[0:len(endFrags)-1], f.decorations, frag)
				} else {
					f.attachToDecoration(endFrags, f.decorations, frag)
				}
			}
			if len(nextFrags) > 0 && next != nil {
				_, nextStmt := next.Node.(ast.Stmt)
				_, nextDecl := next.Node.(ast.Decl)
				nextStart := f.startIndents[next.Node]
				if (nextStmt || nextDecl) && nextStart == start {
					f.attachToDecoration(nextFrags, f.decorations, next)
				}
			}
		case *commentFragment:
			if frag.Attached != nil {
				continue
			}
			var frags []fragment // comment / new-line / empty-line
			var dec *decorationFragment
			var found bool
			var try int
			for !found {
				try++
				switch try {
				case 1:
					frags, dec, found = f.findDecoration(true, true, i, -1, false)
				case 2:
					frags, dec, found = f.findDecoration(false, true, i, 1, false)
				case 3:
					frags, dec, found = f.findDecoration(false, true, i, -1, false)
				case 4:
					frags, dec, found = f.findDecoration(false, false, i, 1, false)
				case 5:
					frags, dec, found = f.findDecoration(false, false, i, -1, false)
				default:
					panic("no decoration found for " + frag.Text)
				}
			}
			f.attachToDecoration(frags, f.decorations, dec)
		}
	}
	for i, frag := range f.fragments {
		switch frag := frag.(type) {
		case *newlineFragment:
			if frag.Attached != nil {
				continue
			}
			nodeBefore, _, foundBefore := f.findNode(i, 1)
			nodeAfter, _, foundAfter := f.findNode(i, -1)
			if foundBefore || foundAfter {
				spaceType := model.NewLine
				if frag.Empty {
					spaceType = model.EmptyLine
				}
				if foundBefore {
					f.before[nodeBefore] = spaceType
				}
				if foundAfter {
					f.after[nodeAfter] = spaceType
				}
				continue
			}
			var dec *decorationFragment
			var found bool
			var try int
			for !found {
				try++
				switch try {
				case 1:
					_, dec, found = f.findDecoration(false, false, i, -1, false)
				case 2:
					_, dec, found = f.findDecoration(false, false, i, 1, false)
				default:
					panic("no decoration found for newline")
				}
			}
			appendNewLine(f.decorations, dec.Node, dec.Name, frag.Empty)
		}
	}
	return
}

func appendDecoration(m map[ast.Node]map[string][]string, n ast.Node, pos, text string) {
	if m[n] == nil {
		m[n] = map[string][]string{}
	}
	m[n][pos] = append(m[n][pos], text)
}

func appendNewLine(m map[ast.Node]map[string][]string, n ast.Node, pos string, empty bool) {
	if m[n] == nil {
		m[n] = map[string][]string{}
	}
	num := 1
	if empty {
		num = 2
	}
	decs := m[n][pos]
	if len(decs) > 0 && strings.HasPrefix(decs[len(decs)-1], "//") {
		num--
	}
	for i := 0; i < num; i++ {
		m[n][pos] = append(m[n][pos], "\n")
	}
}

func (f *fileDecorator) attachToDecoration(frags []fragment, decorations map[ast.Node]map[string][]string, dec *decorationFragment) {
	for _, fr := range frags {
		switch fr := fr.(type) {
		case *commentFragment:
			appendDecoration(decorations, dec.Node, dec.Name, fr.Text)
			fr.Attached = dec
		case *newlineFragment:
			appendNewLine(decorations, dec.Node, dec.Name, fr.Empty)
			fr.Attached = dec
		}
	}
}

func (f *fileDecorator) findDecoration(stopAtNewline, stopAtEmptyLine bool, from int, direction int, onlyClause bool) (swept []fragment, dec *decorationFragment, found bool) {
	var frags []fragment
	for i := from; i < len(f.fragments) && i >= 0; i += direction {
		switch current := f.fragments[i].(type) {
		case *decorationFragment:
			if onlyClause {
				switch current.Node.(type) {
				case *ast.CommClause, *ast.CaseClause:
					if current.Name == "Start" {
						return frags, current, true
					}
					return
				default:
					return
				}
			}
			return frags, current, true
		case *newlineFragment:
			if stopAtNewline {
				return
			}
			if stopAtEmptyLine && current.Empty {
				return
			}
			if current.Attached != nil {
				continue
			}
			if direction == 1 {
				frags = append(frags, current)
			} else {
				frags = append([]fragment{current}, frags...)
			}
		case *commentFragment:
			if current.Attached != nil {
				continue
			}
			if direction == 1 {
				frags = append(frags, current)
			} else {
				frags = append([]fragment{current}, frags...)
			}
		case *tokenFragment, *stringFragment:
			return
		}
	}
	return
}

func (f *fileDecorator) findNode(from int, direction int) (node ast.Node, dec *decorationFragment, found bool) {

	var name string
	switch direction {
	case 1:
		name = "Start"
	case -1:
		name = "End"
	}

	for i := from; i < len(f.fragments) && i >= 0; i += direction {
		switch frag := f.fragments[i].(type) {
		case *decorationFragment:
			if frag.Name == name {
				return frag.Node, frag, true
			}
			return
		case *commentFragment:
			if frag.Attached != nil && frag.Attached.Name == name {
				return frag.Attached.Node, frag.Attached, true
			}
		case *newlineFragment:
			if frag.Attached != nil && frag.Attached.Name == name {
				return frag.Attached.Node, frag.Attached, true
			}
		case *tokenFragment, *stringFragment:
			return
		}
	}
	return
}

func (f *fileDecorator) findIndentedComments(from int, indents [2]int) (frags [2][]fragment, nextDecoration *decorationFragment) {
	var stage int
	var pastNewline bool // while this is false, we're on the same line that the stmt ended, so we accept all comments regardless of the indent (e.g. empty clauses) - see "hanging-indent-same-line" test case.
	for i := from; i < len(f.fragments); i++ {
		switch current := f.fragments[i].(type) {
		case *decorationFragment:
			return frags, current
		case *newlineFragment:
			pastNewline = true
			frags[stage] = append(frags[stage], current)
		case *commentFragment:
			if !pastNewline {
				frags[stage] = append(frags[stage], current)
				continue
			}
			if stage == 0 {
				// Check indent matches. If not, move to second stage or exit if that doesn't match.
				if current.Indent != indents[0] {
					if current.Indent == indents[1] {
						stage = 1
					} else {
						return
					}
				}
			} else if stage == 1 {
				if current.Indent != indents[1] {
					return
				}
			}
			frags[stage] = append(frags[stage], current)
		case *tokenFragment, *stringFragment:
			return
		}
	}
	return
}

type fragment interface {
	Position() token.Pos
	Newline() bool // True if the fragment ends in a newline ("\n" or "//...")
}

type tokenFragment struct {
	Node  ast.Node
	Token token.Token
	Pos   token.Pos
}

type stringFragment struct {
	Node   ast.Node
	String string
	Pos    token.Pos
}

type badFragment struct {
	Node   ast.Node
	Pos    token.Pos
	Length int
}

type commentFragment struct {
	Text     string
	Pos      token.Pos
	Attached *decorationFragment
	Indent   int
}

type newlineFragment struct {
	Pos      token.Pos
	Empty    bool
	Attached *decorationFragment
}

type decorationFragment struct {
	Node ast.Node
	Name string
	Pos  token.Pos
}

func (v *tokenFragment) Position() token.Pos      { return v.Pos }
func (v *stringFragment) Position() token.Pos     { return v.Pos }
func (v *commentFragment) Position() token.Pos    { return v.Pos }
func (v *newlineFragment) Position() token.Pos    { return v.Pos }
func (v *decorationFragment) Position() token.Pos { return v.Pos }
func (v *badFragment) Position() token.Pos        { return v.Pos }

func (v *tokenFragment) Newline() bool      { return false }
func (v *stringFragment) Newline() bool     { return false }
func (v *commentFragment) Newline() bool    { return strings.HasPrefix(v.Text, "//") }
func (v *newlineFragment) Newline() bool    { return true }
func (v *decorationFragment) Newline() bool { return false }
func (v *badFragment) Newline() bool        { return false }

func (f *fileDecorator) addNodeFragments(n ast.Node) {
	if n.Pos().IsValid() {
		f.cursor = int(n.Pos())
	}
	switch n := n.(type) {
	case *ast.ArrayType:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.LBRACK, n.Lbrack)
		f.addDecorationFragment(n, "Lbrack", token.NoPos)
		if n.Len != nil {
			f.addNodeFragments(n.Len)
		}
		f.addTokenFragment(n, token.RBRACK, token.NoPos)
		f.addDecorationFragment(n, "Len", token.NoPos)
		if n.Elt != nil {
			f.addNodeFragments(n.Elt)
		}
		f.addDecorationFragment(n, "End", n.End())
	case *ast.AssignStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		for _, v := range n.Lhs {
			f.addNodeFragments(v)
		}
		f.addTokenFragment(n, n.Tok, n.TokPos)
		f.addDecorationFragment(n, "Tok", token.NoPos)
		for _, v := range n.Rhs {
			f.addNodeFragments(v)
		}
		f.addDecorationFragment(n, "End", n.End())
	case *ast.BadDecl:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addBadFragment(n, n.From, int(n.To-n.From))
		f.addDecorationFragment(n, "End", n.End())
	case *ast.BadExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addBadFragment(n, n.From, int(n.To-n.From))
		f.addDecorationFragment(n, "End", n.End())

	case *ast.BadStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addBadFragment(n, n.From, int(n.To-n.From))
		f.addDecorationFragment(n, "End", n.End())

	case *ast.BasicLit:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addStringFragment(n, n.Value, n.ValuePos)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.BinaryExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addTokenFragment(n, n.Op, n.OpPos)
		f.addDecorationFragment(n, "Op", token.NoPos)
		if n.Y != nil {
			f.addNodeFragments(n.Y)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.BlockStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.LBRACE, n.Lbrace)
		f.addDecorationFragment(n, "Lbrace", token.NoPos)
		for _, v := range n.List {
			f.addNodeFragments(v)
		}
		f.addTokenFragment(n, token.RBRACE, n.Rbrace)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.BranchStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, n.Tok, n.TokPos)
		if n.Label != nil {
			f.addDecorationFragment(n, "Tok", token.NoPos)
		}
		if n.Label != nil {
			f.addNodeFragments(n.Label)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.CallExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Fun != nil {
			f.addNodeFragments(n.Fun)
		}
		f.addDecorationFragment(n, "Fun", token.NoPos)
		f.addTokenFragment(n, token.LPAREN, n.Lparen)
		f.addDecorationFragment(n, "Lparen", token.NoPos)
		for _, v := range n.Args {
			f.addNodeFragments(v)
		}
		if n.Ellipsis.IsValid() {
			f.addTokenFragment(n, token.ELLIPSIS, n.Ellipsis)
		}
		if n.Ellipsis.IsValid() {
			f.addDecorationFragment(n, "Ellipsis", token.NoPos)
		}
		f.addTokenFragment(n, token.RPAREN, n.Rparen)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.CaseClause:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, func() token.Token {
			if n.List == nil {
				return token.DEFAULT
			}
			return token.CASE
		}(), n.Case)
		f.addDecorationFragment(n, "Case", token.NoPos)
		for _, v := range n.List {
			f.addNodeFragments(v)
		}
		f.addTokenFragment(n, token.COLON, n.Colon)
		f.addDecorationFragment(n, "Colon", token.NoPos)
		for _, v := range n.Body {
			f.addNodeFragments(v)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.ChanType:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, func() token.Token {
			if n.Dir == ast.RECV {
				return token.ARROW
			}
			return token.CHAN
		}(), n.Begin)
		if n.Dir == ast.RECV {
			f.addTokenFragment(n, token.CHAN, token.NoPos)
		}
		f.addDecorationFragment(n, "Begin", token.NoPos)
		if n.Dir == ast.SEND {
			f.addTokenFragment(n, token.ARROW, n.Arrow)
		}
		if n.Dir == ast.SEND {
			f.addDecorationFragment(n, "Arrow", token.NoPos)
		}
		if n.Value != nil {
			f.addNodeFragments(n.Value)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.CommClause:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, func() token.Token {
			if n.Comm == nil {
				return token.DEFAULT
			}
			return token.CASE
		}(), n.Case)
		f.addDecorationFragment(n, "Case", token.NoPos)
		if n.Comm != nil {
			f.addNodeFragments(n.Comm)
		}
		if n.Comm != nil {
			f.addDecorationFragment(n, "Comm", token.NoPos)
		}
		f.addTokenFragment(n, token.COLON, n.Colon)
		f.addDecorationFragment(n, "Colon", token.NoPos)
		for _, v := range n.Body {
			f.addNodeFragments(v)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.CompositeLit:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Type != nil {
			f.addNodeFragments(n.Type)
		}
		if n.Type != nil {
			f.addDecorationFragment(n, "Type", token.NoPos)
		}
		f.addTokenFragment(n, token.LBRACE, n.Lbrace)
		f.addDecorationFragment(n, "Lbrace", token.NoPos)
		for _, v := range n.Elts {
			f.addNodeFragments(v)
		}
		f.addTokenFragment(n, token.RBRACE, n.Rbrace)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.DeclStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Decl != nil {
			f.addNodeFragments(n.Decl)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.DeferStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.DEFER, n.Defer)
		f.addDecorationFragment(n, "Defer", token.NoPos)
		if n.Call != nil {
			f.addNodeFragments(n.Call)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.Ellipsis:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.ELLIPSIS, n.Ellipsis)
		if n.Elt != nil {
			f.addDecorationFragment(n, "Ellipsis", token.NoPos)
		}
		if n.Elt != nil {
			f.addNodeFragments(n.Elt)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.EmptyStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		if !n.Implicit {
			f.addTokenFragment(n, token.ARROW, n.Semicolon)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.ExprStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.Field:
		f.addDecorationFragment(n, "Start", n.Pos())
		for _, v := range n.Names {
			f.addNodeFragments(v)
		}
		if n.Type != nil {
			f.addNodeFragments(n.Type)
		}
		if n.Tag != nil {
			f.addDecorationFragment(n, "Type", token.NoPos)
		}
		if n.Tag != nil {
			f.addNodeFragments(n.Tag)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.FieldList:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Opening.IsValid() {
			f.addTokenFragment(n, token.LPAREN, n.Opening)
		}
		f.addDecorationFragment(n, "Opening", token.NoPos)
		for _, v := range n.List {
			f.addNodeFragments(v)
		}
		if n.Closing.IsValid() {
			f.addTokenFragment(n, token.RPAREN, n.Closing)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.File:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.PACKAGE, n.Package)
		f.addDecorationFragment(n, "Package", token.NoPos)
		if n.Name != nil {
			f.addNodeFragments(n.Name)
		}
		f.addDecorationFragment(n, "Name", token.NoPos)
		for _, v := range n.Decls {
			f.addNodeFragments(v)
		}
		for _, v := range n.Imports {
			f.addNodeFragments(v)
		}

	case *ast.ForStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.FOR, n.For)
		f.addDecorationFragment(n, "For", token.NoPos)
		if n.Init != nil {
			f.addNodeFragments(n.Init)
		}
		if n.Init != nil {
			f.addTokenFragment(n, token.SEMICOLON, token.NoPos)
		}
		if n.Init != nil {
			f.addDecorationFragment(n, "Init", token.NoPos)
		}
		if n.Cond != nil {
			f.addNodeFragments(n.Cond)
		}
		if n.Post != nil {
			f.addTokenFragment(n, token.SEMICOLON, token.NoPos)
		}
		if n.Cond != nil {
			f.addDecorationFragment(n, "Cond", token.NoPos)
		}
		if n.Post != nil {
			f.addNodeFragments(n.Post)
		}
		if n.Post != nil {
			f.addDecorationFragment(n, "Post", token.NoPos)
		}
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.FuncDecl:
		f.addDecorationFragment(n, "Start", n.Pos())
		if true {
			f.addTokenFragment(n, token.FUNC, n.Type.Func)
		}
		f.addDecorationFragment(n, "Func", token.NoPos)
		if n.Recv != nil {
			f.addNodeFragments(n.Recv)
		}
		if n.Recv != nil {
			f.addDecorationFragment(n, "Recv", token.NoPos)
		}
		if n.Name != nil {
			f.addNodeFragments(n.Name)
		}
		f.addDecorationFragment(n, "Name", token.NoPos)
		if n.Type.TypeParams != nil {
			f.addNodeFragments(n.Type.TypeParams)
		}
		if n.Type.TypeParams != nil {
			f.addDecorationFragment(n, "Params", token.NoPos)
		}
		if n.Type.Params != nil {
			f.addNodeFragments(n.Type.Params)
		}
		f.addDecorationFragment(n, "Params", token.NoPos)
		if n.Type.Results != nil {
			f.addNodeFragments(n.Type.Results)
		}
		if n.Type.Results != nil {
			f.addDecorationFragment(n, "Results", token.NoPos)
		}
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.FuncLit:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Type != nil {
			f.addNodeFragments(n.Type)
		}
		f.addDecorationFragment(n, "Type", token.NoPos)
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.FuncType:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Func.IsValid() {
			f.addTokenFragment(n, token.FUNC, n.Func)
		}
		if n.Func.IsValid() {
			f.addDecorationFragment(n, "Func", token.NoPos)
		}
		if n.TypeParams != nil {
			f.addNodeFragments(n.TypeParams)
		}
		if n.TypeParams != nil {
			f.addDecorationFragment(n, "Params", token.NoPos)
		}
		if n.Params != nil {
			f.addNodeFragments(n.Params)
		}
		if n.Results != nil {
			f.addDecorationFragment(n, "Params", token.NoPos)
		}
		if n.Results != nil {
			f.addNodeFragments(n.Results)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.GenDecl:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, n.Tok, n.TokPos)
		f.addDecorationFragment(n, "Tok", token.NoPos)
		if n.Lparen.IsValid() {
			f.addTokenFragment(n, token.LPAREN, n.Lparen)
		}
		if n.Lparen.IsValid() {
			f.addDecorationFragment(n, "Lparen", token.NoPos)
		}
		for _, v := range n.Specs {
			f.addNodeFragments(v)
		}
		if n.Rparen.IsValid() {
			f.addTokenFragment(n, token.RPAREN, n.Rparen)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.GoStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.GO, n.Go)
		f.addDecorationFragment(n, "Go", token.NoPos)
		if n.Call != nil {
			f.addNodeFragments(n.Call)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.Ident:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addStringFragment(n, n.Name, n.NamePos)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.IfStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.IF, n.If)
		f.addDecorationFragment(n, "If", token.NoPos)
		if n.Init != nil {
			f.addNodeFragments(n.Init)
		}
		if n.Init != nil {
			f.addDecorationFragment(n, "Init", token.NoPos)
		}
		if n.Cond != nil {
			f.addNodeFragments(n.Cond)
		}
		f.addDecorationFragment(n, "Cond", token.NoPos)
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		if n.Else != nil {
			f.addTokenFragment(n, token.ELSE, token.NoPos)
		}
		if n.Else != nil {
			f.addDecorationFragment(n, "Else", token.NoPos)
		}
		if n.Else != nil {
			f.addNodeFragments(n.Else)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.ImportSpec:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Name != nil {
			f.addNodeFragments(n.Name)
		}
		if n.Name != nil {
			f.addDecorationFragment(n, "Name", token.NoPos)
		}
		if n.Path != nil {
			f.addNodeFragments(n.Path)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.IncDecStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addTokenFragment(n, n.Tok, n.TokPos)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.IndexExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addTokenFragment(n, token.LBRACK, n.Lbrack)
		f.addDecorationFragment(n, "Lbrack", token.NoPos)
		if n.Index != nil {
			f.addNodeFragments(n.Index)
		}
		f.addDecorationFragment(n, "Index", token.NoPos)
		f.addTokenFragment(n, token.RBRACK, n.Rbrack)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.IndexListExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addTokenFragment(n, token.LBRACK, n.Lbrack)
		f.addDecorationFragment(n, "Lbrack", token.NoPos)
		for _, v := range n.Indices {
			f.addNodeFragments(v)
		}
		f.addDecorationFragment(n, "Indices", token.NoPos)
		f.addTokenFragment(n, token.RBRACK, n.Rbrack)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.InterfaceType:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.INTERFACE, n.Interface)
		f.addDecorationFragment(n, "Interface", token.NoPos)
		if n.Methods != nil {
			f.addNodeFragments(n.Methods)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.KeyValueExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Key != nil {
			f.addNodeFragments(n.Key)
		}
		f.addDecorationFragment(n, "Key", token.NoPos)
		f.addTokenFragment(n, token.COLON, n.Colon)
		f.addDecorationFragment(n, "Colon", token.NoPos)
		if n.Value != nil {
			f.addNodeFragments(n.Value)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.LabeledStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Label != nil {
			f.addNodeFragments(n.Label)
		}
		f.addDecorationFragment(n, "Label", token.NoPos)
		f.addTokenFragment(n, token.COLON, n.Colon)
		f.addDecorationFragment(n, "Colon", token.NoPos)
		if n.Stmt != nil {
			f.addNodeFragments(n.Stmt)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.MapType:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.MAP, n.Map)
		f.addTokenFragment(n, token.LBRACK, token.NoPos)
		f.addDecorationFragment(n, "Map", token.NoPos)
		if n.Key != nil {
			f.addNodeFragments(n.Key)
		}
		f.addTokenFragment(n, token.RBRACK, token.NoPos)
		f.addDecorationFragment(n, "Key", token.NoPos)
		if n.Value != nil {
			f.addNodeFragments(n.Value)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.Package:
		for _, v := range n.Files {
			f.addNodeFragments(v)
		}

	case *ast.ParenExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.LPAREN, n.Lparen)
		f.addDecorationFragment(n, "Lparen", token.NoPos)
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addTokenFragment(n, token.RPAREN, n.Rparen)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.RangeStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.FOR, n.For)
		if n.Key != nil {
			f.addDecorationFragment(n, "For", token.NoPos)
		}
		if n.Key != nil {
			f.addNodeFragments(n.Key)
		}
		if n.Value != nil {
			f.addTokenFragment(n, token.COMMA, token.NoPos)
		}
		if n.Key != nil {
			f.addDecorationFragment(n, "Key", token.NoPos)
		}
		if n.Value != nil {
			f.addNodeFragments(n.Value)
		}
		if n.Value != nil {
			f.addDecorationFragment(n, "Value", token.NoPos)
		}
		if n.Tok != token.ILLEGAL {
			f.addTokenFragment(n, n.Tok, n.TokPos)
		}
		f.addTokenFragment(n, token.RANGE, token.NoPos)
		f.addDecorationFragment(n, "Range", token.NoPos)
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "X", token.NoPos)
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.ReturnStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.RETURN, n.Return)
		f.addDecorationFragment(n, "Return", token.NoPos)
		for _, v := range n.Results {
			f.addNodeFragments(v)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.SelectStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.SELECT, n.Select)
		f.addDecorationFragment(n, "Select", token.NoPos)
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.SelectorExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addTokenFragment(n, token.PERIOD, token.NoPos)
		f.addDecorationFragment(n, "X", token.NoPos)
		if n.Sel != nil {
			f.addNodeFragments(n.Sel)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.SendStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Chan != nil {
			f.addNodeFragments(n.Chan)
		}
		f.addDecorationFragment(n, "Chan", token.NoPos)
		f.addTokenFragment(n, token.ARROW, n.Arrow)
		f.addDecorationFragment(n, "Arrow", token.NoPos)
		if n.Value != nil {
			f.addNodeFragments(n.Value)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.SliceExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addTokenFragment(n, token.LBRACK, n.Lbrack)
		if n.Low != nil {
			f.addDecorationFragment(n, "Lbrack", token.NoPos)
		}
		if n.Low != nil {
			f.addNodeFragments(n.Low)
		}
		f.addTokenFragment(n, token.COLON, token.NoPos)
		f.addDecorationFragment(n, "Low", token.NoPos)
		if n.High != nil {
			f.addNodeFragments(n.High)
		}
		if n.Slice3 {
			f.addTokenFragment(n, token.COLON, token.NoPos)
		}
		if n.High != nil {
			f.addDecorationFragment(n, "High", token.NoPos)
		}
		if n.Max != nil {
			f.addNodeFragments(n.Max)
		}
		if n.Max != nil {
			f.addDecorationFragment(n, "Max", token.NoPos)
		}
		f.addTokenFragment(n, token.RBRACK, n.Rbrack)
		f.addDecorationFragment(n, "End", n.End())

	case *ast.StarExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.MUL, n.Star)
		f.addDecorationFragment(n, "Star", token.NoPos)
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.StructType:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.STRUCT, n.Struct)
		f.addDecorationFragment(n, "Struct", token.NoPos)
		if n.Fields != nil {
			f.addNodeFragments(n.Fields)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.SwitchStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.SWITCH, n.Switch)
		f.addDecorationFragment(n, "Switch", token.NoPos)
		if n.Init != nil {
			f.addNodeFragments(n.Init)
		}
		if n.Init != nil {
			f.addDecorationFragment(n, "Init", token.NoPos)
		}
		if n.Tag != nil {
			f.addNodeFragments(n.Tag)
		}
		if n.Tag != nil {
			f.addDecorationFragment(n, "Tag", token.NoPos)
		}
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.TypeAssertExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addTokenFragment(n, token.PERIOD, token.NoPos)
		f.addDecorationFragment(n, "X", token.NoPos)
		f.addTokenFragment(n, token.LPAREN, n.Lparen)
		f.addDecorationFragment(n, "Lparen", token.NoPos)
		if n.Type != nil {
			f.addNodeFragments(n.Type)
		}
		if n.Type == nil {
			f.addTokenFragment(n, token.TYPE, token.NoPos)
		}
		f.addDecorationFragment(n, "Type", token.NoPos)
		f.addTokenFragment(n, token.RPAREN, n.Rparen)
		f.addDecorationFragment(n, "End", n.End())
	case *ast.TypeSpec:
		f.addDecorationFragment(n, "Start", n.Pos())
		if n.Name != nil {
			f.addNodeFragments(n.Name)
		}
		if n.Assign.IsValid() {
			f.addTokenFragment(n, token.ASSIGN, n.Assign)
		}
		f.addDecorationFragment(n, "Name", token.NoPos)
		if n.TypeParams != nil {
			f.addNodeFragments(n.TypeParams)
		}
		if n.TypeParams != nil {
			f.addDecorationFragment(n, "Params", token.NoPos)
		}
		if n.Type != nil {
			f.addNodeFragments(n.Type)
		}
		f.addDecorationFragment(n, "End", n.End())

	case *ast.TypeSwitchStmt:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, token.SWITCH, n.Switch)
		f.addDecorationFragment(n, "Switch", token.NoPos)
		if n.Init != nil {
			f.addNodeFragments(n.Init)
		}
		if n.Init != nil {
			f.addDecorationFragment(n, "Init", token.NoPos)
		}
		if n.Assign != nil {
			f.addNodeFragments(n.Assign)
		}
		f.addDecorationFragment(n, "Assign", token.NoPos)
		if n.Body != nil {
			f.addNodeFragments(n.Body)
		}
		f.addDecorationFragment(n, "End", n.End())
	case *ast.UnaryExpr:
		f.addDecorationFragment(n, "Start", n.Pos())
		f.addTokenFragment(n, n.Op, n.OpPos)
		f.addDecorationFragment(n, "Op", token.NoPos)
		if n.X != nil {
			f.addNodeFragments(n.X)
		}
		f.addDecorationFragment(n, "End", n.End())
	case *ast.ValueSpec:
		f.addDecorationFragment(n, "Start", n.Pos())
		for _, v := range n.Names {
			f.addNodeFragments(v)
		}
		if n.Type != nil {
			f.addNodeFragments(n.Type)
		}
		if n.Values != nil {
			f.addTokenFragment(n, token.ASSIGN, token.NoPos)
		}
		if n.Values != nil {
			f.addDecorationFragment(n, "Assign", token.NoPos)
		}
		for _, v := range n.Values {
			f.addNodeFragments(v)
		}
		f.addDecorationFragment(n, "End", n.End())
	}
}
