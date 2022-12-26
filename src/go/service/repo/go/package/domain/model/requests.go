package model

import "strings"

type (
	TypeFilter struct {
		Name, Type string
		Partial    bool
	}
	FuncFilter struct {
		Name    string
		Partial bool
	}
	FilePath struct {
		Name string
		Dir  string
	}
	ReadReq struct {
		FilePath
	}
	CreateReq struct {
		FilePath
		File *File
	}
	SearchReq struct {
		FilePath
		Id      string
		Type    *TypeFilter
		Func    *FuncFilter
		Package bool
	}
	ModifyReq struct {
		FilePath
	}
	DeleteReq struct {
		FilePath
	}
)

func (f TypeFilter) Filter(t *Type) bool {
	if f.Type != "" {
		switch t.Type.(type) {
		case *Struct:
			if f.Type != "struct" {
				return false
			}
		case *Interface:
			if f.Type != "interface" {
				return false
			}
		default:
			return false
		}
	}
	if f.Partial {
		return strings.Contains(t.Name.Name, f.Name)
	}
	return f.Name == "" || t.Name.Name == f.Name
}

func (f FuncFilter) Filter(t *Func) bool {
	if f.Partial {
		return strings.Contains(t.Name.Name, f.Name)
	}
	return f.Name == "" || t.Name.Name == f.Name
}
