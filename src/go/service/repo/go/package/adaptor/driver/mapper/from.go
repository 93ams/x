package mapper

import (
	. "github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/wrapper/model"
)

func MapType(p *model.Type) any {
	return nil
}
func MapStruct(p *model.Struct) Struct {
	return Struct{
		Fields: nil,
	}
}
func MapInterface(p *model.Interface) Interface {
	return Interface{}
}
func MapDecl() {

}
