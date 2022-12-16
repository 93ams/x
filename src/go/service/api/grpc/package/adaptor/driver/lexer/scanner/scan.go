package scanner

import (
	"grpc/package/domain/model/meta"
	"runtime"
)

type Mode uint

const (
	ScanIdent Mode = 1 << iota
	ScanNumberLit
	ScanStrLit
	ScanBoolLit
	ScanKeyword
	ScanComment
	ScanLit = ScanNumberLit | ScanStrLit | ScanBoolLit
)

func (s *Scanner) scanComment() (string, error) {
	lit := string(s.read())
	ch := s.read()
	switch ch {
	case '/':
		for ch != '\n' {
			lit += string(ch)
			if s.isEOF() {
				return lit, nil
			}
			ch = s.read()
		}
	case '*':
		for {
			if s.isEOF() {
				return lit, s.unexpected(eof, "\n")
			}
			lit += string(ch)
			ch = s.read()
			chn := s.peek()
			if ch == '*' && chn == '/' {
				lit += string(ch)
				lit += string(s.read())
				break
			}
		}
	default:
		return "", s.unexpected(ch, "/ or *")
	}

	return lit, nil
}
func (s *Scanner) scanIdent() string {
	ident := string(s.read())
	for {
		next := s.peek()
		switch {
		case isLetter(next), isDecimalDigit(next), next == '_':
			ident += string(s.read())
		default:
			return ident
		}
	}
}
func (s *Scanner) unexpected(found rune, expected string) error {
	_, file, line, _ := runtime.Caller(1)
	err := &meta.Error{
		Pos:      s.pos.Position,
		Expected: expected,
		Found:    string(found),
	}
	err.SetOccured(file, line)
	return err
}
func (s *Scanner) scanNumberLit() (Token, string, error) {
	lit := string(s.read())
	ch := s.peek()

	switch {
	case lit == "0" && (ch == 'x' || ch == 'X'):
		lit += string(s.read())
		if !isHexDigit(s.peek()) {
			return TILLEGAL, "", s.unexpected(s.peek(), "hexDigit")
		}
		lit += string(s.read())

		for !s.isEOF() {
			if !isHexDigit(s.peek()) {
				break
			}
			lit += string(s.read())
		}
		return TINTLIT, lit, nil
	case lit == ".":
		fractional, err := s.scanFractionPartNoOmit()
		if err != nil {
			return TILLEGAL, "", err
		}
		return TFLOATLIT, lit + fractional, nil
	case ch == '.':
		lit += string(s.read())
		fractional, err := s.scanFractionPart()
		if err != nil {
			return TILLEGAL, "", err
		}
		return TFLOATLIT, lit + fractional, nil
	case ch == 'e' || ch == 'E':
		exp, err := s.scanExponent()
		if err != nil {
			return TILLEGAL, "", err
		}
		return TFLOATLIT, lit + exp, nil
	case lit == "0":
		for !s.isEOF() {
			if !isOctalDigit(s.peek()) {
				break
			}
			lit += string(s.read())
		}
		return TINTLIT, lit, nil
	default:
		for !s.isEOF() {
			if !isDecimalDigit(s.peek()) {
				break
			}
			lit += string(s.read())
		}
		switch s.peek() {
		case '.':
			lit += string(s.read())
			fractional, err := s.scanFractionPart()
			if err != nil {
				return TILLEGAL, "", err
			}
			return TFLOATLIT, lit + fractional, nil
		case 'e', 'E':
			exp, err := s.scanExponent()
			if err != nil {
				return TILLEGAL, "", err
			}
			return TFLOATLIT, lit + exp, nil
		default:
			return TINTLIT, lit, nil
		}
	}
}
func (s *Scanner) scanFractionPart() (string, error) {
	lit := ""
	ch := s.peek()
	switch {
	case isDecimalDigit(ch):
		decimals, err := s.scanDecimals()
		if err != nil {
			return "", err
		}
		lit += decimals
	}
	switch s.peek() {
	case 'e', 'E':
		exp, err := s.scanExponent()
		if err != nil {
			return "", err
		}
		lit += exp
	}
	return lit, nil
}
func (s *Scanner) scanFractionPartNoOmit() (string, error) {
	decimals, err := s.scanDecimals()
	if err != nil {
		return "", err
	}
	switch s.peek() {
	case 'e', 'E':
		exp, err := s.scanExponent()
		if err != nil {
			return "", err
		}
		return decimals + exp, nil
	default:
		return decimals, nil
	}
}
func (s *Scanner) scanExponent() (string, error) {
	ch := s.peek()
	switch ch {
	case 'e', 'E':
		lit := string(s.read())
		switch s.peek() {
		case '+', '-':
			lit += string(s.read())
		}
		decimals, err := s.scanDecimals()
		if err != nil {
			return "", err
		}
		return lit + decimals, nil
	default:
		return "", s.unexpected(ch, "e or E")
	}
}
func (s *Scanner) scanDecimals() (string, error) {
	ch := s.peek()
	if !isDecimalDigit(ch) {
		return "", s.unexpected(ch, "decimalDigit")
	}
	lit := string(s.read())

	for !s.isEOF() {
		if !isDecimalDigit(s.peek()) {
			break
		}
		lit += string(s.read())
	}
	return lit, nil
}

func (s *Scanner) scanStrLit() (string, error) {
	quote := s.read()
	lit := string(quote)

	ch := s.peek()
	for ch != quote {
		cv, err := s.scanCharValue()
		if err != nil {
			return "", err
		}
		lit += cv
		ch = s.peek()
	}

	// consume quote
	lit += string(s.read())
	return lit, nil
}
func (s *Scanner) scanCharValue() (string, error) {
	ch := s.peek()
	switch ch {
	case eof, '\n':
		return "", s.unexpected(ch, `/[^\0\n\\]`)
	case '\\':
		return s.tryScanEscape(), nil
	default:
		return string(s.read()), nil
	}
}
func (s *Scanner) tryScanEscape() string {
	lit := string(s.read())
	isCharEscape := func(r rune) bool {
		cs := []rune{'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '\'', '"'}
		for _, c := range cs {
			if r == c {
				return true
			}
		}
		return false
	}
	ch := s.peek()
	switch {
	case ch == 'x' || ch == 'X':
		lit += string(s.read())
		for i := 0; i < 2; i++ {
			if !isHexDigit(s.peek()) {
				return lit
			}
			lit += string(s.read())
		}
	case isOctalDigit(ch):
		for i := 0; i < 3; i++ {
			if !isOctalDigit(s.peek()) {
				return lit
			}
			lit += string(s.read())
		}
	case isCharEscape(ch):
		lit += string(s.read())
		return lit
	}
	return lit
}
