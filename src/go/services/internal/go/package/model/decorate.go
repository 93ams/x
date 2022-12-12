package model

type (
	SpaceType   int
	Decorations []string
	NodeDecs    struct {
		Before SpaceType
		Start  Decorations
		End    Decorations
		After  SpaceType
	}
)

func (d *Decorations) Append(decs ...string)  { *d = append(*d, decs...) }
func (d *Decorations) Prepend(decs ...string) { *d = append(append([]string{}, decs...), *d...) }
func (d *Decorations) Replace(decs ...string) { *d = append([]string{}, decs...) }
func (d *Decorations) Clear()                 { *d = nil }
func (d *Decorations) All() []string          { return *d }

const (
	None      SpaceType = 0
	NewLine   SpaceType = 1
	EmptyLine SpaceType = 2
)

func (s SpaceType) String() string {
	switch s {
	case None:
		return "None"
	case NewLine:
		return "NewLine"
	case EmptyLine:
		return "EmptyLine"
	}
	return ""
}
