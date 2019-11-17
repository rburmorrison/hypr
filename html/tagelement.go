package html

import (
	"fmt"
	"io"
)

// TagElement represents an HTML element with a tag.
type TagElement struct {
	Tag      string
	Children []Element

	attrs map[string]string
}

// NewTagElement creates a new TagElement and returns a
// reference to it.
func NewTagElement(tag string) *TagElement {
	var el TagElement
	el.Tag = tag
	el.attrs = make(map[string]string)

	return &el
}

func (e *TagElement) Write(w io.Writer) {
	fmt.Fprintf(w, "<%s", e.Tag)
	for k, v := range e.attrs {
		fmt.Fprintf(w, " %s=%q", k, v)
	}
	fmt.Fprint(w, ">")

	for _, child := range e.Children {
		child.Write(w)
	}

	fmt.Fprintf(w, "</%s>", e.Tag)
}

// AddAttr sets the key in the attribute map.
func (e *TagElement) AddAttr(k, v string) {
	e.attrs[k] = v
}

// AddChild adds a child node to the element.
func (e *TagElement) AddChild(child Element) {
	e.Children = append(e.Children, child)
}
