package html

import (
	"bytes"
	"testing"
)

func TestAnchor(t *testing.T) {
	expected := `<a href="https://google.com/"></a>`

	var buf bytes.Buffer
	a := NewAnchor()
	a.AddAttr("href", "https://google.com/")

	if expected != buf.String() {
		t.Errorf("unexpected output, expected %q, got %q", expected, buf.String())
	}
}
