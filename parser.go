package down

import (
	"errors"
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
	if p.location == len(p.source) {
		return true
	}
	return false
}

func (p Parser) parseItalics() Noder {
	return ItalicNode{}
}

func (p Parser) parseBold() Noder {
	return BoldStringNode{}
}

func (p Parser) parseLink() Noder {
	return LinkNode{}
}

func (p Parser) parseLine() Noder {
	return LineNode{}
}

func (p *Parser) String() string {
	// Parse into line
	return p.parseLine().String()
}
