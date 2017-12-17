package board

import (
	"fmt"
	"testing"

	"github.com/nilsbu/fastest"
)

func TestNewPosition(t *testing.T) {
	ft := fastest.T{T: t}

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
		ft.Seq(fmt.Sprintf("#%v: %v", i, tc.p), func(ft fastest.T) {
			p := NewPosition(tc.i, tc.s)
			ft.Equals(tc.p, p)
		})
	}
}

func TestBoard_Put(t *testing.T) {
	ft := fastest.T{T: t}

	tables := []struct {
		pos  Position
		post Marks
		err  bool
	}{
		{[2]int{0, 0}, []Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, false},
		{[2]int{1, 1}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, false},
		{[2]int{4, 2}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, true},
		{[2]int{1, 1}, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, true},
	}

	var b = Data{make(Marks, 3*3), 3}
	var currentPlayer Player = 1

	for i, table := range tables {
		ok, _ := b.Put(table.pos, currentPlayer)

		ft.Implies(table.err == true, !ok, "error expected in step %v but none was returned", i+1)
		ft.Implies(table.err == false, ok, "no error expected in step %v but one was returned", i+1)
		ft.True(b.Marks.Equal(table.post), "board different in step %v:\nexpected:\n%v\n\nactual:\n%v", i+1, table.post, b)

		if table.err == false {
			currentPlayer = currentPlayer%2 + 1
		}
	}
}
