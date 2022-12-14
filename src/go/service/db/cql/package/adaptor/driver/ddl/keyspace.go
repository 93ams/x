package ddl

import (
	"strings"
)

type KeySpaceStmt struct {
	verb        DDLVerb
	Name        string
	Durable     *bool
	Exists      *bool
	Tags        Object
	Replication Object
}

func CreateKeySpace(name string, replication Object) KeySpaceStmt {
	return KeySpaceStmt{verb: VerbCreate, Name: name, Replication: replication}
}
func AlterKeySpace(name string) KeySpaceStmt {
	return KeySpaceStmt{verb: VerbAlter, Name: name}
}
func DropKeySpace(name string) KeySpaceStmt {
	return KeySpaceStmt{verb: VerbDrop, Name: name}
}

func (s KeySpaceStmt) WithDurable(b bool) KeySpaceStmt {
	s.Durable = &b
	return s
}
func (s KeySpaceStmt) WithExists(b bool) KeySpaceStmt {
	s.Exists = &b
	return s
}

func (s KeySpaceStmt) String() string {
	words := []string{string(s.verb), "KEYSPACE", s.Name}
	words = writeExists(words, s.Exists)
	if s.verb != VerbDrop {
		words = append(words, "WITH", "REPLICATION", "=", s.Replication.String())
		if s.Durable != nil {
			words = append(words, "AND", "DURABLE_WRITES", "=")
			if *s.Durable {
				words = append(words, "true")
			} else {
				words = append(words, "false")
			}
		}
		if s.Tags != nil {
			words = append(words, "AND", "TAGS", "=", s.Tags.String())
		}
	}
	return strings.Join(words, " ")
}

func writeExists(in []string, exists *bool) []string {
	if exists != nil {
		in = append(in, "IF")
		if !*exists {
			in = append(in, "NOT")
		}
		in = append(in, "EXISTS")
	}
	return in
}
