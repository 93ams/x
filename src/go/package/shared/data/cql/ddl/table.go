package ddl

import (
	. "github.com/tilau2328/cql/package/shared/data/cql"
	"strings"
)

type (
	Column struct {
		Name    string
		Type    string
		Primary bool
		Static  bool
	}
	TableStmt struct {
		KeySpace   string
		Name       string
		verb       DDLVerb
		Exists     *bool
		cols       []Column
		primaryKey []string
		Properties map[string]Object
	}
)

func CreateTable(keyspace, name string) TableStmt {
	return TableStmt{verb: VerbCreate, KeySpace: keyspace, Name: name}
}
func AlterTable(keyspace, name string) TableStmt {
	return TableStmt{verb: VerbAlter, KeySpace: keyspace, Name: name}
}
func DropTable(keyspace, name string) TableStmt {
	return TableStmt{verb: VerbDrop, KeySpace: keyspace, Name: name}
}
func (s TableStmt) Cols(c ...Column) TableStmt {
	s.cols = c
	for _, v := range c {
		if v.Primary {
			s.primaryKey = append(s.primaryKey, v.Name)
		}
	}
	return s
}
func (s TableStmt) String() string {
	words := []string{string(s.verb), "TABLE", s.KeySpace + "." + s.Name}
	words = writeExists(words, s.Exists)
	switch s.verb {
	case VerbAlter:
		words = append(words, "ADD")
		words = s.writeColumns(words)
		// Add Tag Opts
	case VerbCreate:
		words = s.writeColumns(words)
	case VerbDrop:
	}
	return strings.Join(words, " ")
}

func (s TableStmt) writeColumns(in []string) []string {
	in = append(in, "(")
	var cols []string
	for _, col := range s.cols {
		inner := []string{col.Name, col.Type}
		if col.Static {
			inner = append(inner, "STATIC")
		} else if col.Primary {
			if len(s.primaryKey) == 1 {
				inner = append(inner, "PRIMARY KEY")
			}
		}
		cols = append(cols, strings.Join(inner, " "))
	}
	if len(s.primaryKey) > 1 {
		cols = append(cols, "PRIMARY KEY ("+strings.Join(s.primaryKey, ",")+")")
	}
	in = append(in, strings.Join(cols, ", "))
	return append(in, ")")
}
