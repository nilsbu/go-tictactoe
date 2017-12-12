package actor

import (
	"fmt"
	"testing"

	m "go-tictactoe/board"
	"go-tictactoe/test"
)

type moveCode int

const (
	ok moveCode = iota
	fail
	quit
)

func TestIsAcceptableMove(t *testing.T) {
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
		t.Run(fmt.Sprintf("#%v: \"%v\"", i, tc.s), func(t *testing.T) {
			switch pos, msg, err := isAcceptableMove(b, tc.s); false {
			case test.Cond(tc.status == quit, err != nil):
				t.Errorf("error expected but not returned")
			case test.Cond(tc.status == quit, msg == ""):
				t.Errorf("message has to be \"\" since quit is requested")
			case tc.status != quit:
			case err == nil:
				t.Errorf("no error expected but was returned")
			case test.Cond(tc.status == fail, msg != ""):
				t.Errorf("expected failure but passed")
			case tc.status == ok:
			case msg == "":
				t.Errorf("no message should have been returned")
			case tc.pos == pos:
				t.Errorf("false position, expected = %v, actual = %v", tc.pos, pos)
			}
		})
	}
}
