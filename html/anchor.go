package html

import "io"

// Anchor represents the anchor element in HTML.
type Anchor struct {
	Children ElGroup

	attrs attrs
}

// NewAnchor creates a new Anchor and returns a
// reference to it.
func NewAnchor() *Anchor {
	return &Anchor{}
}

func (a *Anchor) Write(w io.Writer) {}

func (a *Anchor) AddAttr(key string, value string) {}

func (a *Anchor) AddChild(el Element) {}
