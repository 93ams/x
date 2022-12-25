package service_test

import (
	"github.com/tilau2328/x/src/go/service/api/grpc/package/pattern"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	golang "github.com/tilau2328/x/src/go/service/repo/go/package/services"
	"testing"
)

import (
	"github.com/tilau2328/x/src/go/service/api/grpc/package/service"
)

func TestService_Generate(t *testing.T) {
	tests := []struct {
		name    string
		props   pattern.AdaptorProps
		wantErr bool
	}{
		{props: pattern.AdaptorProps{
			Provider: model.SearchReq{
				File: "testdata/provider/provider",
				Name: "Provider",
			},
			Models:    model.SearchReq{File: "testdata/model/models"},
			Grpc:      model.SearchReq{File: "testdata/test_grpc.pb"},
			Proto:     model.SearchReq{File: "testdata/test.pb"},
			Mappers:   "testdata/mapper/mappers",
			Handler:   "testdata/adaptor/server/handler",
			Requester: "testdata/adaptor/client/requester",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service.NewService(service.Options{
				Golang: golang.NewService(),
			})
			if err := s.Generate(tt.props); (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
