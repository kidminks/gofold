package model

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kidminks/gofold/template"
)

func WriteGoMod(module string, f *os.File) error {
	goMod := template.GetGoModTemplate()
	if _, err := fmt.Fprintln(f, goMod); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}

	return nil
}
