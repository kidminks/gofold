package model

import (
	"log/slog"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kidminks/gofold/template"
)

type Handler struct {
	PackageName string
	Name        string
}

func GenerateHandlerFile(name string, config *Config) error {
	tp := strings.Split(config.Handler, "/")
	mp := strings.Split(config.Model, "/")
	s := buildHandler(tp[len(tp)-1], mp[len(mp)-1], name)
	mFileName := config.Handler + "/" + strings.ToLower(name) + ".go"
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

func buildHandler(packageName, modelPackageName, name string) string {
	s := template.GetHandlerTemplate()
	s = strings.ReplaceAll(s, "{package}", packageName)
	s = strings.ReplaceAll(s, "{model_name}", name)
	nameCamel := strcase.ToCamel(name)
	s = strings.ReplaceAll(s, "{model_name_camel}", nameCamel)
	s = strings.ReplaceAll(s, "{model_package}", modelPackageName)
	return s
}
