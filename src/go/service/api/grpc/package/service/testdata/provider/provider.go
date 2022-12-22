package provider

import (
	"context"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/service/testdata/model"
)

type Provider interface {
	SayHello(context.Context, model.HelloRequest) (*model.HelloReply, error)
}
