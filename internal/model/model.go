package model

import (
	"log/slog"
	"os"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
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
	mFileName := config.Model + "/" + strings.ToLower(name) + ".go"
	f, err := os.OpenFile(mFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("error in opening file", "error", err)
		return err
	}
	defer f.Close()
	_, fWriteError := f.WriteString(s)
	if fWriteError != nil {
		slog.Error("error in closing file", "error", fWriteError)
		return fWriteError
	}
	return nil
}

func buildModel(packageName, name string, fields []string) string {
	ff, fs := buildFieldStructure(fields)
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
	rq := rowScanQuery(name, ff)
	s = strings.ReplaceAll(s, "{fetch_row_scan}", rq)
	return s
}

func buildFieldStructure(fields []string) ([]Field, string) {
	ff := []Field{}
	fs := ""
	for _, f := range fields {
		s := strings.Split(f, ":")
		ff = append(ff, Field{
			Key:  strcase.ToCamel(s[0]),
			Type: s[1],
		})
		fs += strcase.ToCamel(s[0]) + " " + s[1] + "\n"
	}
	return ff, fs
}

func buildInsertQuery(name string, fields []Field) (string, string) {
	q := `"INSERT INTO {model_name} ({field}) VALUES ({marks})"`
	iField, iMarks, iExec := "", "", ""
	for _, f := range fields {
		iField += strcase.ToSnake(f.Key) + ","
		iMarks += "?,"
		iExec += strings.ToLower(name) + "." + f.Key + ","
	}
	iField = iField[:len(iField)-1]
	iExec = iExec[:len(iExec)-1]
	iMarks = iMarks[:len(iMarks)-1]
	q = strings.ReplaceAll(q, "{model_name}", name)
	q = strings.ReplaceAll(q, "{field}", iField)
	q = strings.ReplaceAll(q, "{marks}", iMarks)
	return q, iExec
}

func buildUpdateQuery(name string, fields []Field) (string, string) {
	q := `"UPDATE {model_name} SET {field} WHERE id = ?"`
	iMarks, uExec := "", ""
	for _, f := range fields {
		iMarks += strcase.ToSnake(f.Key) + " = ?,"
		uExec += strings.ToLower(name) + "." + f.Key + ","
	}
	uExec = uExec[:len(uExec)-1]
	iMarks = iMarks[:len(iMarks)-1]
	q = strings.ReplaceAll(q, "{model_name}", name)
	q = strings.ReplaceAll(q, "{field}", iMarks)
	return q, uExec
}

func rowScanQuery(name string, fields []Field) string {
	rScan := ""
	for _, f := range fields {
		rScan += "&" + strings.ToLower(name) + "." + f.Key + ","
	}
	rScan = rScan[:len(rScan)-1]
	return rScan
}
