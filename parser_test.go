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
	fmt.Println(p.String())
}
