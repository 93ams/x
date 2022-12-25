package model_test

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"testing"
)

func TestMapper(t *testing.T) {
	tests := []struct {
		name   string
		fields model.Mapper
		want   string
	}{
		{
			fields: model.Mapper{
				Name: "SomethingToProto",
				From: &model.Ident{
					Name: "Something",
					Path: "model",
				},
				To: &model.Ident{
					Name: "Something",
					Path: "adaptor",
				},
				Assignments: []*model.KeyValue{
					{
						Value: &model.Ident{Name: "Foo"},
						Key:   &model.Ident{Name: "Foo"},
					},
					{
						Value: &model.Ident{Name: "Bar"},
						Key:   &model.Ident{Name: "Bar2"},
					},
				},
			},
			want: `package mapper

import (
	"adaptor"
	"model"
)

func SomethingToProto(in model.Something) adaptor.Something {
	return adaptor.Something{
		Foo:  in.Foo,
		Bar2: in.Bar,
	}
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			require.NoError(t, driver.Write(buf, &model.File{
				Name:  &model.Ident{Name: "mapper"},
				Decls: []model.Decl{tt.fields.Func()},
			}))
			require.Equal(t, tt.want, buf.String())
		})
	}
}

func TestBuilder(t *testing.T) {
	tests := []struct {
		name   string
		fields model.Builder
		want   string
	}{
		{
			fields: model.Builder{
				Name: "SomethingBuilder",
				Type: &model.Ident{
					Name: "Something",
					Path: "model",
				},
				Fields:  nil,
				Methods: nil,
			},
			want: `package builder

import "model"

type SomethingBuilder x.Builder[model.Something]

func NewSomethingBuilder() *model.Something {
	return &model.Something{}
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			require.NoError(t, driver.Write(buf, &model.File{
				Name:  &model.Ident{Name: "builder"},
				Decls: tt.fields.Decls(),
			}))
			require.Equal(t, tt.want, buf.String())
		})
	}
}

func TestService(t *testing.T) {
	tests := []struct {
		name   string
		fields model.Service
		want   string
	}{
		{
			fields: model.Service{
				Name: "SomethingService",
				Methods: []*model.Func{
					{
						Name: &model.Ident{Name: "Foo"},
						Type: &model.FuncType{
							Params:  &model.FieldList{},
							Results: &model.FieldList{},
						},
						Body: &model.Block{},
					},
				},
			},
			want: `package builder

type SomethingService struct {
}

func NewSomethingService() *SomethingService {
	return &SomethingService{}
}
func (*SomethingService) Foo() {}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			require.NoError(t, driver.Write(buf, &model.File{
				Name:  &model.Ident{Name: "builder"},
				Decls: tt.fields.Decls(),
			}))
			//coding.Print(&model.File{
			//	Name:  &model.Ident{Name: "builder"},
			//	Decls: tt.fields.Decls(),
			//})
			require.Equal(t, tt.want, buf.String())
		})
	}
}
