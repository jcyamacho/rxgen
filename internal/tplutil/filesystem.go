package tplutil

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/log"
)

func ExecuteFS(ctx context.Context, templateFS fs.FS, templateRoot string, destinationDir string, model any) error {
	logger := log.FromContext(ctx)

	return fs.WalkDir(templateFS, templateRoot, func(fullPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		trimmedPath := strings.TrimLeft(fullPath, templateRoot)
		path, err := ExecuteTemplate(trimmedPath, model)
		if err != nil {
			return err
		}
		if path == "" {
			return nil
		}
		finalPath := filepath.Join(destinationDir, path)

		if d.IsDir() {
			logger.Infof("creating directory if not exists: %s", path)
			return mkdirAll(finalPath)
		}

		logger.Infof("creating file: %s", path)

		fileTpl, err := newTplFromFS(templateFS, fullPath)
		if err != nil {
			return err
		}

		f, err := mkfile(finalPath)
		if err != nil {
			return err
		}
		defer f.Close()

		return fileTpl.Execute(f, model)
	})
}

func mkdirAll(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return err
}

func mkfile(path string) (*os.File, error) {
	if _, err := os.Stat(path); err == nil {
		return nil, &FileAlreadyExistsError{path}
	}
	return os.Create(path)
}
