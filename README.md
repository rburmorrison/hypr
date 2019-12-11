# Hypr

A library for programmatically generating HTML in Go.

# Install

To install, run `go get -u github.com/rburmorrison/hypr/...`.

# Usage

Hypr is used by creating elements and adding children to them. As a base, a document type is provided to you. Here's a
simple example:

```go
package main

import (
    "os"

    "github.com/rburmorrison/hypr/html"
)

func main() {
    // Create a base document with a head
    doc := html.NewDocument()
    doc.Add(html.NewHead())

    // Create a body with a paragraph tag that contains "Hello, world!"
    body := html.NewBody()
    body.Add(html.NewPara("Hello, world!"))
    doc.Add(body)
    
    // Write the result to the terminal
    doc.Write(os.Stdout)
}
```

Tags that are more complicated can be created with one of the fundamental three element types: `TextElement`,
`TagElement`, and `SoloTagElement`. For example, to get the output: `<!DOCTYPE html><head><title>My
Title</title></head><body></body>`, you might write the following code:

```go
package main

import (
    "os"

    "github.com/rburmorrison/hypr/html"
)

func main() {
    // Create an empty document for the <!DOCTYPE html> declaration
    doc := html.NewDocument()
    
    // Define the title element
    title := html.NewTagElement("title")
    title.Add(html.NewTextElement("My Title"))

    // Add the head with the title and an empty body
    doc.Add(html.NewHead(title), html.NewBody())
    
    // Write the result to the terminal
    doc.Write(os.Stdout)
}
```