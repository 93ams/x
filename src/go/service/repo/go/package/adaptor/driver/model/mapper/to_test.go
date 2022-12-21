package mapper_test

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/coding"
	model2 "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/builder"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/visitor/assert"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/model/visitor/filter"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"testing"
)

func TestNewType(t *testing.T) {
	tests := []struct {
		name  string
		props model.Service
		want  *builder.GenBuilder
	}{
		{props: model.Service{
			Struct: model.Struct{
				Ident: model.Ident{
					Name: "Something",
					Path: "model",
				},
			},
			Methods: []model.Func{
				{FuncType: model.FuncType{
					Name: "Foo",
					In: []model.Field{{
						Names: []string{"ctx"},
						Type: model.Ident{
							Path: "context",
							Name: "Context",
						},
					}},
					Out: []model.Field{
						{Type: model.Ident{Name: "error"}},
						{Type: model.Ident{Path: "model", Name: "SomethingElse"}},
					},
				}},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//decls := mapper.NewService(tt.props)
			//lo.Must0(service.WriteDecls(os.Stdout, "test", decls...))
			node := lo.Must(coding.Parse(`package test
			
type Something interface {}
			`))
			model2.Inspect(node, filter.OnInterface(assert.TypeMatch("Something", func(m *model2.Type, t *model2.Interface) {
				coding.Print(m)
			})))
		})
	}
}
