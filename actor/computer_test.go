package actor

import (
	m "go-tictactoe/mechanics"
	"go-tictactoe/test"
	"math"
	"testing"
)

func TestComputerGetMove(t *testing.T) {
	// More detailed tests in TestComputeOptimalMoveSeq
	tables := []struct {
		players m.Player
		id      m.Player
		marks   m.Marks
		idxs    []int
	}{
		{2, 1, m.Marks{1, 0, 0, 1, 2, 0, 2, 0, 0}, []int{2}},
	}

	for i, table := range tables {
		c := Computer{ID: table.id, Players: table.players}
		s := int(math.Sqrt(float64(len(table.marks))))
		b := m.Board{Marks: table.marks, Size: s}
		switch pos, err := c.GetMove(b); false {
		case err == nil:
			t.Errorf("must never return an error")
		case isIndexInList(pos.ToIndex(s), table.idxs):
			t.Errorf("in step %v, %v (= %v) must be in %v", i+1, pos, pos.ToIndex(s),
				table.idxs)
		}
	}
}

func TestComputeOptimalMoveSeq(t *testing.T) {
	const noWinner m.Player = -1

	tables := []struct {
		players m.Player
		id      m.Player
		marks   m.Marks
		idxs    []int
		winner  m.Player
	}{
		{2, 2, m.Marks{2, 0, 2, 0, 1, 0, 1, 1, 0}, []int{1}, 2},
		{2, 1, m.Marks{0, 0, 0, 2, 1, 0, 0, 0, 0}, []int{0, 1, 2, 6, 7, 8}, 1},
		{2, 1, m.Marks{0, 0, 0, 1, 2, 0, 0, 0, 0}, []int{0, 2, 6, 8}, noWinner},
		{2, 2, m.Marks{0, 0, 0, 1, 2, 0, 1, 0, 0}, []int{0}, noWinner},
		{2, 1, m.Marks{2, 0, 0, 1, 2, 0, 1, 0, 0}, []int{8}, noWinner},
		{2, 2, m.Marks{2, 0, 0, 1, 2, 0, 1, 0, 1}, []int{7}, noWinner},
		{2, 1, m.Marks{2, 0, 0, 1, 2, 0, 1, 2, 1}, []int{1}, noWinner},
		{2, 2, m.Marks{0, 0, 1, 1, 2, 0, 0, 0, 0}, []int{0, 1, 6, 7}, noWinner},
		{2, 1, m.Marks{0, 0, 1, 1, 2, 0, 0, 2, 0}, []int{1}, noWinner},
		{2, 2, m.Marks{1, 0, 1, 1, 2, 0, 0, 2, 0}, []int{1}, 2},
		{2, 2, m.Marks{1, 0, 0, 2, 1, 0, 0, 0, 0}, []int{1, 2, 5, 6, 7, 8}, 1},
		{2, 2, m.Marks{2, 0, 0, 0, 1, 0, 1, 0, 0}, []int{2}, noWinner},
	}

	for i, table := range tables {
		s := int(math.Sqrt(float64(len(table.marks))))
		marks := make(m.Marks, len(table.marks))
		copy(marks, table.marks)
		b := m.Board{Marks: marks, Size: s}
		p, w, hw := computeOptimalMoveSeq(b, table.id, table.players)
		switch false {
		case isBoardUnchanged(table.marks, b.Marks):
			t.Errorf("board changed")
		case test.Cond(!hw, w == 0):
			t.Errorf("in step %v hasWinner = false but winner = %v, must be 0",
				i+1, w)
		case test.Cond(table.winner != noWinner, hw):
			t.Errorf("in step %v, winner was expected but none was returned", i+1)
		case test.Cond(table.winner == noWinner, !hw):
			t.Errorf("in step %v no winner was expected but %v was returned", i+1, w)
		case test.Cond(table.winner != noWinner, table.winner == w):
			t.Errorf("in step %v, expected = %v, actual = %v", i+1, table.winner, w)
		case isIndexInList(p, table.idxs):
			t.Errorf("in step %v, %v must be in %v", i+1, p, table.idxs)
		}
	}
}

func isBoardUnchanged(a, b m.Marks) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isIndexInList(pidx int, idxs []int) bool {
	for _, idx := range idxs {
		if pidx == idx {
			return true
		}
	}
	return false
}
