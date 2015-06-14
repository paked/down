package down

import "fmt"

type Parser struct {
	source   []Token
	location int
}

func (p *Parser) Parse() string {
	var res string

	for !p.End() {
		t := p.Peek()
		switch t.typ {
		case BigHeaderToken:
			res += fmt.Sprintf("<h1>%v</h1>", t.val)
		case MediumHeaderToken:
			res += fmt.Sprintf("<h2>%v</h2>", t.val)
		case SmallHeaderToken:
			res += fmt.Sprintf("<h3>%v</h3>", t.val)
		case ParagraphStartToken:
			res += fmt.Sprintf("<p>")
		case ParagraphEndToken:
			res += fmt.Sprintf("</p>")
		case TextToken:
			res += fmt.Sprintf("%v", t.val)
		case BoldStartToken:
			res += fmt.Sprintf("<strong>")
		case BoldEndToken:
			res += fmt.Sprintf("</strong>")
		case ItalicStartToken:
			res += fmt.Sprintf("<em>")
		case ItalicEndToken:
			res += fmt.Sprintf("</em>")
		case ListStartToken:
			res += fmt.Sprintf("<ul>")
		case ListEndToken:
			res += fmt.Sprintf("</ul>")
		case ItemStartToken:
			res += fmt.Sprintf("<li>")
		case ItemEndToken:
			res += fmt.Sprintf("</li>")
		default:
			fmt.Printf("Unlexable token: %v\n", t)
		}

		p.Next()
	}

	return res
}

func (p *Parser) Next() {
	p.location += 1
}

func (p *Parser) End() bool {
	if p.location >= len(p.source) {
		return true
	}

	return false
}

func (p *Parser) Peek() Token {
	return p.source[p.location]
}

func Parse(tokens []Token) string {
	p := Parser{
		source:   tokens,
		location: 0,
	}

	return p.Parse()
}
