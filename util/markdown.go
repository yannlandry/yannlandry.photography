package util

import (
	mdcore "github.com/gomarkdown/markdown"
        mdparser "github.com/gomarkdown/markdown/parser"
        mdhtml "github.com/gomarkdown/markdown/html"
)

var (
	Markdown *MarkdownEngine
)

type MarkdownEngine struct {
	extensions mdparser.Extensions
	renderer *mdhtml.Renderer
	baseURL *URLBuilder
	staticURL *URLBuilder
}

func NewMarkdownEngine(baseURL *URLBuilder, staticURL *URLBuilder) *MarkdownEngine {
	return &MarkdownEngine{
		extensions: mdparser.CommonExtensions | mdparser.AutoHeadingIDs,
		renderer: mdhtml.NewRenderer(mdhtml.RendererOptions{}),
		baseURL: baseURL,
		staticURL: staticURL,
	}
}

func (this *MarkdownEngine) Render(content []byte) []byte {
	parser := mdparser.NewWithExtensions(this.extensions)
	return mdcore.ToHTML(content, parser, this.renderer)
}
