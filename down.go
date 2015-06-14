package down

func Convert(input string) string {
	ts := Lex(input)
	return Parse(ts)
}
