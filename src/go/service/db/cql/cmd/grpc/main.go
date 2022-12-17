package main

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/shared/data/cql"
)

//go:generate protoc -I config/schema --go_out=package --go-grpc_out=package config/schema/ddl.proto config/schema/dml.proto

func main() {
	server, close := lo.Must2(Init(cql.Options{}))
	defer close()
	server.Serve()
}
