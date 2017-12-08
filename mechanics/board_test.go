package mechanics

import (
	"testing"

	"go-tictactoe/test"
)

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
	var currentPlayer Player

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
			currentPlayer = (currentPlayer + 1) % 2
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
