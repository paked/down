package down

import (
	"fmt"
	"testing"
)

func TestHeaders(t *testing.T) {
	fmt.Println("testing heraderz..")
	ts := Lex("#Hello\nHow are you today\n**Hello**\n*Hi* **drugs**\n")
	fmt.Println(ts)
}
