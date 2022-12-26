package model

import "github.com/samber/lo"

func (n *Block) SetDecs(decs BlockDecs) *Block                      { n.Decs = decs; return n }
func (n *ExprStmt) SetDecs(decs ExprStmtDecs) *ExprStmt             { n.Decs = decs; return n }
func (n *Return) SetDecs(decs ReturnDecs) *Return                   { n.Decs = decs; return n }
func (n *ReturnDecs) SetBefore(d SpaceType) *ReturnDecs             { n.Before = d; return n }
func (n *ReturnDecs) SetAfter(d SpaceType) *ReturnDecs              { n.After = d; return n }
func (n *ReturnDecs) SetStart(d Decs) *ReturnDecs                   { n.Start = d; return n }
func (n *ReturnDecs) SetEnd(d Decs) *ReturnDecs                     { n.End = d; return n }
func (n *Import) SetDecs(decs ImportDecs) *Import                   { n.Decs = decs; return n }
func (n *Import) SetName(path *Ident) *Import                       { n.Name = path; return n }
func (n *Import) SetPath(path *Lit) *Import                         { n.Path = path; return n }
func (n *Type) SetDecs(decs TypeDecs) *Type                         { n.Decs = decs; return n }
func (n *Type) SetParams(params *FieldList) *Type                   { n.Params = params; return n }
func (n *Type) SetName(params *Ident) *Type                         { n.Name = params; return n }
func (n *Type) SetAssign(assign bool) *Type                         { n.Assign = assign; return n }
func (n *Value) SetDecs(decs ValueDecs) *Value                      { n.Decs = decs; return n }
func (n *Value) SetValues(values ...Expr) *Value                    { n.Values = values; return n }
func (n *Call) SetDecs(decs CallDecs) *Call                         { n.Decs = decs; return n }
func (n *Call) SetEllipsis(ellipsis bool) *Call                     { n.Ellipsis = ellipsis; return n }
func (n *CompositeLit) SetElts(e ...Expr) *CompositeLit             { n.Elts = append(n.Elts, e...); return n }
func (n *CompositeLit) SetIncomplete(i bool) *CompositeLit          { n.Incomplete = i; return n }
func (n *CompositeLit) SetDecs(decs CompositeLitDecs) *CompositeLit { n.Decs = decs; return n }
func (n *CompositeLitDecs) SetBefore(d SpaceType) *CompositeLitDecs { n.Before = d; return n }
func (n *CompositeLitDecs) SetAfter(d SpaceType) *CompositeLitDecs  { n.After = d; return n }
func (n *CompositeLitDecs) SetStart(d Decs) *CompositeLitDecs       { n.Start = d; return n }
func (n *CompositeLitDecs) SetEnd(d Decs) *CompositeLitDecs         { n.End = d; return n }
func (n *FuncType) SetDecs(decs FuncTypeDecs) *FuncType             { n.Decs = decs; return n }
func (n *FuncType) SetTypeParams(fields *FieldList) *FuncType       { n.TypeParams = fields; return n }
func (n *FuncType) SetParams(fields *FieldList) *FuncType           { n.Params = fields; return n }
func (n *FuncType) SetResults(fields *FieldList) *FuncType          { n.Results = fields; return n }
func (n *FuncType) SetFunc(fn bool) *FuncType                       { n.Func = fn; return n }
func (n *Ident) SetDecs(decs IdentDecs) *Ident                      { n.Decs = decs; return n }
func (n *Ident) SetOnj(onj *Object) *Ident                          { n.Obj = onj; return n }
func (n *Ident) SetPath(path string) *Ident                         { n.Path = path; return n }
func (n *Interface) SetDecs(decs InterfaceDecs) *Interface          { n.Decs = decs; return n }
func (n *Interface) SetIncomplete(incomplete bool) *Interface       { n.Incomplete = incomplete; return n }
func (n *KeyValue) SetDecs(decs KeyValueDecs) *KeyValue             { n.Decs = decs; return n }
func (n *KeyValueDecs) SetBefore(d SpaceType) *KeyValueDecs         { n.Before = d; return n }
func (n *KeyValueDecs) SetAfter(d SpaceType) *KeyValueDecs          { n.After = d; return n }
func (n *KeyValueDecs) SetStart(d Decs) *KeyValueDecs               { n.Start = d; return n }
func (n *KeyValueDecs) SetEnd(d Decs) *KeyValueDecs                 { n.End = d; return n }
func (n *Lit) SetDecs(decs LitDecs) *Lit                            { n.Decs = decs; return n }
func (n *Selector) SetDecs(decs SelectorDecs) *Selector             { n.Decs = decs; return n }
func (n *Struct) SetDecs(decs StructDecs) *Struct                   { n.Decs = decs; return n }
func (n *Struct) SetIncomplete(incomplete bool) *Struct             { n.Incomplete = incomplete; return n }
func (n *Gen) SetDecs(decs GenDecs) *Gen                            { n.Decs = decs; return n }
func (n *Func) SetDecs(decs FuncDecs) *Func                         { n.Decs = decs; return n }
func (n *Func) SetRecv(fields *FieldList) *Func                     { n.Recv = fields; return n }
func (n *Func) SetType(t *FuncType) *Func                           { n.Type = t; return n }
func (n *Func) SetBody(nody *Block) *Func                           { n.Body = nody; return n }
func (f *FuncDecs) SetTypeParams(d Decs) *FuncDecs                  { f.TypeParams = d; return f }
func (f *FuncDecs) SetResults(d Decs) *FuncDecs                     { f.Results = d; return f }
func (f *FuncDecs) SetParams(d Decs) *FuncDecs                      { f.Params = d; return f }
func (f *FuncDecs) SetStart(d Decs) *FuncDecs                       { f.Start = d; return f }
func (f *FuncDecs) SetName(d Decs) *FuncDecs                        { f.Name = d; return f }
func (f *FuncDecs) SetFunc(d Decs) *FuncDecs                        { f.Func = d; return f }
func (f *FuncDecs) SetRecv(d Decs) *FuncDecs                        { f.Recv = d; return f }
func (f *FuncDecs) SetEnd(d Decs) *FuncDecs                         { f.End = d; return f }
func (n *Field) SetDecs(decs FieldDecs) *Field                      { n.Decs = decs; return n }
func (n *Field) SetTag(tag *Lit) *Field                             { n.Tag = tag; return n }
func (n *FieldList) SetDecs(decs FieldListDecs) *FieldList          { n.Decs = decs; return n }
func (n *File) SetDecls(decls ...Decl) *File                        { n.Decls = append(n.Decls, decls...); return n }
func (n *File) SetDecs(decs FileDecs) *File                         { n.Decs = decs; return n }
func (n *File) SetScope(name *Scope) *File                          { n.Scope = name; return n }
func (n *Package) SetScope(scope *Scope) *Package                   { n.Scope = scope; return n }
func (n *Package) SetDecs(decs PackageDecs) *Package                { n.Decs = decs; return n }
func (n *Package) AddFiles(f map[string]*File) *Package             { n.Files = lo.Assign(n.Files, f); return n }
func (n *Package) AddImports(i map[string]*Object) *Package {
	n.Imports = lo.Assign(n.Imports, i)
	return n
}
