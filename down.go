package down

func Convert(input string) string {
	ts := Lex(input + "\n")

	return Parse(ts)
}
