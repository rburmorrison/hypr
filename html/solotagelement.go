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

func NewSoloTagElement(tag string) *SoloTagElement {
	var el SoloTagElement
	el.Tag = tag
	el.attrs = make(map[string]string)

	return &el
}

func (e SoloTagElement) Write(w io.Writer) {
	fmt.Fprintf(w, "<%s", e.Tag)
	for k, v := range e.attrs {
		fmt.Fprintf(w, " %s=%q", k, v)
	}
	fmt.Fprint(w, ">")
}

func (e *SoloTagElement) SetAttr(k, v string) {
	e.attrs[k] = v
}

func (e *SoloTagElement) Add(...Element) {}
