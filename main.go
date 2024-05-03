package main

import (
	"log/slog"

	"os"
)

func generateDefaultFolders(projectName string) error {
	internal := projectName + "/internal"
	cmd := projectName + "/cmd"
	config := projectName + "/config"

	paths := []string{
		internal,
		internal + "/model",
		internal + "/db",
		cmd,
		cmd + "/server",
		config,
	}

	for _, path := range paths {
		if err := generateFolder(path); err != nil {
			return err
		}

	}
	return nil

}

func generateFolder(path string) error {
	slog.Info("creating folder ", "path", path, "status", "started")
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	slog.Info("creating folder ", "path", path, "status", "completed")
	return nil

}

func generateDefaultFile(projectName string) error {
	files := []string{
		projectName + "/cmd/server/main.go",
		projectName + "/.gitignore",
		projectName + "/go.mod",
	}

	for _, file := range files {
		if err := generateFile(file); err != nil {
			return err
		}
	}
	return nil

}

func generateFile(file string) error {
	slog.Info("creating file ", "path", file, "status", "started")
	if _, err := os.Create(file); err != nil {
		return err
	}
	slog.Info("creating folder ", "path", file, "status", "completed")
	return nil
}

func main() {
	projectName := ""
	if len(os.Args) >= 2 {
		projectName = os.Args[1]
	} else {
		slog.Error("project name not given")
		return
	}
	if err := generateDefaultFolders(projectName); err != nil {
		slog.Error("error in folder generation", "error", err)
	}
	if err := generateDefaultFile(projectName); err != nil {
		slog.Error("error in file generation", "error", err)
	}

}
