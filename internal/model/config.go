package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Config struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
	Config  string   `json:"config"`
	Model   string   `json:"model"`
	Main    string   `json:"main"`
}

func WriteDefaultConfig(f *os.File) error {
	configJson := `
{
	"folders": ["cmd/server", "internal/model", "internal/db", "config"],
	"file": ["cmd/server/main.go", ".gitignore", "go.mod"],
	"config": "config",
	"model": "internal/model",
	"handler": "internal/handler"
	"main": "/cmd/server/main.go"
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

func FetchConfig(configFile string) (*Config, error) {
	jsonFile, err := os.Open(configFile)
	if err != nil {
		slog.Error("error in opening file", "file", configFile)
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var result Config
	json.Unmarshal([]byte(byteValue), &result)
	return &result, nil
}
