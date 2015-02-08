package down

import (
	"fmt"
)

type Node interface {
	String() string
}

type LineNode struct {
	Child Node
}

func (ln LineNode) String() string {
	return ln.Child.String()
}

type CompositeStringNode struct {
	Special Node
	Child   Node
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
