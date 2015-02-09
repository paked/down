package down

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	p := Parser{}
	source := "#Hello \nmy *name* is *harrison*"
	p.Parse(source)
	fmt.Println(p.String())
}
