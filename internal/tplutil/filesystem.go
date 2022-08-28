package tplutil

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ExecuteFS(templateFS fs.FS, templateRoot string, destinationDir string, model any) error {
	return fs.WalkDir(templateFS, templateRoot, func(fullPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		trimmedPath := strings.TrimLeft(fullPath, templateRoot)
		path, err := ExecuteTemplate(trimmedPath, model)
		if err != nil {
			return err
		}
		finalPath := filepath.Join(destinationDir, path)

		if d.IsDir() {
			return mkdirAll(finalPath)
		}

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
