package board

import (
	"fmt"
	"testing"

	"tictactoe/test"
)

func TestIsFull(t *testing.T) {
	testCases := []struct {
		marks Marks
		size  int
		fin   bool
	}{
		{Marks{0, 0, 0, 0}, 2, false},
		{Marks{2, 1, 2, 0}, 2, false},
		{Marks{2, 1, 2, 3}, 2, true},
		{Marks{2, 1, 2, 3, 1, 1, 2, 3, 2}, 3, true},
		{Marks{0, 1, 2, 3, 1, 1, 2, 3, 2}, 3, false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			bo := Data{Marks: tc.marks, Size: tc.size}
			switch fin := bo.IsFull(); false {
			case tc.fin == fin:
				t.Errorf("expected = %v, actual %v", tc.fin, fin)
			}
		})
	}
}

func TestGetWinner(t *testing.T) {
	const noWinner Player = -1

	testCases := []struct {
		marks  Marks
		size   int
		winner Player
	}{
		// No winner
		{Marks{0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, noWinner},
		{Marks{2, 1, 2, 1, 1, 2, 1, 2, 1}, 3, noWinner},
		// Winner in row
		{Marks{1, 1, 1, 2, 2, 1, 1, 2, 2}, 3, 1},
		{Marks{1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 1, 1, 2, 1, 1, 2}, 4, 2},
		// Winner in column
		{Marks{1, 2, 3, 2, 1, 3, 1, 2, 3}, 3, 3},
		{Marks{1, 2, 2, 1, 1, 2, 2, 1, 1}, 3, 1},
		// Winner in diagonal
		{Marks{1, 2, 2, 1, 1, 2, 2, 2, 1}, 3, 1},
		{Marks{0, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 0}, 4, 2},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			bo := Data{Marks: tc.marks, Size: tc.size}
			switch winner, hasWinner := bo.GetWinner(); false {
			case test.Cond(tc.winner == noWinner, !hasWinner):
				t.Errorf("no winner expected but one was returned")
			case test.Cond(tc.winner != noWinner, hasWinner):
				t.Errorf("winner expected but none returned")
			case tc.winner != noWinner:
			case tc.winner == winner:
				t.Errorf("wrong winner: expected = %v, actual = %v", tc.winner, winner)
			}
		})
	}
}
