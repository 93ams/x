package model

func (e *EmptyStatement) Accept(v Visitor) {
	if !v.VisitEmptyStatement(e) {
		return
	}
	if e.Comment != nil {
		e.Comment.Accept(v)
	}
}
func (f *EnumField) Accept(v Visitor) {
	if !v.VisitEnumField(f) {
		return
	}
	for _, comment := range f.Comments {
		comment.Accept(v)
	}
	if f.Comment != nil {
		f.Comment.Accept(v)
	}
}
func (e *Enum) Accept(v Visitor) {
	if !v.VisitEnum(e) {
		return
	}
	for _, body := range e.Body {
		body.Accept(v)
	}
	for _, comment := range e.Comments {
		comment.Accept(v)
	}
	if e.Comment != nil {
		e.Comment.Accept(v)
	}
	if e.CommentBLC != nil {
		e.CommentBLC.Accept(v)
	}
}
func (m *Extend) Accept(v Visitor) {
	if !v.VisitExtend(m) {
		return
	}
	for _, body := range m.Body {
		body.Accept(v)
	}
	for _, comment := range m.Comments {
		comment.Accept(v)
	}
	if m.Comment != nil {
		m.Comment.Accept(v)
	}
	if m.CommentBLC != nil {
		m.CommentBLC.Accept(v)
	}
}
func (f *Field) Accept(v Visitor) {
	if !v.VisitField(f) {
		return
	}
	for _, comment := range f.Comments {
		comment.Accept(v)
	}
	if f.Comment != nil {
		f.Comment.Accept(v)
	}
}
func (f *GroupField) Accept(v Visitor) {
	if !v.VisitGroupField(f) {
		return
	}
	for _, body := range f.Body {
		body.Accept(v)
	}
	for _, comment := range f.Comments {
		comment.Accept(v)
	}
	if f.Comment != nil {
		f.Comment.Accept(v)
	}
	if f.CommentBLC != nil {
		f.CommentBLC.Accept(v)
	}
}
func (e *Extensions) Accept(v Visitor) {
	if !v.VisitExtensions(e) {
		return
	}
	for _, comment := range e.Comments {
		comment.Accept(v)
	}
	if e.Comment != nil {
		e.Comment.Accept(v)
	}
}
func (i *Import) Accept(v Visitor) {
	if !v.VisitImport(i) {
		return
	}
	for _, comment := range i.Comments {
		comment.Accept(v)
	}
	if i.Comment != nil {
		i.Comment.Accept(v)
	}
}
func (m *Message) Accept(v Visitor) {
	if !v.VisitMessage(m) {
		return
	}
	for _, body := range m.MessageBody {
		body.Accept(v)
	}
	for _, comment := range m.Comments {
		comment.Accept(v)
	}
	if m.Comment != nil {
		m.Comment.Accept(v)
	}
	if m.CommentBLC != nil {
		m.CommentBLC.Accept(v)
	}
}
func (m *MapField) Accept(v Visitor) {
	if !v.VisitMapField(m) {
		return
	}

	for _, comment := range m.Comments {
		comment.Accept(v)
	}
	if m.Comment != nil {
		m.Comment.Accept(v)
	}
}
func (s *Syntax) Accept(v Visitor) {
	if !v.VisitSyntax(s) {
		return
	}
	for _, comment := range s.Comments {
		comment.Accept(v)
	}
	if s.Comment != nil {
		s.Comment.Accept(v)
	}
}
func (r *RPC) Accept(v Visitor) {
	if !v.VisitRPC(r) {
		return
	}

	for _, comment := range r.Comments {
		comment.Accept(v)
	}
	if r.Comment != nil {
		r.Comment.Accept(v)
	}
}
func (s *Service) Accept(v Visitor) {
	if !v.VisitService(s) {
		return
	}
	for _, body := range s.Body {
		body.Accept(v)
	}
	for _, comment := range s.Comments {
		comment.Accept(v)
	}
	if s.Comment != nil {
		s.Comment.Accept(v)
	}
	if s.CommentBLC != nil {
		s.CommentBLC.Accept(v)
	}
}
func (o *Option) Accept(v Visitor) {
	if !v.VisitOption(o) {
		return
	}
	for _, comment := range o.Comments {
		comment.Accept(v)
	}
	if o.Comment != nil {
		o.Comment.Accept(v)
	}
}
func (f *OneOfField) Accept(v Visitor) {
	if !v.VisitOneofField(f) {
		return
	}
	for _, comment := range f.Comments {
		comment.Accept(v)
	}
	if f.Comment != nil {
		f.Comment.Accept(v)
	}
}
func (o *OneOf) Accept(v Visitor) {
	if !v.VisitOneof(o) {
		return
	}
	for _, field := range o.OneofFields {
		field.Accept(v)
	}
	for _, option := range o.Options {
		option.Accept(v)
	}
	for _, comment := range o.Comments {
		comment.Accept(v)
	}
	if o.Comment != nil {
		o.Comment.Accept(v)
	}
}
func (p *Proto) Accept(v Visitor) {
	if p.Syntax != nil {
		p.Syntax.Accept(v)
	}
	for _, body := range p.ProtoBody {
		body.Accept(v)
	}
}
func (r *Reserved) Accept(v Visitor) {
	if !v.VisitReserved(r) {
		return
	}
	for _, comment := range r.Comments {
		comment.Accept(v)
	}
	if r.Comment != nil {
		r.Comment.Accept(v)
	}
}
func (p *Package) Accept(v Visitor) {
	if !v.VisitPackage(p) {
		return
	}

	for _, comment := range p.Comments {
		comment.Accept(v)
	}
	if p.Comment != nil {
		p.Comment.Accept(v)
	}
}
