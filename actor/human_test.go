package actor

import (
	"fmt"
	"testing"

	"github.com/nilsbu/fastest"

	m "github.com/nilsbu/tictactoe/board"
)

type moveCode int

const (
	ok moveCode = iota
	fail
	quit
)

func TestIsAcceptableMove(t *testing.T) {
	ft := fastest.T{T: t}

	b := m.Data{Marks: make(m.Marks, 9), Size: 3}
	b.Marks[3] = 1

	testCases := []struct {
		s      string
		pos    m.Position
		status moveCode
	}{
		{"quit", m.Position{0, 0}, quit},
		{"exit ", m.Position{0, 0}, quit},
		{" exit", m.Position{0, 0}, quit},
		{"asd", m.Position{0, 0}, fail},
		{"x4,1", m.Position{0, 0}, fail},
		{"4,1,2", m.Position{0, 0}, fail},
		{"1232", m.Position{0, 0}, fail},
		{"", m.Position{0, 0}, fail},
		{"1,0", m.Position{1, 0}, ok},
		{"1,01", m.Position{1, 1}, ok},
		{"2, 1", m.Position{2, 1}, ok},
	}

	for i, tc := range testCases {

		ft.Seq(fmt.Sprintf("#%v: \"%v\"", i, tc.s), func(ft fastest.T) {
			pos, msg, err := isAcceptableMove(b, tc.s)

			ft.Implies(tc.status == quit, err != nil, "error expected but not returned")
			ft.Implies(tc.status == quit, msg == "", "message has to be \"\" since quit is requested")
			ft.Only(tc.status != quit)
			ft.Nil(err, "no error expected but was returned")
			ft.Implies(tc.status == fail, msg != "", "expected failure but passed")
			ft.Only(tc.status == ok)
			ft.Equals("", msg)
			ft.Equals(tc.pos, pos)
		})
	}
}
