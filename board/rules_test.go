package board

import (
	"fmt"
	"testing"

	"github.com/nilsbu/go-tictactoe/test"
)

func TestDataIsFinished(t *testing.T) {
	const draw Player = -1
	//const open Player = 0

	testCases := []struct {
		marks    Marks
		size     int
		finished bool
		winner   Player
	}{
		// No winner
		{Marks{0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, false, draw},
		{Marks{2, 1, 2, 1, 1, 2, 1, 2, 1}, 3, true, draw},
		{Marks{2, 2, 1, 0, 1, 0, 2, 1, 0}, 3, false, draw},
		// Winner in row
		{Marks{1, 1, 1, 2, 2, 1, 1, 2, 2}, 3, true, 1},
		{Marks{1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 1, 1, 2, 1, 1, 2}, 4, true, 2},
		// Winner in column
		{Marks{1, 2, 3, 2, 1, 3, 1, 2, 3}, 3, true, 3},
		{Marks{1, 1, 2, 1, 1, 2, 2, 1, 1}, 3, true, 1},
		// Winner in diagonal
		{Marks{1, 2, 2, 1, 1, 2, 2, 2, 1}, 3, true, 1},
		{Marks{0, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 0}, 4, true, 2},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			bo := Data{Marks: tc.marks, Size: tc.size}
			finished, isDraw, winner := bo.IsFinished()
			switch false {
			case test.Cond(!tc.finished, !finished):
				t.Error("expected not finished but was")
			case test.Cond(!tc.finished, !isDraw):
				t.Error("if game is not finished, is cannot be a draw")
			case test.Cond(!tc.finished, winner == 0):
				t.Errorf("if game is not finished, winner cannot be %v", winner)
			case tc.finished:
			case finished:
				t.Error("expected finished but wasn't")
			case test.Cond(tc.winner == draw, isDraw):
				t.Error("expected draw but wasn't returned")
			case test.Cond(tc.winner == draw, winner == 0):
				t.Errorf("cannot have a winner in a draw, was %v", winner)
			case tc.winner != draw:
			case !isDraw:
				t.Error("when there is a winner, the game cannot be a draw")
			case tc.winner == winner:
				t.Errorf("wrong winner: expected %v, actual = %v", tc.winner, winner)
			}
		})
	}
}
