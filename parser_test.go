package down

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	p := Parser{}
	p.source = "#Hello \nmy *name* is *harrison*"
	fmt.Println(p.source)
}
