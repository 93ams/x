package model

import "strings"

type (
	TypeFilter struct {
		Name    string
		Type    string
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
		Type TypeFilter
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
	return t.Name.Name == f.Name
}
