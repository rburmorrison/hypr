package html

import (
	"fmt"
	"io"
)

// TextElement represents plain text in HTML.
type TextElement string

// NewTextElement creates a new Element and returns a reference to it.
func NewTextElement(text string) *TextElement {
	te := TextElement(text)
	return &te
}

func (e *TextElement) Write(w io.Writer) {
	fmt.Fprint(w, string(*e))
}

// SetAttr is only implemented to make TextElement an Element.
func (e *TextElement) SetAttr(k, v string) {}

// Add is only implemented to make TextElement an Element.
func (e *TextElement) Add(...Element) {}
