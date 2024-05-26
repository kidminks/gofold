package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/kidminks/gofold/template"
)

type Config struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
	Config  string   `json:"config"`
	Model   string   `json:"model"`
	Handler string   `json:"handler"`
	Route   string   `json:"route"`
	Main    string   `json:"main"`
	Db      string   `json:"db"`
	Module  string   `json:"module"`
}

func WriteDefaultConfig(module string, f *os.File) error {
	configJson := template.GetDefaultConfigTemplate()
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
