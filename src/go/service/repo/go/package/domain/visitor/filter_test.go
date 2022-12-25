package visitor

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"testing"
)

type TypeType interface {
	int | float64
}

func TestOn(t *testing.T) {
	type args struct {
		fn  func(model.Node) bool
		ret any
	}
	tests := []struct {
		name string
		args args
		want func(model.Node) bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := make([]*model.Type, 0)
			filter := On(func(t *model.Type) bool {
				ret = append(ret, t)
				return true
			})
			model.Inspect(lo.Must(driver.Parse(`package test
type (
	Stuff interface {
		
	}
	Something struct {

	}
)
`)), filter)
			driver.Print(ret)
		})
	}
}
func TestOnType(t *testing.T) {
	type args struct {
		fn  func(model.Node) bool
		ret any
	}
	tests := []struct {
		name string
		args args
		want func(model.Node) bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := make([]*model.Type, 0)
			filter := OnType(func(t *model.Type, node model.Expr) bool {
				ret = append(ret, t)
				return true
			})
			model.Inspect(lo.Must(driver.Parse(`package test
type (
	Stuff interface {
		
	}
	Something struct {

	}
)
`)), filter)
			driver.Print(ret)
		})
	}
}
