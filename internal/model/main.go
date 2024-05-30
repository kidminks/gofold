package model

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kidminks/gofold/template"
)

func WriteMainFile(module string, f *os.File) error {
	main := template.GetMainTemplate()
	main.Module = module
	main.CreateValueMap()
	main.Replace()
	if _, err := fmt.Fprintln(f, main.Output); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}

	return nil
}
