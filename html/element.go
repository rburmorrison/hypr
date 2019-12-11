package html

import (
	"fmt"
	"io"
)

// Element represents any HTML element.
type Element interface {
	Write(w io.Writer)
	SetAttr(k, v string)
	Add(...Element)
}

// NewLink creates a new Element preformatted as an anchor and returns a reference to it.
func NewLink(link, text string) Element {
	a := NewTagElement("a")
	a.Add(NewTextElement(text))
	a.SetAttr("href", link)

	return a
}

// NewOL creates a new Element preformatted as an unordered list and returns a reference to it.
func NewOL(children ...Element) Element {
	ol := NewTagElement("ol")
	ol.Add(children...)
	return ol
}

// NewUL creates a new Element preformatted as an unordered list and returns a reference to it.
func NewUL(children ...Element) Element {
	ul := NewTagElement("ul")
	ul.Add(children...)
	return ul
}

// NewLI creates a new Element preformatted as an item in an ordered or unordered list and returns a reference to it.
func NewLI(text string) Element {
	li := NewTagElement("li")
	li.Add(NewTextElement(text))

	return li
}

// NewHeader creates a new TagElement for any header tag from level 1-6 and returns it. If level is less than 1, level
// will be corrected to 1. If level is greater than 6, level will be corrected to six.
func NewHeader(level int, text string) Element {
	if level < 1 {
		level = 1
	} else if level > 6 {
		level = 6
	}

	tag := fmt.Sprintf("h%d", level)
	header := NewTagElement(tag)
	header.Add(NewTextElement(text))

	return header
}

// NewH1 creates a new level 1 header tag.
func NewH1(text string) Element {
	return NewHeader(1, text)
}

// NewH2 creates a new level 2 header tag.
func NewH2(text string) Element {
	return NewHeader(2, text)
}

// NewH3 creates a new level 2 header tag.
func NewH3(text string) Element {
	return NewHeader(3, text)
}

// NewHead creates a new head tag with the specified children.
func NewHead(children ...Element) Element {
	return newWithChildren("head", children)
}

// NewBody creates a new body tag with the specified children.
func NewBody(children ...Element) Element {
	return newWithChildren("body", children)
}

// NewHTML creates a new html tag with the specified children.
func NewHTML(children ...Element) Element {
	return newWithChildren("html", children)
}

// NewPara creates a new paragraph tag with the given text.
func NewPara(text string) Element {
	p := NewTagElement("p")
	p.Add(NewTextElement(text))

	return p
}

// NewDiv creates a div tag with the given children.
func NewDiv(children ...Element) Element {
	return newWithChildren("div", children)
}

// NewBreak creates a break tag.
func NewBreak() Element {
	br := NewSoloTagElement("br")
	return br
}

// NewHR creates a new hr tag.
func NewHR() Element {
	hr := NewSoloTagElement("hr")
	return hr
}

func newWithChildren(tag string, children []Element) Element {
	el := NewTagElement(tag)
	for _, child := range children {
		el.Add(child)
	}

	return el
}
