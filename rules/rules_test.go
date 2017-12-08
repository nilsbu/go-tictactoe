package rules

import (
	"testing"

	"go-tictactoe/mechanics"
	"go-tictactoe/test"
)

const noWinner mechanics.Player = -1

func TestGetWinner(t *testing.T) {
	tables := []struct {
		marks  []mechanics.Player
		size   int
		winner mechanics.Player
	}{
		// No winner
		{[]mechanics.Player{0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, noWinner},
		{[]mechanics.Player{2, 1, 2, 1, 1, 2, 1, 2, 1}, 3, noWinner},
		// Winner in row
		{[]mechanics.Player{1, 1, 1, 2, 2, 1, 1, 2, 2}, 3, 1},
		{[]mechanics.Player{1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 1, 1, 2, 1, 1, 2}, 4, 2},
		// Winner in column
		{[]mechanics.Player{1, 2, 3, 2, 1, 3, 1, 2, 3}, 3, 3},
		{[]mechanics.Player{1, 2, 2, 1, 1, 2, 2, 1, 1}, 3, 1},
		//Winner in diagonal
		{[]mechanics.Player{1, 2, 2, 1, 1, 2, 2, 2, 1}, 3, 1},
		{[]mechanics.Player{0, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 0}, 4, 2},
	}

	for i, table := range tables {
		b := mechanics.Board{Marks: table.marks, Size: table.size}
		switch winner, hasWinner := GetWinner(b); false {
		case test.Cond(table.winner == noWinner, !hasWinner):
			t.Errorf("no winner expected in step %v but one was returned", i+1)
		case test.Cond(table.winner != noWinner, hasWinner):
			t.Errorf("winner expected in step %v but none returned", i+1)
		case table.winner != noWinner:
		case table.winner == winner:
			t.Errorf("wrong winner in table %v: expected = %v, actual = %v", i+1,
				table.winner, winner)
		}
	}
}
