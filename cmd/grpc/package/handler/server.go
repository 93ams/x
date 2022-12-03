package handler

import (
	"github.com/tilau2328/cql/cmd/grpc/package/model"
	"google.golang.org/grpc"
)

func NewServer(ddl *DDL, dml *DML) *grpc.Server {
	s := grpc.NewServer()
	model.RegisterDDLServer(s, ddl)
	model.RegisterDMLServer(s, dml)
	return s
}
