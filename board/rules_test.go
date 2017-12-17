package board

import (
	"fmt"
	"testing"

	"github.com/nilsbu/fastest"
)

func TestDataIsFinished(t *testing.T) {
	ft := fastest.T{T: t}

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
		ft.Seq(fmt.Sprintf("#%v: %v", i, tc.marks), func(ft fastest.T) {
			bo := Data{Marks: tc.marks, Size: tc.size}
			finished, isDraw, winner := bo.IsFinished()

			ft.Implies(!tc.finished, !finished, "expected not finished but was")
			ft.Implies(!tc.finished, !isDraw, "if game is not finished, is cannot be a draw")
			ft.Implies(!tc.finished, winner == 0, "if game is not finished, winner cannot be %v", winner)
			ft.Only(tc.finished)
			ft.True(finished, "expected finished but wasn't")
			ft.Implies(tc.winner == draw, isDraw, "expected draw but wasn't returned")
			ft.Implies(tc.winner == draw, winner == 0, "cannot have a winner in a draw, was %v", winner)
			ft.Only(tc.winner != draw)
			ft.True(!isDraw, "when there is a winner, the game cannot be a draw")
			ft.Equals(tc.winner, winner, "wrong winner: expected %v, actual = %v", tc.winner, winner)
		})
	}
}
