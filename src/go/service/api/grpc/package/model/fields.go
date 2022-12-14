package model

type (
	FieldCardinality int
	FieldType        string
)

const (
	CardinalityDefault FieldCardinality = iota
	CardinalityRequired
	CardinalityOptional
	CardinalityRepeated
)
const (
	TypeString FieldType = "string"
	TypeUint64 FieldType = "uint64"
)
