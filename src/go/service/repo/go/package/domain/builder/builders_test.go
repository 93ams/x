package builder_test

import (
	"github.com/tilau2328/cql/src/go/services/gen/go/package/adaptor/driver/resolver"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/adaptor/driver/restorer"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder/model/decl"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder/model/expr"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder/model/node"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder/model/spec"
	. "github.com/tilau2328/cql/src/go/services/gen/go/package/domain/builder/model/stmt"
	"github.com/tilau2328/cql/src/go/services/gen/go/package/domain/model"
	"go/token"
	"strconv"
	"testing"
)

func TestMethodBuilder_Build(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := restorer.NewRestorerWithImports("root", resolver.NewGuessResolver())
			r.Print(File(Ident("asd"), Import(Ident("fmt"))).Decls(
				Structure(),
				Inter(),
				Mapper(MapperProps{Name: "ToTest", From: ""}),
			).Build())
		})
	}
}

type MapperProps struct {
	Name string
	From string
	To   string
}

func Inter() builder.GenBuilder {
	return Gen(token.TYPE, Type(Interface(FieldList())).Name(Ident("Interface")))
}
func Structure() builder.GenBuilder {
	return Gen(token.TYPE, Type(Struct(FieldList(Field(Ident("string"), Ident("Bar"))))).Name(Ident("Struct")))
}
func Mapper(props MapperProps) builder.FuncBuilder {
	return Func(Ident(props.Name)).
		Type(FuncType().
			Params(FieldList(Field(Ident("Struct").Path("model"), Ident("in")))).
			Results(FieldList(Field(Ident("Struct"))))).
		Body(Block(
			Expr(Call(Ident("Println").Path("fmt"),
				Lit(token.STRING, strconv.Quote("Hello World!")))),
			Return(CompositeLit(Ident("Struct"),
				KeyValue(Ident("Foo"), Selector(Ident("in"), Ident("Bar"))).Decs(
					KeyValueDecs().Before(model.NewLine).After(model.NewLine),
				),
			)).Decs(ReturnDecs().Before(model.NewLine)),
		))
}
