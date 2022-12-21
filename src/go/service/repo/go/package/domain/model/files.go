package model

type (
	File struct {
		Pkg, Name string
	}
	Package struct {
		Name  string
		Files []File
	}
	Module struct {
		Name     string
		Packages []Package
	}
)
