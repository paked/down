package down

import (
	"testing"
)

func TestParseHeaders(t *testing.T) {
	sources := []Source{{"#Hello", "<h1>Hello</h1>"}, {"##Hello", "<h2>Hello</h2>"}, {"####Hello", "<h4>Hello</h4>"}}
	test(sources, "headers", t)
}

func TestParseLink(t *testing.T) {
	sources := []Source{{"[Link](http://google.com)", `<p><a href='http://google.com'>Link</a></p>`},
		{"[*Link*](http://google.com)", `<p><a href='http://google.com'>*Link*</a></p>`},
		{"[*Link*]", "<p>*Link*]\n</p>"},
		{"(http://google.com)", "<p>(http://google.com)</p>"},
		{"[Link](Hello world", "<p>Hello world\n</p>"}}
	test(sources, "links", t)
}

func TestParseBold(t *testing.T) {
	sources := []Source{{"**this is bold text**", "<p><b>this is bold text</b></p>"}}
	test(sources, "bold", t)
}

func TestParseItalics(t *testing.T) {
	sources := []Source{{"*this is italic text*", "<p><i>this is italic text</i></p>"}}
	test(sources, "italics", t)
}

func TestParseUnorderedList(t *testing.T) {
	sources := []Source{{"- One\n- Two", "<ul><li>One</li><li>Two</li></ul>"}}
	test(sources, "unordered list", t)
}

type Source struct {
	In, Out string
}

func test(sources []Source, name string, t *testing.T) {
	for _, s := range sources {
		p := Parse(s.In)
		if r := (p == s.Out+"\n"); !r {
			t.Errorf("Error with %v: %v is not %v", name, p, s.Out)
		}
	}
}
