package tplutil

import (
	"bytes"
)

func ExecuteTemplate(text string, model any) (string, error) {
	tpl, err := newTpl("").Parse(text)
	if err != nil {
		return "", err
	}
	buff := new(bytes.Buffer)
	if err := tpl.Execute(buff, model); err != nil {
		return "", err
	}
	return buff.String(), nil
}
