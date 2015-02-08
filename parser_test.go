package down

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	s := scanner{}
	s.text = "Hello would you like some bacon?"

	for !s.End() {
		s.Next()
		fmt.Println(s.Token())
	}
}
