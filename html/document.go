package html

import (
	"fmt"
	"io"
)

// Document represents an HTML document.
type Document struct {
	HTML HTML
}

// NewDocument creates a new Document and returns a reference to it.
func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Write(w io.Writer) {
	fmt.Fprint(w, `<!DOCTYPE html>`)
	d.HTML.Write(w)
}
