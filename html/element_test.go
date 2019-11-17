package html

import (
	"bytes"
	"testing"
)

func TestElement(t *testing.T) {
	expected := `<a href="https://google.com/"></a>`

	var buf bytes.Buffer
	a := NewTagElement("a")
	a.AddAttr("href", "https://google.com/")
	a.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewLink(t *testing.T) {
	expected := `<a href="https://google.com/">Click here!</a>`

	var buf bytes.Buffer
	link := NewLink("https://google.com/", "Click here!")
	link.Write(&buf)

	assertExpectation(t, expected, buf.String())
}
