package html

import (
	"fmt"
	"io"
)

// Element represents a generic HTML element.
type Element struct {
	Tag      string
	Text     string
	Children []Element

	attrs map[string]string
}

// NewElement creates a new Element and returns a
// reference to it.
func NewElement() *Element {
	var el *Element
	el.attrs = make(map[string]string)

	return el
}

func (e *Element) Write(w io.Writer) {
	fmt.Fprintf(w, "<%s", e.Tag)
	for k, v := range e.attrs {
		fmt.Fprintf(w, " %s=%q", k, v)
	}
	fmt.Fprint(w, ">")
}

// AddAttr sets the key in the attribute map.
func (e *Element) AddAttr(k, v string) {
	e.attrs[k] = v
}

// AddChild adds a child node to the element.
func (e *Element) AddChild(Element) {
	e.Children = append(e.Children)
}
