package down

//go:generate stringer -type=TokenType
import (
	"errors"
	"fmt"
)

type FindFunc func(l *lexer, s string) ([]Token, error)

func Lex(s string) ([]Token, error) {
	l := lexer{
		location: 0,
		source:   s,
		finders:  []FindFunc{HeaderFinder},
	}

	return l.lex()
}

type lexer struct {
	location int
	source   string
	finders  []FindFunc
}

func (l *lexer) lex() ([]Token, error) {
	return HeaderFinder(l, l.source)
}

func (l *lexer) End() bool {
	if l.location >= len(l.source) {
		return true
	}

	return false
}

func (l *lexer) Peek() rune {
	return rune(l.source[l.location])
}

func (l *lexer) Next() {
	l.location += 1
}

func (t Token) Is(typ TokenType) bool {
	if t.typ == typ {
		return true
	}

	return false
}

type Token struct {
	typ TokenType
	val string
}

func (t Token) String() string {
	return fmt.Sprint(t.typ)
}

type TokenType int

const (
	HeaderSixType TokenType = iota
	HeaderOneType
	HeaderTwoType
	TextType
	HeaderEndType
)

func HeaderFinder(l *lexer, s string) ([]Token, error) {
	var tokens []Token
	if l.Peek() != '#' {
		return tokens, errors.New("Not a token")
	}

	l.Next()
	c := 1
	for !l.End() {
		if l.Peek() != '#' {
			break
		}

		c += 1
		l.Next()
	}

	t := Token{}
	switch c {
	case 1:
		t.typ = HeaderOneType
	case 2:
		t.typ = HeaderTwoType
	}
	tokens = append(tokens, t)

	var content string
	for !l.End() {
		c := l.Peek()
		if c == '\n' {
			break
		}

		content += string(c)
		l.Next()
	}

	text := Token{typ: TextType, val: content}
	tokens = append(tokens, text)

	l.Next()
	end := Token{typ: HeaderEndType, val: ""}
	tokens = append(tokens, end)

	return tokens, nil
}
