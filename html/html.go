// Package html defines types and functions to
// programmatically generate HTML documents.
package html

import (
	"fmt"
	"io"
)

type attrs map[string]string

func (a attrs) write(w io.Writer) {
	for k, v := range a {
		fmt.Fprintf(w, " %s=%q", k, v)
	}
}

// HTML represents an HTML element.
type HTML struct {
	Lang string

	Head Head
	Body Body
}

func (h *HTML) Write(w io.Writer) {
	fmt.Fprint(w, `<html`)
	if h.Lang != "" {
		fmt.Fprintf(w, ` lang="%s"`, h.Lang)
	}
	fmt.Fprint(w, `>`)
	defer fmt.Fprint(w, `</html>`)

	// Head
	if len(h.Head.Children) > 0 {
		h.Head.Write(w)
	}

	// Body
	if len(h.Body.Children) > 0 {
		h.Body.Write(w)
	}
}
