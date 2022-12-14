package builder

type (
	FileBuilder      struct{}
	MethodBuilder    struct{}
	StructBuilder    struct{}
	InterfaceBuilder struct{}
)

func File() *FileBuilder {
	return &FileBuilder{}
}
func Method() *MethodBuilder {
	return &MethodBuilder{}
}
func Struct() *StructBuilder {
	return &StructBuilder{}
}
func Interface() *InterfaceBuilder {
	return &InterfaceBuilder{}
}
