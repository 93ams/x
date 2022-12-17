package lexer

import (
	"grpc/package/adaptor/driver/lexer/scanner"
	"io"
	"log"
	"path/filepath"
	"runtime"
)

type Lexer struct {
	Token       scanner.Token
	Text        string
	RawText     []rune
	Pos         scanner.Position
	Error       func(lexer *Lexer, err error)
	scanner     *scanner.Scanner
	scannerOpts []Opt[*scanner.Scanner]
	scanErr     error
	debug       bool
}

func WithDebug(debug bool) Opt[*Lexer] { return func(l *Lexer) { l.debug = debug } }
func WithFilename(filename string) Opt[*Lexer] {
	return func(l *Lexer) { l.scannerOpts = append(l.scannerOpts, scanner.WithFilename(filename)) }
}
func NewLexer(input io.Reader, opts ...Opt[*Lexer]) *Lexer {
	lex := new(Lexer)
	for _, opt := range opts {
		opt(lex)
	}
	lex.Error = func(_ *Lexer, err error) { log.Printf(`Lexer encountered the error "%v"`, err) }
	lex.scanner = scanner.NewScanner(input, lex.scannerOpts...)
	return lex
}
func (l *Lexer) Next() {
	defer func() {
		if l.debug {
			if _, file, line, ok := runtime.Caller(2); ok {
				log.Printf("[DEBUG] Text=[%s], Token=[%v], Pos=[%s] called from %s:%d\n",
					l.Text,
					l.Token,
					l.Pos,
					filepath.Base(file),
					line)
			}
		}
	}()
	var err error
	l.Token, l.Text, l.Pos, err = l.scanner.Scan()
	l.RawText = l.scanner.LastScanRaw()
	if err != nil {
		l.scanErr = err
		l.Error(l, err)
	}
}
func (l *Lexer) NextN(n int) {
	for 0 < n {
		l.Next()
		n--
	}
}
func (l *Lexer) NextKeywordOrStrLit() {
	l.nextWithSpecificMode(scanner.ScanKeyword | scanner.ScanStrLit)
}
func (l *Lexer) NextKeyword()   { l.nextWithSpecificMode(scanner.ScanKeyword) }
func (l *Lexer) NextStrLit()    { l.nextWithSpecificMode(scanner.ScanStrLit) }
func (l *Lexer) NextLit()       { l.nextWithSpecificMode(scanner.ScanLit) }
func (l *Lexer) NextNumberLit() { l.nextWithSpecificMode(scanner.ScanNumberLit) }
func (l *Lexer) NextComment()   { l.nextWithSpecificMode(scanner.ScanComment) }
func (l *Lexer) nextWithSpecificMode(nextMode scanner.Mode) {
	mode := l.scanner.Mode
	defer func() {
		l.scanner.Mode = mode
	}()

	l.scanner.Mode = nextMode
	l.Next()
}
func (l *Lexer) IsEOF() bool      { return l.Token == scanner.TEOF }
func (l *Lexer) LatestErr() error { return l.scanErr }
func (l *Lexer) Peek() scanner.Token {
	l.Next()
	defer l.UnNext()
	return l.Token
}
func (l *Lexer) PeekN(n int) scanner.Token {
	var lasts [][]rune
	for 0 < n {
		l.Next()
		lasts = append(lasts, l.RawText)
		n--
	}
	token := l.Token
	for i := len(lasts) - 1; 0 <= i; i-- {
		l.UnNextTo(lasts[i])
	}
	return token
}
func (l *Lexer) UnNext() {
	l.Pos = l.scanner.UnScan()
	l.Token = scanner.TILLEGAL
}
func (l *Lexer) UnNextTo(lastScan []rune) {
	l.scanner.SetLastScanRaw(lastScan)
	l.UnNext()
}
func (l *Lexer) ConsumeToken(t scanner.Token) {
	l.Next()
	if l.Token == t {
		return
	}
	l.UnNext()
}
