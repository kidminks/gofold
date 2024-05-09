package model

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"unicode"

	"github.com/kidminks/gofold/template"
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

type Query struct {
	InsertQuery string
}

func GenerateModelFile(name string, config *Config, fields []string) error {
	tp := strings.Split(config.Model, "/")
	p := tp[len(tp)-1]
	s := buildModel(p, name, fields)
	f, err := os.OpenFile(config.Model, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		slog.Error("error in opening file", "error", err)
		return err
	}
	_, fWriteError := f.WriteString(s)
	if fWriteError != nil {
		slog.Error("error in closing file", "error", fWriteError)
		return fWriteError
	}
	return nil
}

func buildModel(packageName, name string, fields []string) string {
	ff, fs := buildFieldStructure(fields)
	fmt.Println(ff)
	s := template.GetModelTemplate()
	s = strings.ReplaceAll(s, "{package}", packageName)
	s = strings.ReplaceAll(s, "{model_name}", name)
	s = strings.ReplaceAll(s, "{fields}", fs)
	r := []rune(name)
	r[0] = unicode.ToLower(r[0])
	cn := string(r)
	s = strings.ReplaceAll(s, "{model_name_camel}", cn)
	iq, iField := buildInsertQuery(name, ff)
	s = strings.ReplaceAll(s, "{insert_query}", iq)
	s = strings.ReplaceAll(s, "{insert_exec}", iField)
	uq, uField := buildUpdateQuery(name, ff)
	s = strings.ReplaceAll(s, "{update_query}", uq)
	s = strings.ReplaceAll(s, "{update_exec}", uField)
	return s
}

func buildFieldStructure(fields []string) ([]Field, string) {
	ff := []Field{}
	fs := ""
	for _, f := range fields {
		s := strings.Split(f, ":")
		ff = append(ff, Field{
			Key:  s[0],
			Type: s[1],
		})
		fs += s[0] + " " + s[1] + "\n"
	}
	return ff, fs
}

func buildInsertQuery(name string, fields []Field) (string, string) {
	q := "INSERT INTO {model_name} ({field}) VALUES ({marks})"
	iField, iMarks := "", ""
	for _, f := range fields {
		iField += f.Key + ","
		iMarks += "?,"
	}
	iField = iField[:len(iField)-1]
	iMarks = iMarks[:len(iMarks)-1]
	q = strings.ReplaceAll(q, "{model_name}", name)
	q = strings.ReplaceAll(q, "{field}", iField)
	q = strings.ReplaceAll(q, "{marks}", iMarks)
	return q, iField
}

func buildUpdateQuery(name string, fields []Field) (string, string) {
	q := "UPDATE {model_name} SET {field} WHERE id = ?"
	iField, iMarks := "id,", ""
	for _, f := range fields {
		iField += f.Key + ","
		iMarks += f.Key + " = ?,"
	}
	iField = iField[:len(iField)-1]
	iMarks = iMarks[:len(iMarks)-1]
	q = strings.ReplaceAll(q, "{model_name}", name)
	q = strings.ReplaceAll(q, "{name}", name)
	q = strings.ReplaceAll(q, "{field}", iMarks)
	return q, iField
}
