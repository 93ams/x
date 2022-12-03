package main

import (
	"github.com/samber/lo"
	"github.com/tilau2328/cql/package/shared/data/cql"
)

//go:generate protoc -I schema --go_out=package --go-grpc_out=package schema/ddl.proto schema/dml.proto
func main() {
	_, close := lo.Must2(Init(cql.Options{
		Keyspace: "",
	}))
	defer close()

}
