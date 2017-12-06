package rules

import (
	"go-tictactoe/mechanics"
	"testing"
)

func TestGetWinner(t *testing.T) {
	tables := []struct {
		marks  []mechanics.Player
		size   int
		winner mechanics.Player
	}{
		// No winner
		{[]mechanics.Player{0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, NoWinner},
		{[]mechanics.Player{2, 1, 2, 1, 1, 2, 1, 2, 1}, 3, NoWinner},
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
		field := mechanics.Field{Marks: table.marks, Size: table.size}
		winner := GetWinner(field)

		if table.winner != winner {
			t.Errorf("wrong winner in table %v: expected = %v, actual = %v", i+1,
				table.winner, winner)
		}
	}
}
