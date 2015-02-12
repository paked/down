package down

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	p := Parser{}
	source := "#Hello\nmy *name* is *harrison*\nHello"
	p.Parse(source)
	fmt.Println(p.String())
}

func TestParseLink(t *testing.T) {
	p := Parser{}
	source := "#Hello\nthis is a link [Go to google!](http://google.com)"
	p.Parse(source)
	// fmt.Println(p.String())
}

func TestParseBold(t *testing.T) {
	p := Parser{}
	source := "**hello** ***hey how are you* how are you**"
	p.Parse(source)
	fmt.Println(p.String())
	// fmt.Println(p)
}

func TestHeaderNode(t *testing.T) {
	h := HeaderNode{}
	h.Level = 1
	fmt.Println(h.String())
}

func TestParseHeader(t *testing.T) {
	p := Parser{}
	source := "#Hello\n##Hello"
	p.Parse(source)
	fmt.Println(p.String())
}
