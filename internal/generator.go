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
	if err := GenerateFile(path + DefaultConfigFile); err != nil {
		return err
	}
	f, err := os.OpenFile(path+DefaultConfigFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("error in opening file")
		return err
	}
	defer f.Close()
	if err := model.WriteDefaultConfig(f); err != nil {
		return err
	}
	return nil
}

func GenerateStructureUsingConfigFile(path, configFile string, addPath bool) error {
	if addPath {
		configFile = path + "/" + configFile
	}
	config, err := model.FetchConfig(configFile)
	if err != nil {
		return err
	}
	for _, f := range config.Folders {
		if addPath {
			f = path + "/" + f
		}
		GenerateFolder(f)
	}
	for _, f := range config.Files {
		if addPath {
			f = path + "/" + f
		}
		GenerateFile(f)
	}
	return nil
}

func GenerateModel(name, configFile string, fields []string) error {
	config, err := model.FetchConfig(configFile)
	if err != nil {
		return err
	}
	mErr := model.GenerateModelFile(name, config, fields)
	if mErr != nil {
		return mErr
	}
	return nil
}

func GenerateHandler(name, configFile string) error {
	config, err := model.FetchConfig(configFile)
	if err != nil {
		return err
	}
	hErr := model.GenerateHandlerFile(name, config)
	if hErr != nil {
		return hErr
	}
	return nil
}
