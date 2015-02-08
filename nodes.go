package down

import (
	"fmt"
)

type Noder interface {
	String() string
}

type Node struct {
}

func (n Node) String() string {
	return ""
}

type LineNode struct {
	Child Noder
}

func (ln LineNode) String() string {
	return ln.Child.String()
}

type CompositeStringNode struct {
	Special Noder
	Child   Noder
}

func (csn CompositeStringNode) String() string {
	return "<p>" + csn.Special.String() + csn.Child.String() + "</p>"
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
