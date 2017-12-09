package mechanics

import (
	"fmt"
	"testing"

	"go-tictactoe/test"
)

func TestNewPosition(t *testing.T) {
	testCases := []struct {
		i int
		s int
		p Position
	}{
		{0, 2, Position{0, 0}},
		{1, 2, Position{1, 0}},
		{8, 3, Position{2, 2}},
		{8, 4, Position{0, 2}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.p), func(t *testing.T) {
			if p := NewPosition(tc.i, tc.s); tc.p != p {
				t.Errorf("expected = %v, actual = %v", tc.p, p)
			}
		})
	}
}

func TestBoard_Put(t *testing.T) {
	tables := []struct {
		pos  Position
		post Marks
		err  test.ErrorAnticipation
	}{
		{[2]int{0, 0}, []Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, test.NoError},
		{[2]int{1, 1}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, test.NoError},
		{[2]int{4, 2}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, test.AnyError},
		{[2]int{1, 1}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, test.AnyError},
	}

	var b = Board{make(Marks, 3*3), 3}
	var currentPlayer Player = 1

	for i, table := range tables {
		switch err := b.Put(table.pos, currentPlayer); false {
		case test.Cond(table.err == test.AnyError, err != nil):
			t.Errorf("error expected in step %v but none was returned", i+1)
		case test.Cond(table.err == test.NoError, err == nil):
			t.Errorf("no error expected in step %v but one was returned", i+1)
		case b.Marks.Equal(table.post):
			t.Errorf("board different in step %v:\nexpected:\n%v\n\nactual:\n%v",
				i+1, table.post, b)
		}

		if table.err == test.NoError {
			currentPlayer = currentPlayer%2 + 1
		}
	}
}

func (marks Marks) Equal(other Marks) bool {
	if len(marks) != len(other) {
		return false
	}
	for i := 0; i < len(marks); i++ {
		if marks[i] != other[i] {
			return false
		}
	}
	return true
}
