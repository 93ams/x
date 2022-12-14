package model

type (
	File struct {
		Name       string
		Imports    []Import
		Methods    []Method
		Structs    []Struct
		Interfaces []Interface
		Vars       []Var
	}
	Import struct {
		Name string
	}
	Method struct {
		MethodSignature
	}
	MethodSignature struct {
		Name string
	}
	Struct struct {
		Name    string
		Methods []Method
	}
	Interface struct {
		Name    string
		Methods []MethodSignature
	}
	Var struct {
		Name  string
		Type  string
		Const bool
	}
)
