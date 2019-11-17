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
