package model

import (
	"grpc/package/wrapper/model/meta"
)

type (
	EmptyStatement struct{ Comment *Comment }
	ProtoMeta      struct{ Filename string }
	Proto          struct {
		Syntax    *Syntax
		ProtoBody []Visitee
		Meta      *ProtoMeta
	}
	Package struct {
		Name     string
		Comments []*Comment
		Comment  *Comment
		Meta     meta.Meta
	}
	EnumValueOption struct {
		Name     string
		Constant string
	}
	EnumField struct {
		Ident    string
		Number   string
		Options  []*EnumValueOption
		Comments []*Comment
		Comment  *Comment
		Meta     meta.Meta
	}
	Enum struct {
		Name       string
		Body       []Visitee
		Comments   []*Comment
		Comment    *Comment
		CommentBLC *Comment
		Meta       meta.Meta
	}
	Extend struct {
		Type       string
		Body       []Visitee
		Comments   []*Comment
		Comment    *Comment
		CommentBLC *Comment
		Meta       meta.Meta
	}
	FieldOption struct {
		Name     string
		Constant string
	}
	Field struct {
		IsRepeated bool
		IsRequired bool
		IsOptional bool
		Type       string
		Name       string
		Number     string
		Options    []*FieldOption
		Comments   []*Comment
		Comment    *Comment
		Meta       meta.Meta
	}
	GroupField struct {
		IsRepeated bool
		IsRequired bool
		IsOptional bool
		GroupName  string
		Body       []Visitee
		Number     string
		Comments   []*Comment
		Comment    *Comment
		CommentBLC *Comment
		Meta       meta.Meta
	}
	Extensions struct {
		Ranges   []*Range
		Comments []*Comment
		Comment  *Comment
		Meta     meta.Meta
	}
	ImportModifier uint
	Import         struct {
		Modifier ImportModifier
		Location string
		Comments []*Comment
		Comment  *Comment
		Meta     meta.Meta
	}
	Message struct {
		MessageName string
		MessageBody []Visitee
		Comments    []*Comment
		Comment     *Comment
		CommentBLC  *Comment
		Meta        meta.Meta
	}
	MapField struct {
		KeyType      string
		Type         string
		MapName      string
		FieldNumber  string
		FieldOptions []*FieldOption
		Comments     []*Comment
		Comment      *Comment
		Meta         meta.Meta
	}
	Syntax struct {
		ProtobufVersion string
		VersionQuote    string
		Comments        []*Comment
		Comment         *Comment
		Meta            meta.Meta
	}
	Request struct {
		IsStream    bool
		MessageType string
		Meta        meta.Meta
	}
	Response struct {
		IsStream bool
		Type     string
		Meta     meta.Meta
	}
	RPC struct {
		Name       string
		Request    *Request
		Response   *Response
		Options    []*Option
		Comments   []*Comment
		Comment    *Comment
		CommentBLC *Comment
		Meta       meta.Meta
	}
	Service struct {
		Name       string
		Body       []Visitee
		Comments   []*Comment
		Comment    *Comment
		CommentBLC *Comment
		Meta       meta.Meta
	}
	Range struct {
		Begin string
		End   string
	}
	Reserved struct {
		Ranges     []*Range
		FieldNames []string
		Comments   []*Comment
		Comment    *Comment
		Meta       meta.Meta
	}
	Option struct {
		Name     string
		Constant string
		Comments []*Comment
		Comment  *Comment
		Meta     meta.Meta
	}
	OneOfField struct {
		Type     string
		Name     string
		Number   string
		Options  []*FieldOption
		Comments []*Comment
		Comment  *Comment
		Meta     meta.Meta
	}
	OneOf struct {
		OneofFields []*OneOfField
		OneofName   string
		Options     []*Option
		Comments    []*Comment
		Comment     *Comment
		CommentBLC  *Comment
		Meta        meta.Meta
	}
)

const (
	cStyleCommentPrefix     = "/*"
	cStyleCommentSuffix     = "*/"
	cPlusStyleCommentPrefix = "//"
)

const (
	ImportModifierNone ImportModifier = iota
	ImportModifierPublic
	ImportModifierWeak
)
