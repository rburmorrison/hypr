package html

import (
	"fmt"
	"io"
)

// Body represents the body in HTML.
type Body struct {
	Children ElGroup

	attrs attrs
}

func (h *Body) Write(w io.Writer) {
	fmt.Fprint(w, "<head")
	h.attrs.write(w)
	fmt.Fprint(w, ">")

	h.Children.Write(w)
	fmt.Fprint(w, "</head>")
}
