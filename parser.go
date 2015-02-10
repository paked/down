package down

import (
	"errors"
	"fmt"
)

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
	p.location += 1
}

func (p Parser) End() bool {
	if p.location >= len(p.source) {
		return true
	}
	return false
}

func (p Parser) Back() {
	p.location -= 1
}

func (p Parser) parseBold() Noder {
	return BoldStringNode{}
}

func (p Parser) parseLink() Noder {
	return LinkNode{}
}

func (p *Parser) parseLine() Noder {
	line := LineNode{RawTextNode{""}}
	for !p.End() {
		c := p.source[p.location]
		switch c {
		case uint8('#'):
			line.Child = p.parseHeader()
		default:
			line.Child = ParagraphNode{Child: p.parseComposite()}
		}

		p.children = append(p.children, line)
		p.Next()
	}
	return LineNode{}
}

func (p *Parser) parseRaw() RawTextNode {
	var content string

	for !p.End() {
		c := p.source[p.location]
		if c == uint8(10) {
			break
		}

		if c == uint8('*') {
			p.location -= 1
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
	p.location -= 1
	return RawTextNode{"*" + content}
}

func (p *Parser) parseComposite() CompositeStringNode {
	composite := CompositeStringNode{}
	for !p.End() {
		c := p.source[p.location]
		if c == uint8(10) {
			break
		}

		switch c {
		case uint8('*'):
			p.Next()
			composite.AddChild(p.parseItalics())
		default:
			composite.AddChild(p.parseRaw())
		}

		p.Next()
	}
	return composite
}

func (p *Parser) parseHeader() HeaderOneNode {
	text := RawTextNode{}

	for !p.End() {
		if p.source[p.location] == uint8(10) {
			break
		}

		if p.source[p.location] == uint8('#') {
			p.Next()
			continue
		}

		text.Content += string(p.source[p.location])
		p.Next()
	}

	return HeaderOneNode{Child: text}
}

func (p *Parser) Parse(source string) {
	fmt.Printf("Parsing string:\n`%v`\n", source)
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
