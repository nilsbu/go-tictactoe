package actor

import (
	"testing"

	"go-tictactoe/mechanics"
)

const (
	ok = iota
	fail
	quit
)

func TestIsAcceptableMove(t *testing.T) {
	b := mechanics.Board{Marks: make(mechanics.Marks, 9), Size: 3}
	b.Marks[3] = 1

	tables := []struct {
		input  string
		pos    mechanics.Position
		status int
	}{
		{"quit", mechanics.Position{0, 0}, quit},
		{"exit ", mechanics.Position{0, 0}, quit},
		{" exit", mechanics.Position{0, 0}, quit},
		{"asd", mechanics.Position{0, 0}, fail},
		{"x4,1", mechanics.Position{0, 0}, fail},
		{"4,1,2", mechanics.Position{0, 0}, fail},
		{"1232", mechanics.Position{0, 0}, fail},
		{"", mechanics.Position{0, 0}, fail},
		{"1,0", mechanics.Position{1, 0}, ok},
		{"1,01", mechanics.Position{1, 1}, ok},
		{"2, 1", mechanics.Position{2, 1}, ok},
		{"3,2", mechanics.Position{0, 0}, fail}, // out of bounds
		{"0,1", mechanics.Position{0, 0}, fail}, // field not free
		// no need to test bounds and written fields, see board tests
	}

	for i, table := range tables {
		pos, msg, err := isAcceptableMove(b, table.input)

		if table.status == quit {
			if err == nil {
				t.Errorf("error expected in step %v but not returned", i+1)
			} else if msg != "" {
				t.Errorf("message has to be \"\" in step %v since quit is requested", i+1)
			}
			continue
		}
		if err != nil {
			if table.status != quit {
				t.Errorf("no error expected in step %v but was returned", i+1)
			}
			continue
		}
		if table.status == fail {
			if msg == "" {
				t.Errorf("expected failure in step %v but passed", i+1)
			}
			continue
		}
		if msg != "" {
			if table.status != fail {
				t.Errorf("failed unexpectedly in step %v", i+1)
			}
			continue
		}
		if table.pos != pos {
			t.Errorf("false position in step %v, expected = %v, actual = %v",
				i+1, table.pos, pos)
		}
	}
}
