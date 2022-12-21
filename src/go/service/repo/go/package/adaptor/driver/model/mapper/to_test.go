package mapper_test

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/builder"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/mapper"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/service"
	"os"
	"testing"
)

func TestNewType(t *testing.T) {
	tests := []struct {
		name  string
		props model.Struct
		want  *builder.GenBuilder
	}{
		{props: model.Struct{
			Ident: model.Ident{
				Name: "Something",
				Path: "model",
			},
			Fields: []model.StructField{
				{Names: []string{"Foo"}, Type: model.Ident{Name: "string"}},
				{Names: []string{"Bar"}, Type: model.Ident{Path: "model", Name: "SomethingElse"}},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decls := mapper.NewBuilder(tt.props)
			lo.Must0(service.WriteDecls(os.Stdout, "test", decls...))
			//			wrapper.Print(lo.Must(wrapper.Parse(`package test
			//
			//func (b *SomethingBuilder) Bar(v model.SomethingElse) *SomethingBuilder {
			//	b.T.Bar = v
			//	return b
			//}
			//`)))
		})
	}
}
