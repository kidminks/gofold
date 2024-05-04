package internal

import (
	"log/slog"
	"os"

	"github.com/kidminks/gofold/internal/model"
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
	if err := model.WriteDefaultConfig(f); err != nil {
		return err
	}
	return nil
}
