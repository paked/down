package down

import (
	"fmt"
	"testing"
)

func TestLex(t *testing.T) {
	fmt.Println(Lex("###Hello\nHey#wut"))
}
