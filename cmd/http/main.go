package main

//go:generate oapi-codegen -package model -generate types -o ./package/model/models_gen.go schema.yaml
//go:generate oapi-codegen -package model -generate client -o ./package/request/generated.go schema.yaml
//go:generate oapi-codegen -package model -generate chi-server -o ./package/handler/generated.go schema.yaml

func main() {

}
