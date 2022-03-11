package models

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Site struct {
	Model
	Name     string
	Content  string
	ParentID *uint
	Parent   *Site
}

var p = parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs)
var r = html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank})

func (s Site) String() string {
	return string(markdown.ToHTML([]byte(s.Content), p, r))
}
