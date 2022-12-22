package parser

import (
	"fmt"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/adaptor/driver/model/meta"
	"runtime"
)

type (
	parseEnumBodyStatementErr struct {
		parseEnumFieldErr      error
		parseEmptyStatementErr error
	}
	parseExtendBodyStatementErr struct {
		parseFieldErr          error
		parseEmptyStatementErr error
	}
	parseMessageBodyStatementErr struct {
		parseFieldErr          error
		parseEmptyStatementErr error
	}
	parseReservedErr struct {
		parseRangesErr     error
		parseFieldNamesErr error
	}
)

func (p *Parser) unexpected(expected string) error {
	_, file, line, _ := runtime.Caller(1)
	err := &meta.Error{
		Found:    fmt.Sprintf("%q(Token=%v, Pos=%s)", p.lex.Text, p.lex.Token, p.lex.Pos),
		Pos:      p.lex.Pos.Position,
		Expected: expected,
	}
	err.SetOccured(file, line)
	return err
}
func (p *Parser) unexpectedf(format string, a ...interface{}) error {
	return p.unexpected(fmt.Sprintf(format, a...))
}
func (e *parseEnumBodyStatementErr) Error() string {
	return fmt.Sprintf("%v:%v",
		e.parseEnumFieldErr,
		e.parseEmptyStatementErr)
}
func (e *parseExtendBodyStatementErr) Error() string {
	return fmt.Sprintf("%v:%v", e.parseFieldErr, e.parseEmptyStatementErr)
}
func (e *parseMessageBodyStatementErr) Error() string {
	return fmt.Sprintf("%v:%v", e.parseFieldErr, e.parseEmptyStatementErr)
}
func (e *parseReservedErr) Error() string {
	return fmt.Sprintf("%v:%v", e.parseRangesErr, e.parseFieldNamesErr)
}
