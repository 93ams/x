package model

import (
	"grpc/package/wrapper/model/meta"
	"strings"
)

type (
	HasInlineCommentSetter interface{ SetComment(*Comment) }
	Comment                struct {
		Raw  string
		Meta meta.Meta
	}
)

func (c *Comment) IsCStyle() bool   { return strings.HasPrefix(c.Raw, cStyleCommentPrefix) }
func (c *Comment) Accept(v Visitor) { v.VisitComment(c) }
func (c *Comment) Lines() []string {
	raw := c.Raw
	if c.IsCStyle() {
		raw = strings.TrimPrefix(raw, cStyleCommentPrefix)
		raw = strings.TrimSuffix(raw, cStyleCommentSuffix)
	} else {
		raw = strings.TrimPrefix(raw, cPlusStyleCommentPrefix)
	}
	return strings.Split(raw, "\n")
}

func (r *RPC) SetComment(comment *Comment)            { r.Comment = comment }
func (e *Enum) SetComment(comment *Comment)           { e.Comment = comment }
func (o *OneOf) SetComment(comment *Comment)          { o.Comment = comment }
func (f *Field) SetComment(comment *Comment)          { f.Comment = comment }
func (i *Import) SetComment(comment *Comment)         { i.Comment = comment }
func (m *Extend) SetComment(comment *Comment)         { m.Comment = comment }
func (s *Syntax) SetComment(comment *Comment)         { s.Comment = comment }
func (o *Option) SetComment(comment *Comment)         { o.Comment = comment }
func (m *Message) SetComment(comment *Comment)        { m.Comment = comment }
func (p *Package) SetComment(comment *Comment)        { p.Comment = comment }
func (s *Service) SetComment(comment *Comment)        { s.Comment = comment }
func (r *Reserved) SetComment(comment *Comment)       { r.Comment = comment }
func (m *MapField) SetComment(comment *Comment)       { m.Comment = comment }
func (f *EnumField) SetComment(comment *Comment)      { f.Comment = comment }
func (e *Extensions) SetComment(comment *Comment)     { e.Comment = comment }
func (f *OneOfField) SetComment(comment *Comment)     { f.Comment = comment }
func (f *GroupField) SetComment(comment *Comment)     { f.Comment = comment }
func (e *EmptyStatement) SetComment(comment *Comment) { e.Comment = comment }
