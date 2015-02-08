package down

const (
	EOF = -(iota + 1)
	Text
)

type scanner struct {
	text string
	t    token
	loc  position
}

type token struct {
	text  string
	ident int
}

type position struct {
	begin, end int
}

func (s *scanner) Next() {
	s.t = token{}
	s.loc.begin = s.loc.end

	for s.loc.end < len(s.text) {
		c := s.text[s.loc.end]
		if c == ' ' {
			for s.text[s.loc.end] == ' ' {
				s.loc.end += 1
			}
			break
		}

		s.loc.end += 1
	}
	// Identify type

	// Modify text based on that.
}

func (s *scanner) Token() token {
	return token{s.text[s.loc.begin:s.loc.end], Text}
}

func (s scanner) End() bool {
	if s.loc.end == len(s.text) {
		return true
	}

	return false
}
