package model

import (
	"log/slog"
	"os"
	"strings"
)

type Route struct {
	PackageName        string
	ModelPackageName   string
	ModelPackageImport string
	Name               string
}

func GenerateRouteFile(name string, config *Config) error {
	tp := strings.Split(config.Handler, "/")
	mp := strings.Split(config.Model, "/")
	s := buildHandler(&Handler{
		PackageName:        tp[len(tp)-1],
		ModelPackageName:   mp[len(mp)-1],
		Name:               name,
		ModelPackageImport: config.Module + "/" + config.Model,
	})
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
