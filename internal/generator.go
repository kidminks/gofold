package internal

import (
	"fmt"
	"log/slog"
	"os"
)

func GenerateFolder(path string) error {
	slog.Info("creating folder ", "path", path, "status", "started")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	slog.Info("creating folder ", "path", path, "status", "completed")
	return nil
}

func GenerateFile(file string) error {
	slog.Info("creating file ", "path", file, "status", "started")
	if _, err := os.Create(file); err != nil {
		return err
	}
	slog.Info("creating folder ", "path", file, "status", "completed")
	return nil
}

func GenerateDefaultConfigFile(path string) error {
	if err := GenerateFolder(path); err != nil {
		return err
	}
	if err := GenerateFile(path + "/gofold_config.json"); err != nil {
		return err
	}
	f, err := os.OpenFile(path+"/gofold_config.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("error in opening file")
		return err
	}

	configJson := `
{
	"folders": ["cmd/server", "internal/model", "internal/db", "config"],
	"file": ["cmd/server/main.go", ".gitignore", "go.mod"],
	"config": "config"
	"model": "internal/model",
	"main": "/cmd/server/main.go",
}
	`
	if _, err := fmt.Fprintln(f, configJson); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}
	return nil
}
