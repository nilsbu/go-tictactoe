package mechanics

import "testing"

func TestField_Put(t *testing.T) {
	tables := []struct {
		pos  Position
		post Marks
		err  error
	}{
		{[2]int{0, 0}, []Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, nil},
		{[2]int{1, 1}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, nil},
		{[2]int{4, 2}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, errAny},
		{[2]int{1, 1}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, errAny},
	}

	var field = Field{make(Marks, 3*3), 3}
	var currentPlayer Player

	for i, table := range tables {
		err := field.Put(table.pos, currentPlayer)
		if (err == nil) != (table.err == nil) {
			t.Errorf("unexpected error behavior in step %v: expected = \"%v\", actual = \"%v\"",
				i+1, table.err, err)
			continue
		}
		if err != nil {
			continue
		}
		if !field.Marks.Equal(table.post) {
			t.Errorf("field different in step %v:\nexpected:\n%v\n\nactual:\n%v",
				i+1, table.post, field)
		}

		currentPlayer = (currentPlayer + 1) % 2
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
