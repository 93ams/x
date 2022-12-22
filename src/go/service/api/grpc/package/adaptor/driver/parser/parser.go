package parser

import (
	. "github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/adaptor/driver/lexer"
	scanner2 "github.com/tilau2328/x/src/go/service/api/grpc/package/adaptor/driver/lexer/scanner"
	model2 "github.com/tilau2328/x/src/go/service/api/grpc/package/adaptor/driver/model"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/adaptor/driver/model/meta"
	"strings"
	"unicode/utf8"
)

type Parser struct {
	lex                   *lexer.Lexer
	permissive            bool
	bodyIncludingComments bool
}

func Permissive(permissive bool) Opt[*Parser] {
	return func(p *Parser) { p.permissive = permissive }
}
func BodyIncludingComments(bodyIncludingComments bool) Opt[*Parser] {
	return func(p *Parser) { p.bodyIncludingComments = bodyIncludingComments }
}
func NewParser(lex *lexer.Lexer, opts ...Opt[*Parser]) *Parser {
	p := &Parser{lex: lex}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
func (p *Parser) IsEOF() bool {
	p.lex.Next()
	defer p.lex.UnNext()
	return p.lex.IsEOF()
}
func (p *Parser) ParseComments() []*model2.Comment {
	var comments []*model2.Comment
	for {
		comment, err := p.parseComment()
		if err != nil {
			return comments
		}
		comments = append(comments, comment)
	}
}
func (p *Parser) parseComment() (*model2.Comment, error) {
	p.lex.NextComment()
	if p.lex.Token == scanner2.TCOMMENT {
		return &model2.Comment{
			Raw: p.lex.Text,
			Meta: meta.Meta{
				Pos:     p.lex.Pos.Position,
				LastPos: p.lex.Pos.AdvancedBulk(p.lex.Text).Position,
			},
		}, nil
	}
	defer p.lex.UnNext()
	return nil, p.unexpected("comment")
}

func (p *Parser) ParseSyntax() (*model2.Syntax, error) {
	if p.lex.NextKeyword(); p.lex.Token != scanner2.TSYNTAX {
		return nil, p.unexpected("syntax")
	}
	startPos := p.lex.Pos
	if p.lex.Next(); p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	}
	if p.lex.Next(); p.lex.Token != scanner2.TQUOTE {
		return nil, p.unexpected("quote")
	}
	lq := p.lex.Text
	if p.lex.Next(); p.lex.Text != "proto3" && p.lex.Text != "proto2" {
		return nil, p.unexpected("proto3 or proto2")
	}
	version := p.lex.Text
	if p.lex.Next(); p.lex.Token != scanner2.TQUOTE {
		return nil, p.unexpected("quote")
	}
	tq := p.lex.Text
	if p.lex.Next(); p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}
	return &model2.Syntax{
		ProtobufVersion: version,
		VersionQuote:    lq + version + tq,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
func (p *Parser) ParseEnum() (*model2.Enum, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TENUM {
		return nil, p.unexpected("enum")
	}
	startPos := p.lex.Pos
	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("enumName")
	}
	enumName := p.lex.Text
	enumBody, inlineLeftCurly, lastPos, err := p.parseEnumBody()
	if err != nil {
		return nil, err
	}

	return &model2.Enum{
		Name:       enumName,
		Body:       enumBody,
		CommentBLC: inlineLeftCurly,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: lastPos.Position,
		},
	}, nil
}
func (p *Parser) parseEnumBody() (
	[]model2.Visitee,
	*model2.Comment,
	scanner2.Position,
	error,
) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTCURLY {
		return nil, nil, scanner2.Position{}, p.unexpected("{")
	}
	inlineLeftCurly := p.parseInlineComment()
	var stmts []model2.Visitee
	for {
		comments := p.ParseComments()
		p.lex.NextKeyword()
		token := p.lex.Token
		p.lex.UnNext()
		var stmt interface {
			model2.HasInlineCommentSetter
			model2.Visitee
		}
		switch token {
		case scanner2.TRIGHTCURLY:
			if p.bodyIncludingComments {
				for _, comment := range comments {
					stmts = append(stmts, model2.Visitee(comment))
				}
			}
			p.lex.Next()
			lastPos := p.lex.Pos
			if p.permissive {
				// accept a block followed by semicolon. See https://github.com/yoheimuta/go-protoparser/v4/issues/30.
				p.lex.ConsumeToken(scanner2.TSEMICOLON)
				if p.lex.Token == scanner2.TSEMICOLON {
					lastPos = p.lex.Pos
				}
			}
			return stmts, inlineLeftCurly, lastPos, nil
		case scanner2.TOPTION:
			option, err := p.ParseOption()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			option.Comments = comments
			stmt = option
		case scanner2.TRESERVED:
			reserved, err := p.ParseReserved()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			reserved.Comments = comments
			stmt = reserved
		default:
			enumField, enumFieldErr := p.parseEnumField()
			if enumFieldErr == nil {
				enumField.Comments = comments
				stmt = enumField
				break
			}
			p.lex.UnNext()
			emptyErr := p.lex.ReadEmptyStatement()
			if emptyErr == nil {
				stmt = &model2.EmptyStatement{}
				break
			}
			return nil, nil, scanner2.Position{}, &parseEnumBodyStatementErr{
				parseEnumFieldErr:      enumFieldErr,
				parseEmptyStatementErr: emptyErr,
			}
		}
		p.MaybeScanInlineComment(stmt)
		stmts = append(stmts, stmt)
	}
}
func (p *Parser) parseEnumField() (*model2.EnumField, error) {
	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("ident")
	}
	startPos := p.lex.Pos
	ident := p.lex.Text
	p.lex.Next()
	if p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	}
	var intLit string
	p.lex.ConsumeToken(scanner2.TMINUS)
	if p.lex.Token == scanner2.TMINUS {
		intLit = "-"
	}
	p.lex.NextNumberLit()
	if p.lex.Token != scanner2.TINTLIT {
		return nil, p.unexpected("intLit")
	}
	intLit += p.lex.Text
	enumValueOptions, err := p.parseEnumValueOptions()
	if err != nil {
		return nil, err
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}
	return &model2.EnumField{
		Ident:   ident,
		Number:  intLit,
		Options: enumValueOptions,
		Meta:    meta.Meta{Pos: startPos.Position},
	}, nil
}
func (p *Parser) parseEnumValueOptions() ([]*model2.EnumValueOption, error) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTSQUARE {
		p.lex.UnNext()
		return nil, nil
	}
	opt, err := p.parseEnumValueOption()
	if err != nil {
		return nil, p.unexpected("enumValueOption")
	}
	var opts []*model2.EnumValueOption
	opts = append(opts, opt)
	for {
		p.lex.Next()
		if p.lex.Token != scanner2.TCOMMA {
			p.lex.UnNext()
			break
		}
		opt, err = p.parseEnumValueOption()
		if err != nil {
			return nil, p.unexpected("enumValueOption")
		}
		opts = append(opts, opt)
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TRIGHTSQUARE {
		return nil, p.unexpected("]")
	}
	return opts, nil
}

func (p *Parser) parseEnumValueOption() (*model2.EnumValueOption, error) {
	optionName, err := p.parseOptionName()
	if err != nil {
		return nil, err
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	}
	constant, err := p.parseOptionConstant()
	if err != nil {
		return nil, err
	}
	return &model2.EnumValueOption{
		Name:     optionName,
		Constant: constant,
	}, nil
}
func (p *Parser) ParseExtensions() (*model2.Extensions, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TEXTENSIONS {
		return nil, p.unexpected("extensions")
	}
	startPos := p.lex.Pos
	ranges, err := p.parseRanges()
	if err != nil {
		return nil, err
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}
	return &model2.Extensions{
		Ranges: ranges,
		Meta:   meta.Meta{Pos: startPos.Position},
	}, nil
}

func (p *Parser) ParseExtend() (*model2.Extend, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TEXTEND {
		return nil, p.unexpected("extend")
	}
	startPos := p.lex.Pos
	messageType, _, err := p.lex.ReadMessageType()
	if err != nil {
		return nil, err
	}
	extendBody, inlineLeftCurly, lastPos, err := p.parseExtendBody()
	if err != nil {
		return nil, err
	}
	return &model2.Extend{
		Type:       messageType,
		Body:       extendBody,
		CommentBLC: inlineLeftCurly,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: lastPos.Position,
		},
	}, nil
}

func (p *Parser) parseExtendBody() ([]model2.Visitee, *model2.Comment, scanner2.Position, error) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTCURLY {
		return nil, nil, scanner2.Position{}, p.unexpected("{")
	}
	inlineLeftCurly := p.parseInlineComment()
	p.lex.Next()
	if p.lex.Token == scanner2.TRIGHTCURLY {
		lastPos := p.lex.Pos
		if p.permissive {
			p.lex.ConsumeToken(scanner2.TSEMICOLON)
			if p.lex.Token == scanner2.TSEMICOLON {
				lastPos = p.lex.Pos
			}
		}
		return nil, nil, lastPos, nil
	}
	p.lex.UnNext()
	var stmts []model2.Visitee
	for {
		comments := p.ParseComments()
		p.lex.NextKeyword()
		token := p.lex.Token
		p.lex.UnNext()
		var stmt interface {
			model2.HasInlineCommentSetter
			model2.Visitee
		}
		switch token {
		case scanner2.TRIGHTCURLY:
			if p.bodyIncludingComments {
				for _, comment := range comments {
					stmts = append(stmts, model2.Visitee(comment))
				}
			}
			p.lex.Next()
			lastPos := p.lex.Pos
			if p.permissive {
				p.lex.ConsumeToken(scanner2.TSEMICOLON)
				if p.lex.Token == scanner2.TSEMICOLON {
					lastPos = p.lex.Pos
				}
			}
			return stmts, inlineLeftCurly, lastPos, nil
		default:
			field, fieldErr := p.ParseField()
			if fieldErr == nil {
				field.Comments = comments
				stmt = field
				break
			}
			p.lex.UnNext()
			emptyErr := p.lex.ReadEmptyStatement()
			if emptyErr == nil {
				stmt = &model2.EmptyStatement{}
				break
			}

			return nil, nil, scanner2.Position{}, &parseExtendBodyStatementErr{
				parseFieldErr:          fieldErr,
				parseEmptyStatementErr: emptyErr,
			}
		}

		p.MaybeScanInlineComment(stmt)
		stmts = append(stmts, stmt)
	}
}
func (p *Parser) ParseField() (ret *model2.Field, err error) {
	p.lex.NextKeyword()
	ret = &model2.Field{Meta: meta.Meta{Pos: p.lex.Pos.Position}}
	switch p.lex.Token {
	case scanner2.TREPEATED:
		ret.IsRepeated = true
	case scanner2.TREQUIRED:
		ret.IsRequired = true
	case scanner2.TOPTIONAL:
		ret.IsOptional = true
	default:
		p.lex.UnNext()
	}
	if ret.Type, _, err = p.parseType(); err != nil {
		return nil, p.unexpected("type")
	}
	p.lex.Next()
	ret.Name = p.lex.Text
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("fieldName")
	} else if p.lex.Next(); p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	} else if ret.Number, err = p.parseFieldNumber(); err != nil {
		return nil, p.unexpected("fieldNumber")
	} else if ret.Options, err = p.parseFieldOptionsOption(); err != nil {
		return nil, err
	} else if p.lex.Next(); p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}
	return
}
func (p *Parser) parseFieldOptionsOption() ([]*model2.FieldOption, error) {
	if p.lex.Next(); p.lex.Token == scanner2.TLEFTSQUARE {
		if fieldOptions, err := p.parseFieldOptions(); err != nil {
			return nil, err
		} else if p.lex.Next(); p.lex.Token != scanner2.TRIGHTSQUARE {
			return nil, p.unexpected("]")
		} else {
			return fieldOptions, nil
		}
	}
	p.lex.UnNext()
	return nil, nil
}

func (p *Parser) parseFieldOptions() (opts []*model2.FieldOption, err error) {
	opt, err := p.parseFieldOption()
	if err != nil {
		return nil, err
	}
	opts = append(opts, opt)
	for {
		if p.lex.Next(); p.lex.Token != scanner2.TCOMMA {
			p.lex.UnNext()
			break
		} else if opt, err = p.parseFieldOption(); err != nil {
			return nil, p.unexpected("fieldOption")
		}
		opts = append(opts, opt)
	}
	return opts, nil
}

func (p *Parser) parseFieldOption() (*model2.FieldOption, error) {
	if optionName, err := p.parseOptionName(); err != nil {
		return nil, err
	} else if p.lex.Next(); p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	} else if constant, err := p.parseOptionConstant(); err != nil {
		return nil, err
	} else {
		return &model2.FieldOption{Name: optionName, Constant: constant}, nil
	}
}

var typeConstants = map[string]struct{}{
	"double":   {},
	"float":    {},
	"int32":    {},
	"int64":    {},
	"uint32":   {},
	"uint64":   {},
	"sint32":   {},
	"sint64":   {},
	"fixed32":  {},
	"fixed64":  {},
	"sfixed32": {},
	"sfixed64": {},
	"bool":     {},
	"string":   {},
	"bytes":    {},
}

func (p *Parser) parseType() (string, scanner2.Position, error) {
	p.lex.Next()
	if _, ok := typeConstants[p.lex.Text]; ok {
		return p.lex.Text, p.lex.Pos, nil
	}
	p.lex.UnNext()
	messageOrEnumType, startPos, err := p.lex.ReadMessageType()
	if err != nil {
		return "", scanner2.Position{}, err
	}
	return messageOrEnumType, startPos, nil
}

func (p *Parser) parseFieldNumber() (string, error) {
	if p.lex.NextNumberLit(); p.lex.Token != scanner2.TINTLIT {
		return "", p.unexpected("intLit")
	}
	return p.lex.Text, nil
}
func (p *Parser) ParseGroupField() (*model2.GroupField, error) {
	var isRepeated bool
	var isRequired bool
	var isOptional bool
	p.lex.NextKeyword()
	startPos := p.lex.Pos
	if p.lex.Token == scanner2.TREPEATED {
		isRepeated = true
	} else if p.lex.Token == scanner2.TREQUIRED {
		isRequired = true
	} else if p.lex.Token == scanner2.TOPTIONAL {
		isOptional = true
	} else {
		p.lex.UnNext()
	}

	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TGROUP {
		return nil, p.unexpected("group")
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("groupName")
	}
	if !isCapitalized(p.lex.Text) {
		return nil, p.unexpectedf("groupName %q must begin with capital letter.", p.lex.Text)
	}
	groupName := p.lex.Text

	p.lex.Next()
	if p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	}

	fieldNumber, err := p.parseFieldNumber()
	if err != nil {
		return nil, p.unexpected("fieldNumber")
	}

	messageBody, inlineLeftCurly, lastPos, err := p.parseMessageBody()
	if err != nil {
		return nil, err
	}

	return &model2.GroupField{
		IsRepeated: isRepeated,
		IsRequired: isRequired,
		IsOptional: isOptional,
		GroupName:  groupName,
		Number:     fieldNumber,
		Body:       messageBody,

		CommentBLC: inlineLeftCurly,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: lastPos.Position,
		},
	}, nil
}
func (p *Parser) peekIsGroup() bool {
	p.lex.NextKeyword()
	switch p.lex.Token {
	case scanner2.TREPEATED,
		scanner2.TREQUIRED,
		scanner2.TOPTIONAL:
		defer p.lex.UnNextTo(p.lex.RawText)
	default:
		p.lex.UnNext()
	}

	p.lex.NextKeyword()
	defer p.lex.UnNextTo(p.lex.RawText)
	if p.lex.Token != scanner2.TGROUP {
		return false
	}

	p.lex.Next()
	defer p.lex.UnNextTo(p.lex.RawText)
	if p.lex.Token != scanner2.TIDENT {
		return false
	}
	if !isCapitalized(p.lex.Text) {
		return false
	}

	p.lex.Next()
	defer p.lex.UnNextTo(p.lex.RawText)
	if p.lex.Token != scanner2.TEQUALS {
		return false
	}

	_, err := p.parseFieldNumber()
	defer p.lex.UnNextTo(p.lex.RawText)
	if err != nil {
		return false
	}

	p.lex.Next()
	defer p.lex.UnNextTo(p.lex.RawText)
	if p.lex.Token != scanner2.TLEFTCURLY {
		return false
	}
	return true
}

func isCapitalized(s string) bool {
	if s == "" {
		return false
	}
	r, _ := utf8.DecodeRuneInString(s)
	return isUpper(r)
}

func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func (p *Parser) ParseImport() (*model2.Import, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TIMPORT {
		return nil, p.unexpected(`"import"`)
	}
	startPos := p.lex.Pos

	var modifier model2.ImportModifier
	p.lex.NextKeywordOrStrLit()
	switch p.lex.Token {
	case scanner2.TPUBLIC:
		modifier = model2.ImportModifierPublic
	case scanner2.TWEAK:
		modifier = model2.ImportModifierWeak
	case scanner2.TSTRLIT:
		modifier = model2.ImportModifierNone
		p.lex.UnNext()
	}

	p.lex.NextStrLit()
	if p.lex.Token != scanner2.TSTRLIT {
		return nil, p.unexpected("strLit")
	}
	location := p.lex.Text

	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}

	return &model2.Import{
		Modifier: modifier,
		Location: location,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
func (p *Parser) MaybeScanInlineComment(hasSetter model2.HasInlineCommentSetter) {
	inlineComment := p.parseInlineComment()
	if inlineComment == nil {
		return
	}
	hasSetter.SetComment(inlineComment)
}
func (p *Parser) parseInlineComment() *model2.Comment {
	currentPos := p.lex.Pos
	comment, err := p.parseComment()
	if err != nil {
		return nil
	}
	if currentPos.Line != comment.Meta.Pos.Line {
		p.lex.UnNext()
		return nil
	}
	return comment
}
func (p *Parser) ParsePackage() (*model2.Package, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TPACKAGE {
		return nil, p.unexpected("package")
	}
	startPos := p.lex.Pos

	ident, _, err := p.lex.ReadFullIdent()
	if err != nil {
		return nil, p.unexpected("fullIdent")
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}

	return &model2.Package{
		Name: ident,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
func (p *Parser) ParseMessage() (*model2.Message, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TMESSAGE {
		return nil, p.unexpected("message")
	}
	startPos := p.lex.Pos

	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("messageName")
	}
	messageName := p.lex.Text

	messageBody, inlineLeftCurly, lastPos, err := p.parseMessageBody()
	if err != nil {
		return nil, err
	}

	return &model2.Message{
		MessageName: messageName,
		MessageBody: messageBody,
		CommentBLC:  inlineLeftCurly,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: lastPos.Position,
		},
	}, nil
}

func (p *Parser) parseMessageBody() (
	[]model2.Visitee,
	*model2.Comment,
	scanner2.Position,
	error,
) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTCURLY {
		return nil, nil, scanner2.Position{}, p.unexpected("{")
	}
	inlineLeftCurly := p.parseInlineComment()
	p.lex.Next()
	if p.lex.Token == scanner2.TRIGHTCURLY {
		lastPos := p.lex.Pos
		if p.permissive {
			// accept a block followed by semicolon. See https://github.com/yoheimuta/go-protoparser/v4/issues/30.
			p.lex.ConsumeToken(scanner2.TSEMICOLON)
			if p.lex.Token == scanner2.TSEMICOLON {
				lastPos = p.lex.Pos
			}
		}

		return nil, nil, lastPos, nil
	}
	p.lex.UnNext()
	var stmts []model2.Visitee
	for {
		comments := p.ParseComments()
		p.lex.NextKeyword()
		token := p.lex.Token
		p.lex.UnNext()
		var stmt interface {
			model2.HasInlineCommentSetter
			model2.Visitee
		}
		switch token {
		case scanner2.TRIGHTCURLY:
			if p.bodyIncludingComments {
				for _, comment := range comments {
					stmts = append(stmts, model2.Visitee(comment))
				}
			}
			p.lex.Next()

			lastPos := p.lex.Pos
			if p.permissive {
				// accept a block followed by semicolon. See https://github.com/yoheimuta/go-protoparser/v4/issues/30.
				p.lex.ConsumeToken(scanner2.TSEMICOLON)
				if p.lex.Token == scanner2.TSEMICOLON {
					lastPos = p.lex.Pos
				}
			}
			return stmts, inlineLeftCurly, lastPos, nil
		case scanner2.TENUM:
			enum, err := p.ParseEnum()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			enum.Comments = comments
			stmt = enum
		case scanner2.TMESSAGE:
			message, err := p.ParseMessage()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			message.Comments = comments
			stmt = message
		case scanner2.TOPTION:
			option, err := p.ParseOption()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			option.Comments = comments
			stmt = option
		case scanner2.TONEOF:
			oneof, err := p.ParseOneof()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			oneof.Comments = comments
			stmt = oneof
		case scanner2.TMAP:
			mapField, err := p.ParseMapField()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			mapField.Comments = comments
			stmt = mapField
		case scanner2.TEXTEND:
			extend, err := p.ParseExtend()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			extend.Comments = comments
			stmt = extend
		case scanner2.TRESERVED:
			reserved, err := p.ParseReserved()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			reserved.Comments = comments
			stmt = reserved
		case scanner2.TEXTENSIONS:
			extensions, err := p.ParseExtensions()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			extensions.Comments = comments
			stmt = extensions
		default:
			var ferr error
			isGroup := p.peekIsGroup()
			if isGroup {
				groupField, groupErr := p.ParseGroupField()
				if groupErr == nil {
					groupField.Comments = comments
					stmt = groupField
					break
				}
				ferr = groupErr
				p.lex.UnNext()
			} else {
				field, fieldErr := p.ParseField()
				if fieldErr == nil {
					field.Comments = comments
					stmt = field
					break
				}
				ferr = fieldErr
				p.lex.UnNext()
			}

			emptyErr := p.lex.ReadEmptyStatement()
			if emptyErr == nil {
				stmt = &model2.EmptyStatement{}
				break
			}

			return nil, nil, scanner2.Position{}, &parseMessageBodyStatementErr{
				parseFieldErr:          ferr,
				parseEmptyStatementErr: emptyErr,
			}
		}

		p.MaybeScanInlineComment(stmt)
		stmts = append(stmts, stmt)
	}
}

func (p *Parser) ParseMapField() (*model2.MapField, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TMAP {
		return nil, p.unexpected("map")
	}
	startPos := p.lex.Pos
	p.lex.Next()
	if p.lex.Token != scanner2.TLESS {
		return nil, p.unexpected("<")
	}
	keyType, err := p.parseKeyType()
	if err != nil {
		return nil, err
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TCOMMA {
		return nil, p.unexpected(",")
	}
	typeValue, _, err := p.parseType()
	if err != nil {
		return nil, p.unexpected("type")
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TGREATER {
		return nil, p.unexpected(">")
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("mapName")
	}
	mapName := p.lex.Text
	p.lex.Next()
	if p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	}
	fieldNumber, err := p.parseFieldNumber()
	if err != nil {
		return nil, p.unexpected("fieldNumber")
	}
	fieldOptions, err := p.parseFieldOptionsOption()
	if err != nil {
		return nil, err
	}
	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}
	return &model2.MapField{
		KeyType:      keyType,
		Type:         typeValue,
		MapName:      mapName,
		FieldNumber:  fieldNumber,
		FieldOptions: fieldOptions,
		Meta:         meta.Meta{Pos: startPos.Position},
	}, nil
}

var keyTypeConstants = map[string]struct{}{
	"int32":    {},
	"int64":    {},
	"uint32":   {},
	"uint64":   {},
	"sint32":   {},
	"sint64":   {},
	"fixed32":  {},
	"fixed64":  {},
	"sfixed32": {},
	"sfixed64": {},
	"bool":     {},
	"string":   {},
}

func (p *Parser) parseKeyType() (string, error) {
	p.lex.Next()
	if _, ok := keyTypeConstants[p.lex.Text]; ok {
		return p.lex.Text, nil
	}
	return "", p.unexpected("keyType constant")
}

func (p *Parser) ParseService() (*model2.Service, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TSERVICE {
		return nil, p.unexpected("service")
	}
	startPos := p.lex.Pos
	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("serviceName")
	}
	serviceName := p.lex.Text
	serviceBody, inlineLeftCurly, lastPos, err := p.parseServiceBody()
	if err != nil {
		return nil, err
	}
	return &model2.Service{
		Name:       serviceName,
		Body:       serviceBody,
		CommentBLC: inlineLeftCurly,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: lastPos.Position,
		},
	}, nil
}

func (p *Parser) parseServiceBody() (
	[]model2.Visitee,
	*model2.Comment,
	scanner2.Position,
	error,
) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTCURLY {
		return nil, nil, scanner2.Position{}, p.unexpected("{")
	}

	inlineLeftCurly := p.parseInlineComment()

	var stmts []model2.Visitee
	for {
		comments := p.ParseComments()

		p.lex.NextKeyword()
		token := p.lex.Token
		p.lex.UnNext()

		var stmt interface {
			model2.HasInlineCommentSetter
			model2.Visitee
		}

		switch token {
		case scanner2.TRIGHTCURLY:
			if p.bodyIncludingComments {
				for _, comment := range comments {
					stmts = append(stmts, model2.Visitee(comment))
				}
			}
			p.lex.Next()

			lastPos := p.lex.Pos
			if p.permissive {
				// accept a block followed by semicolon. See https://github.com/yoheimuta/go-protoparser/v4/issues/30.
				p.lex.ConsumeToken(scanner2.TSEMICOLON)
				if p.lex.Token == scanner2.TSEMICOLON {
					lastPos = p.lex.Pos
				}
			}
			return stmts, inlineLeftCurly, lastPos, nil
		case scanner2.TOPTION:
			option, err := p.ParseOption()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			option.Comments = comments
			stmt = option
		case scanner2.TRPC:
			rpc, err := p.parseRPC()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
			rpc.Comments = comments
			stmt = rpc
		default:
			err := p.lex.ReadEmptyStatement()
			if err != nil {
				return nil, nil, scanner2.Position{}, err
			}
		}

		p.MaybeScanInlineComment(stmt)
		stmts = append(stmts, stmt)
	}
}
func (p *Parser) parseRPC() (*model2.RPC, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TRPC {
		return nil, p.unexpected("rpc")
	}
	startPos := p.lex.Pos

	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("serviceName")
	}
	rpcName := p.lex.Text

	rpcRequest, err := p.parseRPCRequest()
	if err != nil {
		return nil, err
	}

	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TRETURNS {
		return nil, p.unexpected("returns")
	}

	rpcResponse, err := p.parseRPCResponse()
	if err != nil {
		return nil, err
	}

	var opts []*model2.Option
	var inlineLeftCurly *model2.Comment
	p.lex.Next()
	lastPos := p.lex.Pos
	switch p.lex.Token {
	case scanner2.TLEFTCURLY:
		p.lex.UnNext()
		opts, inlineLeftCurly, err = p.parseRPCOptions()
		if err != nil {
			return nil, err
		}
		lastPos = p.lex.Pos
		if p.permissive {
			// accept a block followed by semicolon. See https://github.com/yoheimuta/go-protoparser/v4/issues/30.
			p.lex.ConsumeToken(scanner2.TSEMICOLON)
			if p.lex.Token == scanner2.TSEMICOLON {
				lastPos = p.lex.Pos
			}
		}
	case scanner2.TSEMICOLON:
		break
	default:
		return nil, p.unexpected("{ or ;")
	}

	return &model2.RPC{
		Name:       rpcName,
		Request:    rpcRequest,
		Response:   rpcResponse,
		Options:    opts,
		CommentBLC: inlineLeftCurly,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: lastPos.Position,
		},
	}, nil
}
func (p *Parser) parseRPCRequest() (*model2.Request, error) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTPAREN {
		return nil, p.unexpected("(")
	}
	startPos := p.lex.Pos

	p.lex.NextKeyword()
	isStream := true
	if p.lex.Token != scanner2.TSTREAM {
		isStream = false
		p.lex.UnNext()
	}

	messageType, _, err := p.lex.ReadMessageType()
	if err != nil {
		return nil, err
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TRIGHTPAREN {
		return nil, p.unexpected(")")
	}

	return &model2.Request{
		IsStream:    isStream,
		MessageType: messageType,
		Meta:        meta.Meta{Pos: startPos.Position},
	}, nil
}
func (p *Parser) parseRPCResponse() (*model2.Response, error) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTPAREN {
		return nil, p.unexpected("(")
	}
	startPos := p.lex.Pos

	p.lex.NextKeyword()
	isStream := true
	if p.lex.Token != scanner2.TSTREAM {
		isStream = false
		p.lex.UnNext()
	}

	messageType, _, err := p.lex.ReadMessageType()
	if err != nil {
		return nil, err
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TRIGHTPAREN {
		return nil, p.unexpected(")")
	}

	return &model2.Response{
		IsStream: isStream,
		Type:     messageType,
		Meta:     meta.Meta{Pos: startPos.Position},
	}, nil
}
func (p *Parser) parseRPCOptions() ([]*model2.Option, *model2.Comment, error) {
	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTCURLY {
		return nil, nil, p.unexpected("{")
	}

	inlineLeftCurly := p.parseInlineComment()

	var options []*model2.Option
	for {
		p.lex.NextKeyword()
		token := p.lex.Token
		p.lex.UnNext()

		switch token {
		case scanner2.TOPTION:
			option, err := p.ParseOption()
			if err != nil {
				return nil, nil, err
			}
			options = append(options, option)
		case scanner2.TRIGHTCURLY:
			// This spec is not documented, but allowed in general.
			break
		default:
			err := p.lex.ReadEmptyStatement()
			if err != nil {
				return nil, nil, err
			}
		}

		p.lex.Next()
		if p.lex.Token == scanner2.TRIGHTCURLY {
			return options, inlineLeftCurly, nil
		}
		p.lex.UnNext()
	}
}
func (p *Parser) ParseOption() (*model2.Option, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TOPTION {
		return nil, p.unexpected("option")
	}
	startPos := p.lex.Pos

	optionName, err := p.parseOptionName()
	if err != nil {
		return nil, err
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	}

	constant, err := p.parseOptionConstant()
	if err != nil {
		return nil, err
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}

	return &model2.Option{
		Name:     optionName,
		Constant: constant,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
func (p *Parser) parseCloudEndpointsOptionConstant() (string, error) {
	var ret string

	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTCURLY {
		return "", p.unexpected("{")
	}
	ret += p.lex.Text

	for {
		p.lex.Next()
		if p.lex.Token != scanner2.TIDENT {
			return "", p.unexpected("ident")
		}
		ret += p.lex.Text

		needSemi := false
		p.lex.Next()
		switch p.lex.Token {
		case scanner2.TLEFTCURLY:
			if !p.permissive {
				return "", p.unexpected(":")
			}
			p.lex.UnNext()
		case scanner2.TCOLON:
			ret += p.lex.Text
			if p.lex.Peek() == scanner2.TLEFTCURLY && p.permissive {
				needSemi = true
			}
		default:
			if p.permissive {
				return "", p.unexpected("{ or :")
			}
			return "", p.unexpected(":")
		}

		constant, err := p.parseOptionConstant()
		if err != nil {
			return "", err
		}
		ret += constant

		p.lex.Next()
		if p.lex.Token == scanner2.TSEMICOLON && needSemi && p.permissive {
			ret += p.lex.Text
			p.lex.Next()
		}

		switch {
		case p.lex.Token == scanner2.TCOMMA, p.lex.Token == scanner2.TSEMICOLON:
			ret += p.lex.Text
			if p.lex.Peek() == scanner2.TRIGHTCURLY && p.permissive {
				p.lex.Next()
				ret += p.lex.Text
				return ret, nil
			}
		case p.lex.Token == scanner2.TRIGHTCURLY:
			ret += p.lex.Text
			return ret, nil
		default:
			ret += "\n"
			p.lex.UnNext()
		}
	}
}
func (p *Parser) parseOptionName() (string, error) {
	var optionName string
	p.lex.Next()
	switch p.lex.Token {
	case scanner2.TIDENT:
		optionName = p.lex.Text
	case scanner2.TLEFTPAREN:
		optionName = p.lex.Text

		// protoc accepts "(." fullIndent ")". See #63
		if p.permissive {
			p.lex.Next()
			if p.lex.Token == scanner2.TDOT {
				optionName += "."
			} else {
				p.lex.UnNext()
			}
		}

		fullIdent, _, err := p.lex.ReadFullIdent()
		if err != nil {
			return "", err
		}
		optionName += fullIdent

		p.lex.Next()
		if p.lex.Token != scanner2.TRIGHTPAREN {
			return "", p.unexpected(")")
		}
		optionName += p.lex.Text
	default:
		return "", p.unexpected("ident or left paren")
	}

	for {
		p.lex.Next()
		if p.lex.Token != scanner2.TDOT {
			p.lex.UnNext()
			break
		}
		optionName += p.lex.Text

		p.lex.Next()
		if p.lex.Token != scanner2.TIDENT {
			return "", p.unexpected("ident")
		}
		optionName += p.lex.Text
	}
	return optionName, nil
}
func (p *Parser) parseOptionConstant() (constant string, err error) {
	switch p.lex.Peek() {
	case scanner2.TLEFTCURLY:
		if !p.permissive {
			return "", p.unexpected("constant or permissive mode")
		}

		// parses empty fields within an option
		if p.lex.PeekN(2) == scanner2.TRIGHTCURLY {
			p.lex.NextN(2)
			return "{}", nil
		}

		constant, err = p.parseCloudEndpointsOptionConstant()
		if err != nil {
			return "", err
		}

	case scanner2.TLEFTSQUARE:
		if !p.permissive {
			return "", p.unexpected("constant or permissive mode")
		}
		p.lex.Next()

		// parses empty fields within an option
		if p.lex.Peek() == scanner2.TRIGHTSQUARE {
			p.lex.Next()
			return "[]", nil
		}

		constant, err = p.parseOptionConstants()
		if err != nil {
			return "", err
		}
		p.lex.Next()
		constant = "[" + constant + "]"

	default:
		constant, _, err = p.lex.ReadConstant(p.permissive)
		if err != nil {
			return "", err
		}
	}
	return constant, nil
}
func (p *Parser) parseOptionConstants() (constant string, err error) {
	opt, err := p.parseOptionConstant()
	if err != nil {
		return "", err
	}
	var opts []string
	opts = append(opts, opt)
	for {
		p.lex.Next()
		if p.lex.Token != scanner2.TCOMMA {
			p.lex.UnNext()
			break
		}
		opt, err = p.parseOptionConstant()
		if err != nil {
			return "", p.unexpected("optionConstant")
		}
		opts = append(opts, opt)
	}
	return strings.Join(opts, ","), nil
}

func (p *Parser) ParseOneof() (*model2.OneOf, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TONEOF {
		return nil, p.unexpected("oneof")
	}
	startPos := p.lex.Pos
	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("oneofName")
	}
	oneofName := p.lex.Text

	p.lex.Next()
	if p.lex.Token != scanner2.TLEFTCURLY {
		return nil, p.unexpected("{")
	}

	inlineLeftCurly := p.parseInlineComment()

	var oneofFields []*model2.OneOfField
	var options []*model2.Option
	for {
		comments := p.ParseComments()

		err := p.lex.ReadEmptyStatement()
		if err == nil {
			continue
		}

		p.lex.NextKeyword()
		token := p.lex.Token
		p.lex.UnNext()
		if token == scanner2.TOPTION {
			// See https://github.com/yoheimuta/go-protoparser/issues/57
			option, err := p.ParseOption()
			if err != nil {
				return nil, err
			}
			option.Comments = comments
			p.MaybeScanInlineComment(option)
			options = append(options, option)
		} else {
			oneofField, err := p.parseOneofField()
			if err != nil {
				return nil, err
			}
			oneofField.Comments = comments
			p.MaybeScanInlineComment(oneofField)
			oneofFields = append(oneofFields, oneofField)
		}

		p.lex.Next()
		if p.lex.Token == scanner2.TRIGHTCURLY {
			break
		} else {
			p.lex.UnNext()
		}
	}

	lastPos := p.lex.Pos
	if p.permissive {
		// accept a block followed by semicolon. See https://github.com/yoheimuta/go-protoparser/v4/issues/30.
		p.lex.ConsumeToken(scanner2.TSEMICOLON)
		if p.lex.Token == scanner2.TSEMICOLON {
			lastPos = p.lex.Pos
		}
	}

	return &model2.OneOf{
		OneofFields: oneofFields,
		OneofName:   oneofName,
		Options:     options,
		CommentBLC:  inlineLeftCurly,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: lastPos.Position,
		},
	}, nil
}
func (p *Parser) parseOneofField() (*model2.OneOfField, error) {
	typeValue, startPos, err := p.parseType()
	if err != nil {
		return nil, p.unexpected("type")
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TIDENT {
		return nil, p.unexpected("fieldName")
	}
	fieldName := p.lex.Text

	p.lex.Next()
	if p.lex.Token != scanner2.TEQUALS {
		return nil, p.unexpected("=")
	}

	fieldNumber, err := p.parseFieldNumber()
	if err != nil {
		return nil, p.unexpected("fieldNumber")
	}

	fieldOptions, err := p.parseFieldOptionsOption()
	if err != nil {
		return nil, err
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}

	return &model2.OneOfField{
		Type:    typeValue,
		Name:    fieldName,
		Number:  fieldNumber,
		Options: fieldOptions,
		Meta:    meta.Meta{Pos: startPos.Position},
	}, nil
}

func (p *Parser) ParseProto() (*model2.Proto, error) {
	syntaxComments := p.ParseComments()
	syntax, err := p.ParseSyntax()
	if err != nil {
		return nil, err
	}
	syntax.Comments = syntaxComments
	p.MaybeScanInlineComment(syntax)

	protoBody, err := p.parseProtoBody()
	if err != nil {
		return nil, err
	}

	return &model2.Proto{
		Syntax:    syntax,
		ProtoBody: protoBody,
		Meta: &model2.ProtoMeta{
			Filename: p.lex.Pos.Filename,
		},
	}, nil
}
func (p *Parser) parseProtoBody() ([]model2.Visitee, error) {
	var protoBody []model2.Visitee
	for {
		comments := p.ParseComments()
		if p.IsEOF() {
			if p.bodyIncludingComments {
				for _, comment := range comments {
					protoBody = append(protoBody, model2.Visitee(comment))
				}
			}
			return protoBody, nil
		}
		p.lex.NextKeyword()
		token := p.lex.Token
		p.lex.UnNext()
		var stmt interface {
			model2.HasInlineCommentSetter
			model2.Visitee
		}
		switch token {
		case scanner2.TIMPORT:
			importValue, err := p.ParseImport()
			if err != nil {
				return nil, err
			}
			importValue.Comments = comments
			stmt = importValue
		case scanner2.TPACKAGE:
			packageValue, err := p.ParsePackage()
			if err != nil {
				return nil, err
			}
			packageValue.Comments = comments
			stmt = packageValue
		case scanner2.TOPTION:
			option, err := p.ParseOption()
			if err != nil {
				return nil, err
			}
			option.Comments = comments
			stmt = option
		case scanner2.TMESSAGE:
			message, err := p.ParseMessage()
			if err != nil {
				return nil, err
			}
			message.Comments = comments
			stmt = message
		case scanner2.TENUM:
			enum, err := p.ParseEnum()
			if err != nil {
				return nil, err
			}
			enum.Comments = comments
			stmt = enum
		case scanner2.TSERVICE:
			service, err := p.ParseService()
			if err != nil {
				return nil, err
			}
			service.Comments = comments
			stmt = service
		case scanner2.TEXTEND:
			extend, err := p.ParseExtend()
			if err != nil {
				return nil, err
			}
			extend.Comments = comments
			stmt = extend
		default:
			err := p.lex.ReadEmptyStatement()
			if err != nil {
				return nil, err
			}
			protoBody = append(protoBody, &model2.EmptyStatement{})
		}
		p.MaybeScanInlineComment(stmt)
		protoBody = append(protoBody, stmt)
	}
}
func (p *Parser) ParseReserved() (*model2.Reserved, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner2.TRESERVED {
		return nil, p.unexpected("reserved")
	}
	startPos := p.lex.Pos

	parse := func() ([]*model2.Range, []string, error) {
		ranges, err := p.parseRanges()
		if err == nil {
			return ranges, nil, nil
		}
		fieldNames, ferr := p.parseFieldNames()
		if ferr == nil {
			return nil, fieldNames, nil
		}
		return nil, nil, &parseReservedErr{
			parseRangesErr:     err,
			parseFieldNamesErr: ferr,
		}
	}

	ranges, fieldNames, err := parse()
	if err != nil {
		return nil, err
	}

	p.lex.Next()
	if p.lex.Token != scanner2.TSEMICOLON {
		return nil, p.unexpected(";")
	}
	return &model2.Reserved{
		Ranges:     ranges,
		FieldNames: fieldNames,
		Meta:       meta.Meta{Pos: startPos.Position},
	}, nil
}
func (p *Parser) parseRanges() ([]*model2.Range, error) {
	var ranges []*model2.Range
	rangeValue, err := p.parseRange()
	if err != nil {
		return nil, err
	}
	ranges = append(ranges, rangeValue)

	for {
		p.lex.Next()
		if p.lex.Token != scanner2.TCOMMA {
			p.lex.UnNext()
			break
		}

		rangeValue, err := p.parseRange()
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, rangeValue)
	}
	return ranges, nil
}
func (p *Parser) parseRange() (*model2.Range, error) {
	p.lex.NextNumberLit()
	if p.lex.Token != scanner2.TINTLIT {
		p.lex.UnNext()
		return nil, p.unexpected("intLit")
	}
	begin := p.lex.Text

	p.lex.Next()
	if p.lex.Text != "to" {
		p.lex.UnNext()
		return &model2.Range{
			Begin: begin,
		}, nil
	}

	p.lex.NextNumberLit()
	switch {
	case p.lex.Token == scanner2.TINTLIT,
		p.lex.Text == "max":
		return &model2.Range{
			Begin: begin,
			End:   p.lex.Text,
		}, nil
	default:
		break
	}
	return nil, p.unexpected(`"intLit | "max"`)
}
func (p *Parser) parseFieldNames() ([]string, error) {
	var fieldNames []string

	fieldName, err := p.parseQuotedFieldName()
	if err != nil {
		return nil, err
	}
	fieldNames = append(fieldNames, fieldName)

	for {
		p.lex.Next()
		if p.lex.Token != scanner2.TCOMMA {
			p.lex.UnNext()
			break
		}

		fieldName, err = p.parseQuotedFieldName()
		if err != nil {
			return nil, err
		}
		fieldNames = append(fieldNames, fieldName)
	}
	return fieldNames, nil
}
func (p *Parser) parseQuotedFieldName() (string, error) {
	p.lex.NextStrLit()
	if p.lex.Token != scanner2.TSTRLIT {
		p.lex.UnNext()
		return "", p.unexpected("quotedFieldName")
	}
	return p.lex.Text, nil
}
