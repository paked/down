package down

import (
	"testing"
)

func TestParser(t *testing.T) {
	if Convert("#Hello\n") != "<h1>Hello</h1>" {
		t.Fail()
	}

	if Convert("##Hello\n") != "<h2>Hello</h2>" {
		t.Fail()
	}

	if Convert("###Hello\n") != "<h3>Hello</h3>" {
		t.Fail()
	}

	if Convert("**Hello**\n") != "<p><strong>Hello</strong></p>" {
		t.Fail()
	}

	if Convert("*Hello*\n") != "<p><em>Hello</em></p>" {
		t.Fail()
	}

	if Convert("**Hello** *Hello*\n") != "<p><strong>Hello</strong> <em>Hello</em></p>" {
		t.Fail()
	}

}
