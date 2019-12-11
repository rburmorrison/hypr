package html

import (
	"fmt"
	"io"
)

// Document represents an HTML document.
type Document struct {
	Children []Element
}

// NewDocument creates a new Document and returns a reference to it.
func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Write(w io.Writer) {
	fmt.Fprint(w, `<!DOCTYPE html>`)
	for _, child := range d.Children {
		child.Write(w)
	}
}

// Add appends the given children to the document's children.
func (d *Document) Add(children ...Element) {
	d.Children = append(d.Children, children...)
}
