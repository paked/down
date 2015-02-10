package down

import (
	"fmt"
)

type Noder interface {
	String() string
}

type Node struct{}

func (n Node) String() string {
	return ""
}

type LineNode struct {
	Child Noder
}

func (ln LineNode) String() string {
	return ln.Child.String() + "\n"
}

type CompositeStringNode struct {
	children []Noder
}

func (csn CompositeStringNode) String() string {
	var content string
	for _, node := range csn.children {
		content += node.String()
	}

	return content
}

func (csn *CompositeStringNode) AddChild(n Noder) {
	csn.children = append(csn.children, n)
}

type RawTextNode struct {
	Content string
}

func (rtn RawTextNode) String() string {
	return rtn.Content
}

type LinkNode struct {
	Text RawTextNode
	Link RawTextNode
}

func (ln LinkNode) String() string {
	return fmt.Sprintf("<a href='%v'>%v</a>", ln.Text.String(), ln.Link.String())
}

type BoldStringNode struct {
	Child CompositeStringNode
}

func (bsn BoldStringNode) String() string {
	return "<b>" + bsn.Child.String() + "</b>"
}

type ItalicNode struct {
	Child CompositeStringNode
}

func (in ItalicNode) String() string {
	return "<i>" + in.Child.String() + "</i>"
}

type HeaderOneNode struct {
	Child RawTextNode
}

func (hon HeaderOneNode) String() string {
	return fmt.Sprintf("<h1>%v</h1>", hon.Child.String())
}

type ParagraphNode struct {
	Child CompositeStringNode
}

func (pn ParagraphNode) String() string {
	return fmt.Sprintf("<p>%v</p>", pn.Child.String())
}
