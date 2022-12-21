package scanner

import (
	"grpc/package/wrapper/model/meta"
	"unicode/utf8"
)

type Position struct {
	meta.Position
	columns map[int]int
}

func NewPosition() *Position {
	return &Position{
		Position: meta.Position{Line: 1, Column: 1},
		columns:  make(map[int]int),
	}
}
func (pos Position) String() string { return pos.Position.String() }
func (pos *Position) Advance(r rune) {
	length := utf8.RuneLen(r)
	pos.Offset += length
	if r == '\n' {
		pos.columns[pos.Line] = pos.Column
		pos.Line++
		pos.Column = 1
	} else {
		pos.Column++
	}
}
func (pos Position) AdvancedBulk(s string) Position {
	for _, r := range s {
		pos.Advance(r)
	}
	last, _ := utf8.DecodeLastRuneInString(s)
	pos.Revert(last)
	return pos
}
func (pos *Position) Revert(r rune) {
	length := utf8.RuneLen(r)
	pos.Offset -= length
	if r == '\n' {
		pos.Line--
		pos.Column = pos.columns[pos.Line]
	} else {
		pos.Column--
	}
}
