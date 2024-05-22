package model

import (
	"log/slog"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/kidminks/gofold/template"
)

type Route struct {
	PackageName        string
	ModelPackageName   string
	ModelPackageImport string
	Name               string
}

func GenerateRouteFile(name string, config *Config) error {
	tp := strings.Split(config.Route, "/")
	mp := strings.Split(config.Model, "/")
	s := buildHRoute(&Route{
		PackageName:        tp[len(tp)-1],
		ModelPackageName:   mp[len(mp)-1],
		Name:               name,
		ModelPackageImport: config.Module + "/" + config.Model,
	})
	mFileName := config.Route + "/" + strings.ToLower(name) + ".go"
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

func buildHRoute(h *Route) string {
	s := template.GetRouteTemplate()
	s = strings.ReplaceAll(s, "{package}", h.PackageName)
	s = strings.ReplaceAll(s, "{model_name}", h.Name)
	nameCamel := strings.ToLower(strcase.ToCamel(h.Name))
	s = strings.ReplaceAll(s, "{model_name_camel}", nameCamel)
	s = strings.ReplaceAll(s, "{model_package}", h.ModelPackageName)
	s = strings.ReplaceAll(s, "{model_package_import}", h.ModelPackageImport)
	return s
}
