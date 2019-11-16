package html

import (
	"fmt"
	"io"
)

// Element represents any HTML element.
type Element interface {
	Write(io.Writer)
	AddAttr(string, string)
	AddChild(Element)
}

type attrs map[string]string

func (a attrs) write(w io.Writer) {
	for k, v := range a {
		fmt.Fprintf(w, " %s=%q", k, v)
	}
}
