package model

type (
	ImportType int
)

const (
	ImportDefault ImportType = iota
	ImportPublic
	ImportWeak
)
