package util

import (
	"io"

	mdast "github.com/gomarkdown/markdown/ast"
	mdcore "github.com/gomarkdown/markdown"
        mdhtml "github.com/gomarkdown/markdown/html"
        mdparser "github.com/gomarkdown/markdown/parser"
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
	options := mdhtml.RendererOptions{
		RenderNodeHook: enhance,
	}

	return &MarkdownEngine{
		extensions: mdparser.CommonExtensions | mdparser.AutoHeadingIDs,
		renderer: mdhtml.NewRenderer(options),
		baseURL: baseURL,
		staticURL: staticURL,
	}
}

func (this *MarkdownEngine) Render(content []byte) []byte {
	parser := mdparser.NewWithExtensions(this.extensions)
	return mdcore.ToHTML(content, parser, this.renderer)
}

func enhance(writer io.Writer, node mdast.Node, entering bool) (mdast.WalkStatus, bool) {
	if h, ok := node.(*mdast.Heading); ok && entering {
		// Add permalinks to all headings
		permalink := &mdast.Link{}
		permalink.Destination = []byte("#" + h.HeadingID)
		permalink.AdditionalAttributes = []string{"class=\"permalink\""}
		children := h.GetChildren()
		children = append(children, permalink)
		h.SetChildren(children)
	}
	if a, ok := node.(*mdast.Link); ok && entering {
		// Prepend all local URLs with the base URL, except anchor links
		if len(a.Destination) > 0 && a.Destination[0] != '#' {
			a.Destination = []byte(BaseURL.With(string(a.Destination)))
		}
	}
	if img, ok := node.(*mdast.Image); ok && entering {
		// Prepend all images with the static URL
		img.Destination = []byte(StaticURL.With(string(img.Destination)))
	}

	return mdast.GoToNext, false
}
