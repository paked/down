package down

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	fmt.Println("testing heraderz..")
	ts := Lex("* *Hello*\n")
	fmt.Println(ts)
}
