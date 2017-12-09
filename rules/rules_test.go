package rules

import (
	"fmt"
	"testing"

	m "go-tictactoe/mechanics"
	"go-tictactoe/test"
)

func TestIsFull(t *testing.T) {
	testCases := []struct {
		marks m.Marks
		size  int
		fin   bool
	}{
		{m.Marks{0, 0, 0, 0}, 2, false},
		{m.Marks{2, 1, 2, 0}, 2, false},
		{m.Marks{2, 1, 2, 3}, 2, true},
		{m.Marks{2, 1, 2, 3, 1, 1, 2, 3, 2}, 3, true},
		{m.Marks{0, 1, 2, 3, 1, 1, 2, 3, 2}, 3, false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			b := m.Board{Marks: tc.marks, Size: tc.size}
			switch fin := IsFull(b); false {
			case tc.fin == fin:
				t.Errorf("expected = %v, actual %v", tc.fin, fin)
			}
		})
	}
}

func TestGetWinner(t *testing.T) {
	const noWinner m.Player = -1

	testCases := []struct {
		marks  m.Marks
		size   int
		winner m.Player
	}{
		// No winner
		{m.Marks{0, 0, 0, 0, 0, 0, 0, 0, 0}, 3, noWinner},
		{m.Marks{2, 1, 2, 1, 1, 2, 1, 2, 1}, 3, noWinner},
		// Winner in row
		{m.Marks{1, 1, 1, 2, 2, 1, 1, 2, 2}, 3, 1},
		{m.Marks{1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 1, 1, 2, 1, 1, 2}, 4, 2},
		// Winner in column
		{m.Marks{1, 2, 3, 2, 1, 3, 1, 2, 3}, 3, 3},
		{m.Marks{1, 2, 2, 1, 1, 2, 2, 1, 1}, 3, 1},
		// Winner in diagonal
		{m.Marks{1, 2, 2, 1, 1, 2, 2, 2, 1}, 3, 1},
		{m.Marks{0, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 0}, 4, 2},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			b := m.Board{Marks: tc.marks, Size: tc.size}
			switch winner, hasWinner := GetWinner(b); false {
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
