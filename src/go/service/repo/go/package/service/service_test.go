package service

import (
	"github.com/tilau2328/x/src/go/services/repo/go/package/domain/model"
	"log"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		req  CreateReq
	}{
		{
			req: CreateReq{
				Pkg:  "test",
				File: "./mapper.go",
				Props: model.Mapper{
					From: model.Struct{
						Path: "model",
						Name: "Struct",
						Fields: []model.StructField{
							{Names: []string{"Foo"}},
							{Names: []string{"Bar"}},
						},
					},
					To: model.Struct{Name: "Dest"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println(NewService().Create(tt.req))
		})
	}
}
