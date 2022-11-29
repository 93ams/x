package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gql/package/exec"
	"gql/package/resolver"
	"log"
	"net/http"
)

func main() {
	srv := handler.NewDefaultServer(exec.NewExecutableSchema(exec.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "7878")
	log.Fatal(http.ListenAndServe(":7878", nil))
}
