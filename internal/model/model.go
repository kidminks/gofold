package model

import (
	"fmt"
	"strings"
)

type Field struct {
	Key  string
	Type string
}

type Model struct {
	PackageName string
	Name        string
	Fields      []Field
}

func BuildModel(packageName, name string, fields []string) {
	ff := buildFieldStructure(fields)
	fmt.Println(ff)
}

func buildFieldStructure(fields []string) []Field {
	ff := []Field{}
	for _, f := range fields {
		s := strings.Split(f, ":")
		ff = append(ff, Field{
			Key:  s[0],
			Type: s[1],
		})
	}
	return ff
}
