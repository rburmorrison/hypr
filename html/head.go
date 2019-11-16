package html

import (
	"fmt"
	"io"
)

// Head represents the header in HTML.
type Head struct {
	Children ElGroup

	attrs attrs
}

func (h *Head) Write(w io.Writer) {
	fmt.Fprint(w, "<head")
	h.attrs.write(w)
	fmt.Fprint(w, ">")

	h.Children.Write(w)
	fmt.Fprint(w, "</head>")
}
