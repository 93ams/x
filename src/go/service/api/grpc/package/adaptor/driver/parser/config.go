package parser

import (
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/adaptor/driver/lexer"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/adaptor/driver/model"
	"io"
)

type ParseConfig struct {
	debug                 bool
	permissive            bool
	bodyIncludingComments bool
	filename              string
}

func WithDebug(debug bool) x.Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.debug = debug }
}
func WithFilename(filename string) x.Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.filename = filename }
}
func WithPermissive(permissive bool) x.Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.permissive = permissive }
}
func WithBodyIncludingComments(bodyIncludingComments bool) x.Opt[*ParseConfig] {
	return func(c *ParseConfig) { c.bodyIncludingComments = bodyIncludingComments }
}
func Parse(input io.Reader, options ...x.Opt[*ParseConfig]) (*model.Proto, error) {
	config := x.Apply(&ParseConfig{permissive: true}, options)
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
