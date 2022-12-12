package main

//go:generate swagger-cli bundle config/core.yaml --outfile config/schema.yaml --type yaml
//go:generate oapi-codegen -package model -generate types -o ./package/model/models_gen.go config/schema.yaml
//go:generate oapi-codegen -package request -generate client -o ./package/request/generated.go config/schema.yaml
//go:generate oapi-codegen -package handler -generate chi-server -o ./package/handler/generated.go config/schema.yaml

func main() {

}
