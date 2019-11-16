package html

import (
	"bytes"
	"testing"
)

func TestEmptyDocument(t *testing.T) {
	expected := `<!DOCTYPE html><html></html>`

	var buf bytes.Buffer
	doc := NewDocument()
	doc.Write(&buf)

	if expected != buf.String() {
		t.Errorf("unexpected output, expected %q, got %q", expected, buf.String())
	}
}
