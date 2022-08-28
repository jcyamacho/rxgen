package component

type Style string

const (
	StyleCssMod          Style = "cssmod"
	StyleScssMod         Style = "scssmod"
	StyleStyledComponent Style = "styled"
	StyleEmotion         Style = "emotion"
)

var Styles = []Style{StyleCssMod, StyleScssMod, StyleStyledComponent, StyleEmotion}

type Options struct {
	Name          string
	Style         Style
	StyledPostfix string
}
