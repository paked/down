// syntax
// # Hello = <h1>Hello</h1>
// ## Hello = <h2>Hello</h2>
// ### Hello = <h2>Hello</h2>
//
// *hello* = <em>hello</em>
// **hello** = <strong></strong>
//
// * Hello
// * Goodbye
// 		= <ul><li>Hello</li><li>Goodbye</li><ul>
package down

import (
	"errors"
	"fmt"
	"unicode"
)

const (
	BigHeaderToken      string = "Big Header Token"
	MediumHeaderToken   string = "Medium Header Token"
	SmallHeaderToken    string = "Small Header Token"
	TerminatorToken     string = "Terminator"
	ParagraphStartToken string = "Paragraph Start"
	ParagraphEndToken   string = "Paragraph End"
	TextToken           string = "Text"
	BoldStartToken      string = "Bold Start"
	BoldEndToken        string = "Bold End"
	ItalicStartToken    string = "Italics Start"
	ItalicEndToken      string = "Italics End"
	ListStartToken      string = "List Start"
	ListEndToken        string = "List end"
)

type Lexer struct {
	source   string
	location int
}

func (l *Lexer) Lex() []Token {
	var tokens []Token
	for !l.End() {
		var ts []Token
		var err error

		if ts, err = l.try(l.titles, l.newline); err == nil {
			tokens = append(tokens, ts...)
			continue
		}

		if ts, err = l.try(l.text, l.newline); err == nil {
			tokens = append(tokens, ts...)
			continue
		}

		fmt.Println("You messed up.")
		break
	}

	return tokens
}

func (l *Lexer) try(fns ...func() ([]Token, error)) ([]Token, error) {
	var tokens []Token

	for _, fn := range fns {
		loc := l.location
		ts, err := fn()

		if err != nil {
			fmt.Println("error: ", err)
			l.location = loc

			return tokens, err
		}

		tokens = append(tokens, ts...)

	}

	return tokens, nil
}

func (l *Lexer) words() ([]Token, error) {
	var tokens []Token

	var content string
	for !l.End() {
		c := l.Peek()
		if c == '*' || c == '\n' || !unicode.IsPrint(c) {
			break
		}

		content += string(c)
		l.Next()
	}

	if content == "" || l.End() {
		return tokens, errors.New("no words or at end")
	}

	tokens = append(tokens, Token{TextToken, content})

	return tokens, nil
}

func (l *Lexer) match(sequence string) bool {
	if len(sequence) > (len(l.source) - l.location) {
		return false
	}

	cut := l.source[l.location : l.location+len(sequence)]
	if sequence != cut {
		return false
	}

	return true
}

func (l *Lexer) bold() ([]Token, error) {
	var tokens []Token

	if !l.match("**") {
		return tokens, errors.New("not a starter")
	}

	l.Next()
	l.Next()

	ts, err := l.words()
	if err != nil {
		fmt.Println(l.source[:l.location])
		return tokens, err
	}

	if !l.match("**") {
		return tokens, errors.New("not an ender")
	}

	l.Next()
	l.Next()

	tokens = append(tokens, Token{typ: BoldStartToken})
	tokens = append(tokens, ts...)
	tokens = append(tokens, Token{typ: BoldEndToken})

	return tokens, nil
}

func (l *Lexer) italic() ([]Token, error) {
	var tokens []Token

	if l.Peek() != '*' {
		return tokens, errors.New("not a correct italic starter")
	}

	l.Next()

	tokens = append(tokens, Token{typ: ItalicStartToken})

	ts, err := l.words()
	if err != nil {
		return tokens, err
	}

	tokens = append(tokens, ts...)

	if l.Peek() != '*' {
		return tokens, errors.New("not a correct itlaic ender")
	}

	l.Next()

	tokens = append(tokens, Token{typ: ItalicEndToken})

	return tokens, nil
}

func (l *Lexer) text() ([]Token, error) {
	var tokens []Token

	tokens = append(tokens, Token{typ: ParagraphStartToken})

	for !l.End() {
		c := l.Peek()
		if c == '\n' {
			break
		}

		if ts, err := l.try(l.words); err == nil {
			tokens = append(tokens, ts...)
		}

		if ts, err := l.try(l.bold); err == nil {
			tokens = append(tokens, ts...)
		}

		if ts, err := l.try(l.italic); err == nil {
			tokens = append(tokens, ts...)
		}

	}

	if len(tokens) == 1 || l.End() {
		return tokens, errors.New("no tokens or at end")
	}

	tokens = append(tokens, Token{typ: ParagraphEndToken})

	return tokens, nil
}

func (l *Lexer) newline() ([]Token, error) {
	var tokens []Token
	if l.Peek() != '\n' {
		return tokens, errors.New("Not a terminator")
	}

	l.Next()
	tokens = append(tokens, Token{TerminatorToken, ""})

	return tokens, nil
}

func (l *Lexer) titles() ([]Token, error) {
	var tokens []Token

	if l.Peek() != '#' {
		return tokens, errors.New("not a title!")
	}

	l.Next()

	level := 1

	for !l.End() {
		if l.Peek() != '#' {
			break
		}

		level += 1

		l.Next()
	}

	var content string
	for !l.End() {
		c := l.Peek()
		if !unicode.IsPrint(c) || c == '\n' {
			break
		}

		content += string(c)
		l.Next()
	}

	if content == "" {
		return tokens, errors.New("no content")
	}

	tok := Token{val: content}
	switch level {
	case 1:
		tok.typ = BigHeaderToken
	case 2:
		tok.typ = MediumHeaderToken
	case 3:
		tok.typ = SmallHeaderToken
	default:
		return tokens, errors.New("not a valid number of hashes")
	}

	tokens = append(tokens, tok)

	return tokens, nil
}

func (l *Lexer) Next() {
	l.location += 1
}

func (l *Lexer) End() bool {
	if l.location >= len(l.source) {
		return true
	}

	return false
}

func (l *Lexer) Peek() rune {
	return rune(l.source[l.location])
}

func Lex(input string) []Token {
	l := Lexer{
		source:   input,
		location: 0,
	}

	return l.Lex()
}

type Token struct {
	typ string
	val string
}

func (t Token) String() string {
	return t.typ + ` "` + t.val + `",`
}
