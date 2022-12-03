package handler

import (
	"google.golang.org/grpc"
	"grpc/package/model"
)

func NewServer(ddl DDL, dml DML) *grpc.Server {
	s := grpc.NewServer()
	model.RegisterDDLServer(s, ddl)
	model.RegisterDMLServer(s, dml)
	return s
}
