package component

import (
	"context"
	"embed"

	"github.com/jcyamacho/rxgen/internal/tplutil"
)

var (
	//go:embed templates/*
	templates   embed.FS
	templateMap = map[Style]string{
		StyleCssMod:          "templates/css-module",
		StyleScssMod:         "templates/scss-module",
		StyleStyledComponent: "templates/styled-components",
		StyleEmotion:         "templates/emotion",
	}
)

func Create(ctx context.Context, destinationDir string, opt *Options) error {
	templateRoot, ok := templateMap[opt.Style]
	if !ok {
		return &InvalidStyleError{opt.Style}
	}

	return tplutil.ExecuteFS(ctx, templates, templateRoot, destinationDir, opt)
}
