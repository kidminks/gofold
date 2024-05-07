package model

type Field struct {
	Key  string
	Type string
}

type Model struct {
	PackageName string
	Name        string
	Fields      []Field
}
