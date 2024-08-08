package parser

type Type string

type GoStruct struct {
	Name   string
	Fields map[string]Type
}
