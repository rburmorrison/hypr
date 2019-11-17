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

func TestNewLI(t *testing.T) {
	expected := `<li>Item 1</li>`

	var buf bytes.Buffer
	li := NewLI("Item 1")
	li.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewOL(t *testing.T) {
	expected := `<ol><li>Ordered Item 1</li></ol>`

	var buf bytes.Buffer
	ol := NewOL()
	ol.AddChild(NewLI("Ordered Item 1"))
	ol.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewUL(t *testing.T) {
	expected := `<ul><li>Unordered Item 1</li></ul>`

	var buf bytes.Buffer
	ol := NewUL()
	ol.AddChild(NewLI("Unordered Item 1"))
	ol.Write(&buf)

	assertExpectation(t, expected, buf.String())
}
