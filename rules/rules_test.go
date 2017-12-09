package rules

import (
	"fmt"
	"testing"

	b "go-tictactoe/board"
	"go-tictactoe/test"
)

func TestIsFull(t *testing.T) {
	testCases := []struct {
		marks b.Marks
		size  int
		fin   bool
	}{
		{b.Marks{0, 0, 0, 0}, 2, false},
		{b.Marks{2, 1, 2, 0}, 2, false},
		{b.Marks{2, 1, 2, 3}, 2, true},
		{b.Marks{2, 1, 2, 3, 1, 1, 2, 3, 2}, 3, true},
		{b.Marks{0, 1, 2, 3, 1, 1, 2, 3, 2}, 3, false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			bo := b.Board{Marks: tc.marks, Size: tc.size}
			switch fin := IsFull(bo); false {
			case tc.fin == fin:
				t.Errorf("expected = %v, actual %v", tc.fin, fin)
			}
		})
	}
}

func TestGetWinner(t *testing.T) {
	const noWinner b.Player = -1

	testCases := []struct {
		marks  b.Marks
		size   int
		winner b.Player
	}{
		// No winner
		{b.Marks{0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, noWinner},
		{b.Marks{2, 1, 2, 1, 1, 2, 1, 2, 1}, 3, noWinner},
		// Winner in row
		{b.Marks{1, 1, 1, 2, 2, 1, 1, 2, 2}, 3, 1},
		{b.Marks{1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 1, 1, 2, 1, 1, 2}, 4, 2},
		// Winner in column
		{b.Marks{1, 2, 3, 2, 1, 3, 1, 2, 3}, 3, 3},
		{b.Marks{1, 2, 2, 1, 1, 2, 2, 1, 1}, 3, 1},
		// Winner in diagonal
		{b.Marks{1, 2, 2, 1, 1, 2, 2, 2, 1}, 3, 1},
		{b.Marks{0, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 0}, 4, 2},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			bo := b.Board{Marks: tc.marks, Size: tc.size}
			switch winner, hasWinner := GetWinner(bo); false {
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
