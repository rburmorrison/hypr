package html


func NewAnchor(text, link str) *Element {
    return Element{
        "a",
        []Element,
        map[string]string{"href": text}
    }
}