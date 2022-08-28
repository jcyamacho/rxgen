package tplutil

import (
	"io/fs"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
)

var funcMap = template.FuncMap{
	"ToCamelCase":  strcase.ToLowerCamel,
	"ToPascalCase": strcase.ToCamel,
	"ToKebabCase":  strcase.ToKebab,
	"ToSnakeCase":  strcase.ToSnake,
}

func newTpl(name string) *template.Template {
	return template.New(name).Funcs(funcMap)
}

func newTplFromFS(tplFS fs.FS, path string) (*template.Template, error) {
	return newTpl(filepath.Base(path)).ParseFS(tplFS, path)
}
