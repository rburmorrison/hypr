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

// NewOL creates a new Element preformatted as an
// unordered list and returns a reference to it.
func NewOL() Element {
	ol := NewTagElement("ol")
	return ol
}

// NewUL creates a new Element preformatted as an
// unordered list and returns a reference to it.
func NewUL() Element {
	ul := NewTagElement("ul")
	return ul
}

// NewLI creates a new Element preformatted as
// an item in an ordered or unordered list and
// returns a reference to it.
func NewLI(text string) Element {
	li := NewTagElement("li")
	li.AddChild(NewTextElement(text))

	return li
}
