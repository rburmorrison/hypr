package html

import (
	"bytes"
	"testing"
)

func TestEmptyDocument(t *testing.T) {
	expected := `<!DOCTYPE html>`

	var buf bytes.Buffer
	doc := NewDocument()
	doc.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestBasicDocument(t *testing.T) {
	expected := `<!DOCTYPE html><head></head><body></body>`

	var buf bytes.Buffer
	doc := NewDocument()
	doc.AddChild(NewHead())
	doc.AddChild(NewBody())

	doc.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestBaseDocument(t *testing.T) {
	expected := `<!DOCTYPE html><html><head></head><body></body></html>`

	var buf bytes.Buffer
	doc := NewDocument()
	doc.AddChild(NewHTML(NewHead(), NewBody()))
	doc.Write(&buf)

	assertExpectation(t, expected, buf.String())
}
