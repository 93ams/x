package model

type (
	ColumnKey struct {
		KeySpace string
		Table    string
		Name     string
	}
	Column struct {
		ColumnKey
		Order     string
		NameBytes string
		Kind      string
		Pos       int
		Type      string
		Static    bool
		Primary   bool
	}
)

func (k ColumnKey) String() string { return k.KeySpace + "." + k.Table + "." + k.Name }
func (k ColumnKey) Raw() any       { return k }
