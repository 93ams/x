// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type KeySpace struct {
	Name   string   `json:"name"`
	Tables []*Table `json:"tables"`
}

type NewKeyspace struct {
	Name string `json:"name"`
}

type NewTable struct {
	Name   string    `json:"name"`
	Params []*string `json:"params"`
}

type Table struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}