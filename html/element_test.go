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

func TestNewHeader(t *testing.T) {
	expected := `<h5>Heading 5</h5>`

	var buf bytes.Buffer
	h := NewHeader(5, "Heading 5")
	h.Write(&buf)

	assertExpectation(t, expected, buf.String())

	expected = `<h1>Heading 1</h1>`

	buf = bytes.Buffer{}
	h = NewHeader(0, "Heading 1")
	h.Write(&buf)

	assertExpectation(t, expected, buf.String())

	expected = `<h6>Heading 6</h6>`

	buf = bytes.Buffer{}
	h = NewHeader(10, "Heading 6")
	h.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewH1(t *testing.T) {
	expected := `<h1>Heading 1</h1>`

	var buf bytes.Buffer
	h := NewH1("Heading 1")
	h.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewH2(t *testing.T) {
	expected := `<h2>Heading 2</h2>`

	var buf bytes.Buffer
	h := NewH2("Heading 2")
	h.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewH3(t *testing.T) {
	expected := `<h3>Heading 3</h3>`

	var buf bytes.Buffer
	h := NewH3("Heading 3")
	h.Write(&buf)

	assertExpectation(t, expected, buf.String())
}
