package scanner

func (s *Scanner) isEOF() bool { return s.peek() == eof }
func isBoolLit(ident string) bool {
	switch ident {
	case "true", "false":
		return true
	default:
		return false
	}
}
func isQuote(ch rune) bool {
	switch ch {
	case '\'', '"':
		return true
	default:
		return false
	}
}
func isLetter(r rune) bool {
	if r < 'A' {
		return false
	}

	if r > 'z' {
		return false
	}

	if r > 'Z' && r < 'a' {
		return false
	}

	return true
}
func isDecimalDigit(r rune) bool { return '0' <= r && r <= '9' }
func isOctalDigit(r rune) bool   { return '0' <= r && r <= '7' }
func isHexDigit(r rune) bool {
	if '0' <= r && r <= '9' {
		return true
	}
	if 'A' <= r && r <= 'F' {
		return true
	}
	if 'a' <= r && r <= 'f' {
		return true
	}
	return false
}
func isFloatLitKeyword(ident string) bool {
	switch ident {
	case "inf", "nan":
		return true
	default:
		return false
	}
}
