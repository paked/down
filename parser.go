package down

import (
	"errors"
	"fmt"
)

func Parse(s string) string {
	p := Parser{}
	p.Parse(s)
	return p.String()
}

type Parser struct {
	source   string
	location int
	children []Noder
}

func (p Parser) Peek() (uint8, error) {
	if p.location+1 == len(p.source) {
		return ' ', errors.New("Cannot peek EOF")
	}

	return p.source[p.location+1], nil
}

func (p *Parser) Next() {
	p.location++
}

func (p Parser) End() bool {
	if p.location >= len(p.source) {
		return true
	}
	return false
}

func (p Parser) Back() {
	p.location--
}

func (p *Parser) parseLine() Noder {
	line := LineNode{RawTextNode{""}}
	for !p.End() {
		c := p.source[p.location]
		switch c {
		case uint8('#'):
			line.Child = p.parseHeader()
		case uint8('-'):
			line.Child = p.parseList()
		default:
			line.Child = ParagraphNode{Child: p.parseComposite()}
		}

		p.children = append(p.children, line)
		p.Next()
	}
	return LineNode{}
}

func (p *Parser) parseList() Noder {
	list := UnorderedListNode{}
	for !p.End() {
		c := p.source[p.location]
		peek, _ := p.Peek()
		if peek == uint8('-') {
			p.Next()
			c = p.source[p.location]
		}

		if c != uint8('-') {
			fmt.Println("FINISHED LIST WITH", string(c), c, p.source[:p.location])
			return list
		}

		p.Next()
		list.AddChild(UnorderedListItemNode{p.parseListItem()})
	}

	return list
}

func (p *Parser) parseListItem() Noder {
	if c := p.source[p.location]; c != uint8('-') && c != uint8(' ') {
		return p.parseComposite()
	}
	p.Next()
	return p.parseComposite()
}

func (p *Parser) parseRaw() RawTextNode {
	var content string

	for !p.End() {
		c := p.source[p.location]
		if c == uint8(10) {
			break
		}

		if c == uint8('*') || c == uint8('[') || c == uint8('-') || c == uint8('#') {
			break
		}

		content += string(p.source[p.location])
		p.Next()
	}

	return RawTextNode{Content: content}
}

func (p *Parser) parseItalics() Noder {
	var content string
	node := ItalicNode{}
	for !p.End() {
		c := p.source[p.location]
		if c == uint8(10) {
			break
		}

		if c == uint8('*') {
			node.Child.AddChild(RawTextNode{Content: content})
			return node
		}

		content += string(c)
		p.Next()
	}

	p.Back()
	return RawTextNode{"*" + content}
}

func (p *Parser) parseBold() Noder {
	var content string
	node := BoldStringNode{}
	if p.source[p.location] != uint8('*') {
		p.Back()
	} else {
		p.Next()
	}

	for !p.End() {
		c := p.source[p.location]

		if c == uint8(10) {
			break
		}

		if c == uint8('*') {
			peek, _ := p.Peek()
			p.Next()
			if peek != c {
				node.Child.AddChild(RawTextNode{Content: content})
				node.Child.AddChild(p.parseItalics())
				content = ""
				continue
			}
			p.Next()

			node.Child.AddChild(RawTextNode{Content: content})
			return node
		}

		content += string(c)

		p.Next()
	}

	return RawTextNode{"**" + content}
}

func (p *Parser) parseComposite() CompositeStringNode {
	composite := CompositeStringNode{}

	for !p.End() {
		c := p.source[p.location]
		if c == uint8(10) {
			break
		}

		if c == uint8(')') || c == uint8('-') {
			break
		}

		if c == uint8('#') {
			fmt.Println("BREAKING")
			p.location--
			peek, _ := p.Peek()
			fmt.Println("PEEK: ", string(peek))
			break
		}

		switch c {
		case uint8('*'):
			peek, _ := p.Peek()
			p.Next()
			if peek == c {
				p.Next()
				composite.AddChild(p.parseBold())
				continue
			}

			composite.AddChild(p.parseItalics())
		case uint8('['):
			composite.AddChild(p.parseLink())
		default:
			composite.AddChild(p.parseRaw())
			fmt.Println("up to", p.source[:p.location])
			if p.source[p.location] == uint8('[') || p.source[p.location] == uint8('#') || p.source[p.location] == uint8(10) {
				continue
			}
		}

		if c == uint8(10) {
			break
		}

		p.Next()
	}
	return composite
}

func (p *Parser) parseLink() Noder {
	link := LinkNode{}
	var place string
	if p.source[p.location] == uint8('[') {
		p.Next()
	}

	for !p.End() {
		c := p.source[p.location]
		switch c {
		case ']':
			link.Text = RawTextNode{place}
		case '(':
			place = ""
			p.Next()
			continue
		case ')':
			link.Addr = RawTextNode{place}
			return link
		}

		place += string(c)
		p.Next()
	}

	return RawTextNode{Content: place}
}

func (p *Parser) parseHeader() HeaderNode {
	text := RawTextNode{}
	var count int

	for !p.End() {
		if p.source[p.location] == uint8(10) {
			break
		}

		if p.source[p.location] == uint8('#') {
			p.Next()
			count++
			continue
		}

		text.Content += string(p.source[p.location])
		p.Next()
	}

	return HeaderNode{Child: text, Level: count}
}

func (p *Parser) Parse(source string) {
	fmt.Printf("\nParsing string:\n`%v`\n", source)
	p.source = source + "\n"
	p.parseLine()
}

func (p *Parser) String() string {
	var content string
	for _, node := range p.children {
		content += node.String()
	}
	return content
}
