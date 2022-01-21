package util

import (
	"io"
	"strings"

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
		RenderNodeHook: processNode,
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

func processNode(writer io.Writer, node mdast.Node, entering bool) (mdast.WalkStatus, bool) {
	if h, ok := node.(*mdast.Heading); ok && entering {
		processHeading(h)
	}
	if p, ok := node.(*mdast.Paragraph); ok && entering {
		processParagraph(p)
	}
	if a, ok := node.(*mdast.Link); ok && entering {
		processLink(a)
	}
	if img, ok := node.(*mdast.Image); ok && entering {
		processImage(img)
	}

	return mdast.GoToNext, false
}

func processHeading(h *mdast.Heading) {
	// Add permalinks to all headings
	permalink := &mdast.Link{}
	permalink.Destination = []byte("#" + h.HeadingID)
	permalink.AdditionalAttributes = []string{"class=\"permalink\""}
	children := h.GetChildren()
	children = append(children, permalink)
	h.SetChildren(children)
}

func processParagraph(p *mdast.Paragraph) {
	// Add optional classes to paragraphs
	children := p.GetChildren()

	// The paragraph must have at least two children
	if len(children) < 2 {
		return
	}

	// The first child must be an empty text node
	if text, ok := children[0].(*mdast.Text); !ok || strings.TrimSpace(string(text.Leaf.Literal)) != "" {
		return
	}

	// The second child must be a code node
	code, ok := children[1].(*mdast.Code)
	if !ok {
		return
	}

	// The code node must start with `class:`
	content := strings.TrimSpace(string(code.Leaf.Literal))
	if !strings.HasPrefix(content, "class:") {
		return
	}

	// Add special class to paragraph
	addAttribute(p)
	p.Attribute.Classes = append(p.Attribute.Classes, []byte(content[6:]))

	// Remove code tag
	p.SetChildren(children[2:])
}

func processLink(a *mdast.Link) {
	// Prepend all local URLs with the base URL, except anchor links
	if len(a.Destination) > 0 && a.Destination[0] != '#' {
		a.Destination = []byte(BaseURL.With(string(a.Destination)))
	}
}

func processImage(img *mdast.Image) {
	// Prepend all images with the static URL
	img.Destination = []byte(StaticURL.With(string(img.Destination)))
}

func addAttribute(node mdast.Node) {
	// Adds an *Attribute to a Container/Leaf if not present
	if container := node.AsContainer(); container != nil {
		if container.Attribute == nil {
			container.Attribute = &mdast.Attribute{
				ID: []byte{},
				Classes: [][]byte{},
				Attrs: map[string][]byte{},
			}
		}
	} else if leaf := node.AsLeaf(); leaf != nil {
		// NOTE: the renderer ignores `*Attribute` on `Leaf` instances, for whatever reason
		if leaf.Attribute == nil {
			leaf.Attribute = &mdast.Attribute{
				ID: []byte{},
				Classes: [][]byte{},
				Attrs: map[string][]byte{},
			}
		}
	}
}
