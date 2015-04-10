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
		finders:  []FindFunc{HeaderFinder, TextFinder},
	}

	return l.lex()
}

type lexer struct {
	location int
	source   string
	finders  []FindFunc
	tokens   []Token
}

func (l *lexer) lex() ([]Token, error) {
	for !l.End() {
		for _, f := range l.finders {
			if l.try(f) {
				break
			}
		}
	}

	return l.tokens, nil
}

func (l *lexer) try(f FindFunc) bool {
	old := l.location
	ts, err := f(l, l.source)

	if err != nil {
		l.location = old
		return false
	}

	l.tokens = append(l.tokens, ts...)

	return true
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

func TextFinder(l *lexer, s string) ([]Token, error) {
	var tokens []Token
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
	return tokens, nil
}

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
	texts, err := TextFinder(l, s)
	if err != nil {
		return tokens, err
	}
	tokens = append(tokens, texts...)

	l.Next()
	end := Token{typ: HeaderEndType, val: ""}
	tokens = append(tokens, end)

	return tokens, nil
}
