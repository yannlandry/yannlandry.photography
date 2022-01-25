package util

import (
	"fmt"
	"net/url"
	"path"
)

var (
	BaseURL   *URLBuilder
	StaticURL *URLBuilder
)

type URLBuilder struct {
	base *url.URL
}

func NewURLBuilder(base string) (*URLBuilder, error) {
	parsed, err := url.Parse(base)
	if err != nil {
		return nil, fmt.Errorf("failed parsing URL `%s`: %s", base, err)
	}

	return &URLBuilder{
		base: parsed,
	}, nil
}

func (this *URLBuilder) With(extension string) string {
	// URLs that are already complete are not joined with `base`
	parsed, err := url.Parse(extension)
	if err != nil {
		return extension
	} else if parsed.IsAbs() {
		return parsed.String()
	}

	full := *this.base
	full.Path = path.Join(full.Path, parsed.Path)
	return full.String()
}

func (this *URLBuilder) Get() string {
	return this.base.String()
}
