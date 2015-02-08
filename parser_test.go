package down

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	fmt.Println(LineNode{Child: BoldStringNode{CompositeStringNode{RawTextNode{"Oh yes"}, Node{}}}})
}
