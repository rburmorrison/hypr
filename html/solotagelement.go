package html

import (
	"fmt"
	"io"
)

// SoloTagElement represents a TagElement that is self-closing.
type SoloTagElement struct {
	Tag string

	attrs map[string]string
}

// NewSoloTagElement creates a new SoloTagElement and returns a reference to it.
func NewSoloTagElement(tag string) *SoloTagElement {
	var el SoloTagElement
	el.Tag = tag
	el.attrs = make(map[string]string)

	return &el
}

func (e SoloTagElement) Write(w io.Writer) {
	fmt.Fprintf(w, "<%s", e.Tag)
	for k, v := range e.attrs {
		if v != "" {
			fmt.Fprintf(w, " %s=%q", k, v)
		} else {
			fmt.Fprintf(w, " %s", k)
		}
	}
	fmt.Fprint(w, ">")
}

// SetAttr sets the key in the attribute map.
func (e *SoloTagElement) SetAttr(k, v string) {
	e.attrs[k] = v
}

// Add is only implemented to make SoloTagElement an Element.
func (e *SoloTagElement) Add(...Element) {}
