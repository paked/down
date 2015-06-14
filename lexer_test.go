package down

import (
	"fmt"
	"testing"
)

func TestHeaders(t *testing.T) {
	fmt.Println("testing heraderz..")
	ts := Lex("* *Hello*\n")
	fmt.Println(ts)
}
