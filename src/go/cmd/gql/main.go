package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tilau2328/cql/src/go/cmd/gql/package/exec"
	"github.com/tilau2328/cql/src/go/cmd/gql/package/resolver"
	"log"
	"net/http"
)

//go:generate gqlgen generate --config gencfg.yaml

func main() {
	srv := handler.NewDefaultServer(exec.NewExecutableSchema(exec.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "7878")
	log.Fatal(http.ListenAndServe(":7878", nil))
}
