package down

import (
	"fmt"
)

// Noder is the interface used to describe a down Node that will be converted to html.
type Noder interface {
	String() string
}

// Node is a basic Node to show how Nodes should be constructed.
type Node struct{}

func (n Node) String() string {
	return ""
}

// GroupNode is a Node Primitive for nodes that contain n amounts of children.
type GroupNode struct {
	children []Noder
}

func (gn GroupNode) Content() string {
	var content string
	for _, n := range gn.children {
		content += n.String()
	}
	return content
}

// Add a node to a composite Node
func (gn *GroupNode) AddChild(n Noder) {
	gn.children = append(gn.children, n)
}

// LineNode describes a line of markdown.
type LineNode struct {
	Child Noder
}

func (ln LineNode) String() string {
	return ln.Child.String() + "\n"
}

// CompositeStringNode describes a string which can contain any number of RawTextNodes, ItalicNodes,
// BoldNodes and LinkNodes.
type CompositeStringNode struct {
	GroupNode
}

func (csn CompositeStringNode) String() string {
	return csn.Content()
}

// RawTextNode is a node that represents bare text.
type RawTextNode struct {
	Content string
}

func (rtn RawTextNode) String() string {
	return rtn.Content
}

// LinkNode is a node that represents a html link.
type LinkNode struct {
	Text Noder
	Addr Noder
}

func (ln LinkNode) String() string {
	return fmt.Sprintf("<a href='%v'>%v</a>", ln.Addr.String(), ln.Text.String())
}

// BoldStringNode is a node to represent bold text.
type BoldStringNode struct {
	Child CompositeStringNode
}

func (bsn BoldStringNode) String() string {
	return "<b>" + bsn.Child.String() + "</b>"
}

// ItalicStringNode is a node to represent italic text.
type ItalicNode struct {
	Child CompositeStringNode
}

func (in ItalicNode) String() string {
	return "<i>" + in.Child.String() + "</i>"
}

// HeaderNode is a node that can represent any html header tag (h[1...6])
type HeaderNode struct {
	Child RawTextNode
	Level int
}

func (hn HeaderNode) String() string {
	return fmt.Sprintf("<h%[1]v>%[2]v</h%[1]v>", hn.Level, hn.Child.String())
	// return fmt.Sprintf("%[2] %[1]", ln., 22)
}

// ParagraphNode is a node to represent a paragraph of raw, links, bold and italic text.
type ParagraphNode struct {
	Child CompositeStringNode
}

func (pn ParagraphNode) String() string {
	return fmt.Sprintf("<p>%v</p>", pn.Child.String())
}

// UnorderedListNode is a node to represent a html unordered list element
type UnorderedListNode struct {
	GroupNode
}

func (uln UnorderedListNode) String() string {
	return fmt.Sprintf("<ul>%v</ul>", uln.Content())
}

// UnorderedListItemNode is a node to represent a list element node, which is inside an
// unordered list.
type UnorderedListItemNode struct {
	Child Noder
}

func (ulin UnorderedListItemNode) String() string {
	return fmt.Sprintf("<li>%v</li>", ulin.Child.String())
}
