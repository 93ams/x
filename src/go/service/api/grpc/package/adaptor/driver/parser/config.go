package parser

import (
	"grpc/package/adaptor/driver/lexer"
	"grpc/package/domain/model"
	"io"
)

type ParseConfig struct {
	debug                 bool
	permissive            bool
	bodyIncludingComments bool
	filename              string
}

func WithDebug(debug bool) Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.debug = debug }
}
func WithFilename(filename string) Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.filename = filename }
}
func WithPermissive(permissive bool) Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.permissive = permissive }
}
func WithBodyIncludingComments(bodyIncludingComments bool) Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.bodyIncludingComments = bodyIncludingComments }
}
func Parse(input io.Reader, options ...Opt[*ParseConfig]) (*model.Proto, error) {
	config := Apply(&ParseConfig{permissive: true}, options)
	p := NewParser(
		lexer.NewLexer(input,
			lexer.WithDebug(config.debug),
			lexer.WithFilename(config.filename),
		),
		Permissive(config.permissive),
		BodyIncludingComments(config.bodyIncludingComments),
	)
	return p.ParseProto()
}
