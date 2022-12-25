package meta

import "fmt"

type (
	Position struct {
		Filename string
		Offset   int
		Line     int
		Column   int
	}
	Meta struct {
		Pos     Position
		LastPos Position
	}
	Error struct {
		Pos       Position
		Expected  string
		Found     string
		occuredIn string
		occuredAt int
	}
)

func (e *Error) Error() string {
	if e.occuredAt == 0 && e.occuredIn == "" {
		return fmt.Sprintf("found %q but expected [%s]", e.Found, e.Expected)
	}
	return fmt.Sprintf("found %q but expected [%s] at %s:%d", e.Found, e.Expected, e.occuredIn, e.occuredAt)
}
func (e *Error) SetOccured(occuredIn string, occuredAt int) {
	e.occuredIn = occuredIn
	e.occuredAt = occuredAt
}
func (pos Position) String() string {
	s := pos.Filename
	if s == "" {
		s = "<input>"
	}
	s += fmt.Sprintf(":%d:%d", pos.Line, pos.Column)
	return s
}
