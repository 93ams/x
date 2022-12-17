package lexer

import (
	"grpc/package/adaptor/driver/lexer/scanner"
	"grpc/package/domain/model/meta"
	"runtime"
	"strings"
)

func (l *Lexer) ReadConstant(permissive bool) (string, scanner.Position, error) {
	l.NextLit()
	startPos, cons := l.Pos, l.Text
	switch {
	case l.Token == scanner.TSTRLIT:
		if permissive {
			return l.mergeMultilineStrLit(), startPos, nil
		}
		return cons, startPos, nil
	case l.Token == scanner.TBOOLLIT:
		return cons, startPos, nil
	case l.Token == scanner.TIDENT:
		l.UnNext()
		fullIdent, pos, err := l.ReadFullIdent()
		if err != nil {
			return "", scanner.Position{}, err
		}
		return fullIdent, pos, nil
	case l.Token == scanner.TINTLIT, l.Token == scanner.TFLOATLIT:
		return cons, startPos, nil
	case l.Text == "-" || l.Text == "+":
		l.NextLit()
		switch l.Token {
		case scanner.TINTLIT, scanner.TFLOATLIT:
			cons += l.Text
			return cons, startPos, nil
		default:
			return "", scanner.Position{}, l.unexpected(l.Text, "TINTLIT or TFLOATLIT")
		}
	default:
		return "", scanner.Position{}, l.unexpected(l.Text, "constant")
	}
}
func (l *Lexer) ReadEmptyStatement() error {
	l.Next()
	if l.Token == scanner.TSEMICOLON {
		return nil
	}
	l.UnNext()
	return l.unexpected(l.Text, ";")
}
func (l *Lexer) ReadEnumType() (string, scanner.Position, error) { return l.ReadMessageType() }
func (l *Lexer) ReadFullIdent() (string, scanner.Position, error) {
	if l.Next(); l.Token != scanner.TIDENT {
		return "", scanner.Position{}, l.unexpected(l.Text, "TIDENT")
	}
	startPos, fullIdent := l.Pos, l.Text
	l.Next()
	for !l.IsEOF() {
		if l.Token != scanner.TDOT {
			l.UnNext()
			break
		}
		if l.Next(); l.Token != scanner.TIDENT {
			return "", scanner.Position{}, l.unexpected(l.Text, "TIDENT")
		}
		fullIdent += "." + l.Text
		l.Next()
	}
	return fullIdent, startPos, nil
}
func (l *Lexer) ReadMessageType() (string, scanner.Position, error) {
	l.Next()
	startPos := l.Pos

	var messageType string
	if l.Token == scanner.TDOT {
		messageType = l.Text
	} else {
		l.UnNext()
	}

	l.Next()
	for !l.IsEOF() {
		if l.Token != scanner.TIDENT {
			return "", scanner.Position{}, l.unexpected(l.Text, "ident")
		}
		messageType += l.Text

		l.Next()
		if l.Token != scanner.TDOT {
			l.UnNext()
			break
		}
		messageType += l.Text

		l.Next()
	}

	return messageType, startPos, nil
}

func (l *Lexer) mergeMultilineStrLit() string {
	q := "'"
	if strings.HasPrefix(l.Text, "\"") {
		q = "\""
	}
	var b strings.Builder
	b.WriteString(q)
	for l.Token == scanner.TSTRLIT {
		strippedString := strings.Trim(l.Text, q)
		b.WriteString(strippedString)
		l.NextLit()
	}
	l.UnNext()
	b.WriteString(q)
	return b.String()
}
func (l *Lexer) unexpected(found, expected string) error {
	err := &meta.Error{
		Pos:      l.Pos.Position,
		Expected: expected,
		Found:    l.Text,
	}
	if l.debug {
		_, file, line, _ := runtime.Caller(1)
		err.SetOccured(file, line)
	}
	return err
}
