package handler

import (
	"github.com/tilau2328/cql/src/go/cmd/grpc/package/model"
	"google.golang.org/grpc"
	"net/http"
)

type Server struct {
	*grpc.Server
}

func NewServer(ddl *DDLHandler, dml *DML) Server {
	s := grpc.NewServer()
	model.RegisterDDLServer(s, ddl)
	model.RegisterDMLServer(s, dml)
	return Server{Server: s}
}

func (s Server) Serve() error {
	return http.ListenAndServe(":9099", s)
}
