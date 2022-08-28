package tplutil

import "fmt"

type FileAlreadyExistsError struct {
	Path string
}

func (f *FileAlreadyExistsError) Error() string {
	return fmt.Sprintf("file already exists: %s", f.Path)
}
