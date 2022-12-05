package gql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
)

type (
	Endpoint string
	Server   struct {
		ep Endpoint
		*handler.Server
	}
)

func NewServer(s *handler.Server, ep Endpoint) *Server {
	return &Server{Server: s, ep: ep}
}
func (s *Server) Playground() http.HandlerFunc {
	return playground.Handler("GraphQL playground", string(s.ep))
}
func (s *Server) WithPlayground(path string) *Server {
	http.Handle(path, s.Playground())
	return s
}
func (s *Server) Handle() { http.Handle(string(s.ep), s.Server) }
func (s *Server) Start() error {
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "7878")
	return http.ListenAndServe(":7878", nil)
}
