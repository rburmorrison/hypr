package html

import (
	"bytes"
	"testing"
)

func TestElement(t *testing.T) {
	expected := `<a href="https://google.com/"></a>`

	var buf bytes.Buffer
	a := NewTagElement("a")
	a.SetAttr("href", "https://google.com/")
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
	ol.Add(NewLI("Ordered Item 1"))
	ol.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewUL(t *testing.T) {
	expected := `<ul><li>Unordered Item 1</li></ul>`

	var buf bytes.Buffer
	ol := NewUL()
	ol.Add(NewLI("Unordered Item 1"))
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

func TestNewHTML(t *testing.T) {
	expected := `<html><head></head><body></body></html>`

	var buf bytes.Buffer
	h := NewHTML(NewHead(), NewBody())
	h.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewPara(t *testing.T) {
	expected := `<p>Hello, world!</p>`

	var buf bytes.Buffer
	p := NewPara("Hello, world!")
	p.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewDiv(t *testing.T) {
	expected := `<div><h1>My Title</h1><p>Hello, world!</p></div>`

	var buf bytes.Buffer
	div := NewDiv(
		NewH1("My Title"),
		NewPara("Hello, world!"),
	)
	div.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewBreak(t *testing.T) {
	expected := `<br>`

	var buf bytes.Buffer
	br := NewBreak()
	br.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestNewHR(t *testing.T) {
	expected := `<hr>`

	var buf bytes.Buffer
	br := NewHR()
	br.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestSoloTagElement_SetAttr(t *testing.T) {
	expected := `<test key="value">`

	var buf bytes.Buffer
	test := NewSoloTagElement("test")
	test.SetAttr("key", "value")
	test.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

func TestSoloTagElement_Write(t *testing.T) {
	expected := `<!DOCTYPE html>`

	var buf bytes.Buffer
	test := NewSoloTagElement("!DOCTYPE")
	test.SetAttr("html", "")
	test.Write(&buf)

	assertExpectation(t, expected, buf.String())
}

// TestTextElement_Add exists only to provide full coverage for the TextElement.
func TestTextElement_Add(t *testing.T) {
	e := NewTextElement("Hello, world!")
	e.Add(NewTextElement("Invalid"))
}

// TestTextElement_AddAttr exists only to provide full coverage for the TextElement.
func TestTextElement_AddAttr(t *testing.T) {
	e := NewTextElement("Hello, world!")
	e.SetAttr("key", "value")
}

// TestSoloTagElement_Add exists only to provide full coverage for the TextElement.
func TestSoloTagElement_Add(t *testing.T) {
	e := NewSoloTagElement("test")
	e.Add(NewTextElement("Hello, world!"))
}
