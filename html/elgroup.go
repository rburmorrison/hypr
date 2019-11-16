package html

import "io"

// ElGroup represents a list of elements.
type ElGroup []Element

func (eg ElGroup) Write(w io.Writer) {
	for _, el := range eg {
		el.Write(w)
	}
}
