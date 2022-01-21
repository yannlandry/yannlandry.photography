package util

import (
	"path/filepath"
)

type Path struct {
	base string
}

func NewPath(base string) *Path {
	return &Path{
		base: base,
	}
}

func (this *Path) With(extension string) string {
	return filepath.Join(this.base, extension)
}
