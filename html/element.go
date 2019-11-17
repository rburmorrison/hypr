package html

import (
	"io"
)

// Element represents any HTML element.
type Element interface {
	Write(io.Writer)
	AddAttr(string, string)
	AddChild(Element)
}
