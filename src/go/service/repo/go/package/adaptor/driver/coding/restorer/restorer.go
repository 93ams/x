package restorer

import (
	"fmt"
	resolver2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/coding/resolver"
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	Restorer struct {
		Resolver resolver2.RestorerResolver
		Fset     *token.FileSet
		Path     string
		Extras   bool
		model2.Map
	}
	FileRestorer struct {
		*Restorer
		Alias           map[string]string
		Name            string
		file            *model2.File
		lines           []int
		comments        []*ast.CommentGroup
		base            int
		cursor          token.Pos
		nodeDecl        map[*ast.Object]model2.Node
		nodeData        map[*ast.Object]model2.Node
		cursorAtNewLine token.Pos
		packageNames    map[string]string
	}
)

func NewRestorer() *Restorer {
	return &Restorer{Map: model2.NewMap(), Fset: token.NewFileSet(), Resolver: resolver2.NewGuessResolver(), Path: "root"}
}
func NewRestorerWithImports(path string, resolver resolver2.RestorerResolver) *Restorer {
	res := NewRestorer()
	res.Path = path
	res.Resolver = resolver
	return res
}
func (pr *Restorer) Print(f *model2.File) error { return pr.Fprint(os.Stdout, f) }
func (pr *Restorer) Fprint(w io.Writer, f *model2.File) error {
	af, err := pr.RestoreFile(f)
	if err != nil {
		return err
	}
	return format.Node(w, pr.Fset, af)
}
func (pr *Restorer) RestoreFile(file *model2.File) (*ast.File, error) {
	return pr.FileRestorer().RestoreFile(file)
}
func (pr *Restorer) FileRestorer() *FileRestorer {
	return &FileRestorer{Restorer: pr, Alias: map[string]string{}}
}

func (r *FileRestorer) Print(f *model2.File) error { return r.Fprint(os.Stdout, f) }
func (r *FileRestorer) Fprint(w io.Writer, f *model2.File) error {
	af, err := r.RestoreFile(f)
	if err != nil {
		return err
	}
	return format.Node(w, r.Fset, af)
}
func (r *FileRestorer) RestoreFile(file *model2.File) (*ast.File, error) {
	if r.Resolver == nil && r.Path != "" {
		panic("Restorer Path should be empty when Resolver is nil")
	}
	if r.Resolver != nil && r.Path == "" {
		panic("Restorer Path should be set when Resolver is set")
	}
	if r.Fset == nil {
		r.Fset = token.NewFileSet()
	}
	r.file = file
	r.lines = []int{0}
	r.nodeDecl = map[*ast.Object]model2.Node{}
	r.nodeData = map[*ast.Object]model2.Node{}
	r.packageNames = map[string]string{}
	r.comments = []*ast.CommentGroup{}
	r.cursorAtNewLine = 0
	r.packageNames = map[string]string{}
	r.base = r.Fset.Base()
	r.cursor = token.Pos(r.base)
	if err := r.updateImports(); err != nil {
		return nil, err
	}
	f := r.restoreNode(r.file, "", "", "", false).(*ast.File)
	for _, cg := range r.comments {
		f.Comments = append(f.Comments, cg)
	}
	ff := r.Fset.AddFile(r.Name, r.base, r.fileSize())
	if !ff.SetLines(r.lines) {
		panic("ff.SetLines failed")
	}
	if r.Extras {
		for o, dn := range r.nodeDecl {
			o.Decl = r.restoreNode(dn, "", "", "", true)
		}
		for o, dn := range r.nodeData {
			o.Data = r.restoreNode(dn, "", "", "", true)
		}
	}
	return f, nil
}

func (r *FileRestorer) updateImports() error {
	if r.Resolver == nil {
		return nil
	}
	var blocks []*model2.Gen
	var hasCgoBlock bool
	importsFound := map[string]string{}
	packagesInUse := map[string]bool{}
	importsRequired := map[string]bool{}
	model2.Inspect(r.file, func(n model2.Node) bool {
		switch n := n.(type) {
		case *model2.Ident:
			if n.Path == "" {
				return true
			}
			if n.Path == r.Path {
				return true
			}
			packagesInUse[n.Path] = true
			importsRequired[n.Path] = true

		case *model2.Gen:
			if n.Tok != token.IMPORT {
				return true
			}
			if len(n.Specs) == 1 && mustUnquote(n.Specs[0].(*model2.Import).Path.Value) == "C" {
				hasCgoBlock = true
				return true
			}
			blocks = append(blocks, n)
		case *model2.Import:
			path := mustUnquote(n.Path.Value)
			if n.Name == nil {
				importsFound[path] = ""
			} else {
				importsFound[path] = n.Name.Name
			}
			if path == "C" {
				importsRequired["C"] = true
			}
		}
		return true
	})
	resolved := map[string]string{}
	effectiveAlias := map[string]string{}
	for path, alias := range importsFound {
		if alias == "" {
			continue
		} else if a, ok := r.Alias[path]; ok && a == "" {
			continue
		} else if alias == "_" && packagesInUse[path] {
			continue
		}
		effectiveAlias[path] = alias
	}
	for path, alias := range r.Alias {
		if alias == "" {
			continue
		} else if alias == "_" && packagesInUse[path] {
			continue
		}
		effectiveAlias[path] = alias
	}
	for path, alias := range effectiveAlias {
		if alias == "_" {
			importsRequired[path] = true
		}
	}
	for path := range packagesInUse {
		if _, ok := effectiveAlias[path]; ok {
			continue
		} else if name, err := r.Resolver.ResolvePackage(path); err != nil {
			return fmt.Errorf("could not resolve package %s: %w", path, err)
		} else {
			resolved[path] = name
		}
	}
	importsRequiredOrdered := make([]string, len(importsRequired))
	i := 0
	for path := range importsRequired {
		importsRequiredOrdered[i] = path
		i++
	}
	sort.Slice(importsRequiredOrdered, func(i, j int) bool { return packagePathOrderLess(importsRequiredOrdered[i], importsRequiredOrdered[j]) })
	aliases := map[string]string{}
	r.packageNames = map[string]string{}
	conflict := func(name string) bool {
		for _, n := range r.packageNames {
			if name == n {
				return true
			}
		}
		return false
	}
	findAlias := func(path, preferred string) (name, alias string) {
		aliased := preferred != ""
		if !aliased {
			preferred = resolved[path]
		}
		modifier := 1
		current := preferred
		for conflict(current) {
			current = fmt.Sprintf("%s%d", preferred, modifier)
			modifier++
		}
		if !aliased && current == resolved[path] {
			return current, ""
		}
		return current, current
	}
	for _, path := range importsRequiredOrdered {
		alias := effectiveAlias[path]
		if alias == "." || alias == "_" {
			r.packageNames[path], aliases[path] = "", alias
			continue
		}
		r.packageNames[path], aliases[path] = findAlias(path, alias)
	}
	var added bool
	for _, path := range importsRequiredOrdered {
		if _, ok := importsFound[path]; ok {
			continue
		}
		added = true
		if len(blocks) == 0 {
			gd := &model2.Gen{
				Tok: token.IMPORT,
				Decs: model2.GenDecorations{
					NodeDecs: model2.NodeDecs{Before: model2.EmptyLine, After: model2.EmptyLine},
				},
			}
			if hasCgoBlock {
				r.file.Decls = append([]model2.Decl{r.file.Decls[0], gd}, r.file.Decls[1:]...)
			} else {
				r.file.Decls = append([]model2.Decl{gd}, r.file.Decls...)
			}
			blocks = append(blocks, gd)
		}
		is := &model2.Import{
			Path: &model2.Lit{Kind: token.STRING, Value: fmt.Sprintf("%q", path)},
		}
		if aliases[path] != "" {
			is.Name = &model2.Ident{
				Name: aliases[path],
			}
		}
		blocks[0].Specs = append(blocks[0].Specs, is)
	}
	if added {
		sort.Slice(blocks[0].Specs, func(i, j int) bool {
			return packagePathOrderLess(
				mustUnquote(blocks[0].Specs[i].(*model2.Import).Path.Value),
				mustUnquote(blocks[0].Specs[j].(*model2.Import).Path.Value),
			)
		})
	}
	deleteBlocks := map[model2.Decl]bool{}
	for _, block := range blocks {
		specs := make([]model2.Spec, 0, len(block.Specs))
		for _, spec := range block.Specs {
			spec := spec.(*model2.Import)
			path := mustUnquote(spec.Path.Value)
			if importsRequired[path] {
				if spec.Name == nil && aliases[path] != "" {

					spec.Name = &model2.Ident{Name: aliases[path]}
				} else if spec.Name != nil && aliases[path] == "" {

					spec.Name = nil
				} else if spec.Name != nil && aliases[path] != spec.Name.Name {

					spec.Name.Name = aliases[path]
				}
				specs = append(specs, spec)
			}
		}
		count := len(specs)
		if count != len(block.Specs) {
			block.Specs = specs
			if count == 0 {
				deleteBlocks[block] = true
			} else if count == 1 {
				block.Lparen = false
				block.Rparen = false
			} else {
				block.Lparen = true
				block.Rparen = true
			}
		}
	}
	if added {
		var foundDomainImport bool
		for _, spec := range blocks[0].Specs {
			path := mustUnquote(spec.(*model2.Import).Path.Value)
			if strings.Contains(path, ".") && !foundDomainImport {
				spec.Decorations().Before = model2.EmptyLine
				spec.Decorations().After = model2.NewLine
				foundDomainImport = true
				continue
			}
			spec.Decorations().Before = model2.NewLine
			spec.Decorations().After = model2.NewLine
		}
		if len(blocks[0].Specs) == 1 {
			blocks[0].Lparen = false
			blocks[0].Rparen = false
		} else {
			blocks[0].Lparen = true
			blocks[0].Rparen = true
		}
	}
	if len(deleteBlocks) > 0 {
		decls := make([]model2.Decl, 0, len(r.file.Decls))
		for _, decl := range r.file.Decls {
			if deleteBlocks[decl] {
				continue
			}
			decls = append(decls, decl)
		}
		r.file.Decls = decls
	}
	return nil
}
func (r *FileRestorer) restoreIdent(n *model2.Ident, parentName, parentField, parentFieldType string, allowDuplicate bool) ast.Node {
	if r.Resolver == nil && n.Path != "" {
		panic("This syntax has been decorated with import management enabled, but the restorer does not have import management enabled. Use NewRestorerWithImports to create a restorer with import management. See the Imports section of the readme for more information.")
	}
	var name string
	if r.Resolver != nil && n.Path != "" {
		if model2.Avoid[parentName+"."+parentField] {
			panic(fmt.Sprintf("Path %s set on illegal Ident %s: parentName %s, parentField %s, parentFieldType %s", n.Path, n.Name, parentName, parentField, parentFieldType))
		}
		if n.Path != r.Path {
			name = r.packageNames[n.Path]
		}
		if name == "." {
			name = ""
		}
	}
	if name == "" {
		return nil
	}
	out := &ast.SelectorExpr{}
	r.Ast.Nodes[n] = out
	r.Dst.Nodes[out] = n
	r.Dst.Nodes[out.Sel] = n
	r.Dst.Nodes[out.X] = n
	r.applySpace(n, "Before", n.Decs.Before)
	r.applyDecorations(out, "Start", n.Decs.Start, false)
	out.X = r.restoreNode(model2.NewIdent(name), "Selector", "X", "Expr", allowDuplicate).(ast.Expr)
	r.cursor += token.Pos(len(token.PERIOD.String()))
	r.applyDecorations(out, "X", n.Decs.X, false)
	out.Sel = r.restoreNode(model2.NewIdent(n.Name), "Selector", "Sel", "Ident", allowDuplicate).(*ast.Ident)
	r.applyDecorations(out, "End", n.Decs.End, true)
	r.applySpace(n, "After", n.Decs.After)
	return out
}
func packagePathOrderLess(pi, pj string) bool {
	idot := strings.Contains(pi, ".")
	jdot := strings.Contains(pj, ".")
	if idot != jdot {
		return jdot
	}
	return pi < pj
}
func (r *FileRestorer) fileSize() int {
	end := int(r.cursor)
	for _, cg := range r.comments {
		if int(cg.End()) >= end {
			end = int(cg.End()) + 1
		}
	}
	for _, lineOffset := range r.lines {
		pos := lineOffset + r.base
		if pos >= end {
			end = pos + 1
		}
	}
	return end - r.base
}

func (r *FileRestorer) applyLiteral(text string) {
	isMultiLine := strings.HasPrefix(text, "`") && strings.Contains(text, "\n")
	if !isMultiLine {
		return
	}
	for charIndex, char := range text {
		if char == '\n' {
			lineOffset := int(r.cursor) - r.base + charIndex
			r.lines = append(r.lines, lineOffset)
		}
	}
}

func (r *FileRestorer) hasCommentField(n ast.Node) bool {
	switch n.(type) {
	case *ast.Field, *ast.ValueSpec, *ast.TypeSpec, *ast.ImportSpec:
		return true
	}
	return false
}

func (r *FileRestorer) addCommentField(n ast.Node, slash token.Pos, text string) {
	c := &ast.Comment{Slash: slash, Text: text}
	switch n := n.(type) {
	case *ast.Field:
		if n.Comment == nil {
			n.Comment = &ast.CommentGroup{}
			r.comments = append(r.comments, n.Comment)
		}
		n.Comment.List = append(n.Comment.List, c)
	case *ast.ImportSpec:
		if n.Comment == nil {
			n.Comment = &ast.CommentGroup{}
			r.comments = append(r.comments, n.Comment)
		}
		n.Comment.List = append(n.Comment.List, c)
	case *ast.ValueSpec:
		if n.Comment == nil {
			n.Comment = &ast.CommentGroup{}
			r.comments = append(r.comments, n.Comment)
		}
		n.Comment.List = append(n.Comment.List, c)
	case *ast.TypeSpec:
		if n.Comment == nil {
			n.Comment = &ast.CommentGroup{}
			r.comments = append(r.comments, n.Comment)
		}
		n.Comment.List = append(n.Comment.List, c)
	}
}

func (r *FileRestorer) applyDecorations(node ast.Node, name string, decorations model2.Decs, end bool) {
	firstLine := true
	_, isNodeFile := node.(*ast.File)
	isPackageComment := isNodeFile && name == "Start"

	for _, d := range decorations {

		isNewline := d == "\n"
		isLineComment := strings.HasPrefix(d, "//")
		isInlineComment := strings.HasPrefix(d, "/*")
		isComment := isLineComment || isInlineComment
		isMultiLineComment := isInlineComment && strings.Contains(d, "\n")

		if end && r.cursorAtNewLine == r.cursor {
			r.cursor++
		}

		if isMultiLineComment {
			for charIndex, char := range d {
				if char == '\n' {
					lineOffset := int(r.cursor) - r.base + charIndex
					r.lines = append(r.lines, lineOffset)
				}
			}
		}

		if isComment {
			if firstLine && end && r.hasCommentField(node) {

				r.addCommentField(node, r.cursor, d)
			} else {
				r.comments = append(r.comments, &ast.CommentGroup{List: []*ast.Comment{{Slash: r.cursor, Text: d}}})
			}
			r.cursor += token.Pos(len(d))
		}

		if isLineComment || isNewline {
			lineOffset := int(r.cursor) - r.base
			r.lines = append(r.lines, lineOffset)
			r.cursor++

			r.cursorAtNewLine = r.cursor
		}

		if isNewline || isLineComment {
			firstLine = false
		}
	}
	if isPackageComment {

		r.cursor++
	}
}

func (r *FileRestorer) applySpace(node model2.Node, position string, space model2.SpaceType) {
	switch node.(type) {
	case *model2.BadDecl, *model2.BadExpr, *model2.BadStmt:
		if position == "After" {

			space = model2.EmptyLine
		}
	}
	var newlines int
	switch space {
	case model2.NewLine:
		newlines = 1
	case model2.EmptyLine:
		newlines = 2
	}
	if r.cursor == r.cursorAtNewLine {
		newlines--
	}
	for i := 0; i < newlines; i++ {

		r.cursor++

		lineOffset := int(r.cursor) - r.base
		r.lines = append(r.lines, lineOffset)
		r.cursor++
		r.cursorAtNewLine = r.cursor
	}
}

func (r *FileRestorer) restoreObject(o *model2.Object) *ast.Object {
	if !r.Extras {
		return nil
	} else if o == nil {
		return nil
	} else if ro, ok := r.Ast.Objects[o]; ok {
		return ro
	}
	out := &ast.Object{}
	r.Ast.Objects[o] = out
	r.Dst.Objects[out] = o
	out.Kind = ast.ObjKind(o.Kind)
	out.Name = o.Name

	switch decl := o.Decl.(type) {
	case *model2.Scope:
		out.Decl = r.restoreScope(decl)
	case model2.Node:
		r.nodeDecl[out] = decl
	case nil:
	default:
		panic(fmt.Sprintf("o.Decl is %T", o.Decl))
	}
	switch data := o.Data.(type) {
	case int:
		out.Data = data
	case *model2.Scope:
		out.Data = r.restoreScope(data)
	case model2.Node:
		r.nodeData[out] = data
	case nil:
	default:
		panic(fmt.Sprintf("o.Data is %T", o.Data))
	}
	return out
}
func (r *FileRestorer) restoreScope(s *model2.Scope) *ast.Scope {
	if !r.Extras {
		return nil
	}
	if s == nil {
		return nil
	}
	if rs, ok := r.Ast.Scopes[s]; ok {
		return rs
	}
	out := &ast.Scope{}
	r.Ast.Scopes[s] = out
	r.Dst.Scopes[out] = s
	out.Outer = r.restoreScope(s.Outer)
	out.Objects = map[string]*ast.Object{}
	for k, v := range s.Objects {
		out.Objects[k] = r.restoreObject(v)
	}
	return out
}
func mustUnquote(s string) string {
	out, err := strconv.Unquote(s)
	if err != nil {
		panic(err)
	}
	return out
}
