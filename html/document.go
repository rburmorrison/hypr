package html

import (
	"fmt"
	"io"
)

// Document represents an HTML document.
type Document struct {
	Lang string
	Head Head
	Body Body

	Pretty bool
}

// NewDocument creates a new Document and returns a reference to it.
func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Write(w io.Writer) {
	// DOCTYPE
	fmt.Fprint(w, `<!DOCTYPE html>`)

	// HTML
	fmt.Fprint(w, `<html`)
	if d.Lang != "" {
		fmt.Fprintf(w, ` lang="%s"`, d.Lang)
	}
	fmt.Fprint(w, `>`)
	defer fmt.Fprint(w, `</html>`)

	// Head
	if len(d.Head.Elements) > 0 {
		d.Head.Write(w)
	}

	// Body
	if len(d.Body.Elements) > 0 {
		d.Body.Write(w)
	}
}
