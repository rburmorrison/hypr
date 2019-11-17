package html

import (
	"io"
)

// Element represents any HTML element.
type Element interface {
	Write(w io.Writer)
	AddAttr(k, v string)
	AddChild(Element)
}

// NewLink creates a new Element preformatted
// as an anchor and returns a reference to it.
func NewLink(link, text string) Element {
	a := NewTagElement("a")
	a.AddChild(NewTextElement(text))
	a.AddAttr("href", link)

	return a
}
