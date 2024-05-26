package model

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kidminks/gofold/template"
)

func WriteDbFile(module string, f *os.File) error {
	db := template.GetDbTemplate()
	if _, err := fmt.Fprintln(f, db); err != nil {
		slog.Error("error in writing default json to file", "error", err)
		return err
	}

	if err := f.Close(); err != nil {
		slog.Error("error in closing file", "error", err)
		return err
	}
	return nil
}
