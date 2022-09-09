package main

import (
	"fmt"
	"strings"
)

const indent = 2

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indentSize int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)

	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))

	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}

	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElement{rootName, "", []HtmlElement{}},
	}
}

func main() {
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "Hello").
		AddChildFluent("li", "Hello word 2").
		AddChildFluent("li", "Hello word 2").
		AddChildFluent("li", "Hello word 2")

	fmt.Println(b.String())
}
