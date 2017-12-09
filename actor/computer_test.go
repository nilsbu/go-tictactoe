package actor

import (
	"fmt"
	"math"
	"testing"

	m "go-tictactoe/mechanics"
	"go-tictactoe/test"
)

const noWinner m.Player = -1

var testCases = []struct {
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

func TestComputerGetMove(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			c := Computer{ID: tc.id, Players: tc.players}
			s := int(math.Sqrt(float64(len(tc.marks))))
			b := m.Board{Marks: tc.marks, Size: s}
			switch pos, err := c.GetMove(b); false {
			case err == nil:
				t.Errorf("must never return an error")
			case isIndexInList(pos.ToIndex(s), tc.idxs):
				t.Errorf("%v (= %v) must be in %v", pos, pos.ToIndex(s),
					tc.idxs)
			}
		})
	}
}

func TestComputeOptimalMovePar(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			s := int(math.Sqrt(float64(len(tc.marks))))
			marks := make(m.Marks, len(tc.marks))
			copy(marks, tc.marks)
			b := m.Board{Marks: marks, Size: s}
			switch p := computeOptimalMovePar(b, tc.id, tc.players); false {
			case isIndexInList(p, tc.idxs):
				t.Errorf("%v must be in %v", p, tc.idxs)
			}
		})
	}
}

func TestComputeOptimalMoveSeq(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			s := int(math.Sqrt(float64(len(tc.marks))))
			marks := make(m.Marks, len(tc.marks))
			copy(marks, tc.marks)
			b := m.Board{Marks: marks, Size: s}
			p, w, hw := computeOptimalMoveSeq(b, tc.id, tc.players)
			switch false {
			case isBoardUnchanged(tc.marks, b.Marks):
				t.Errorf("board changed")
			case test.Cond(!hw, w == 0):
				t.Errorf("hasWinner = false but winner = %v, must be 0", w)
			case test.Cond(tc.winner != noWinner, hw):
				t.Errorf("winner was expected but none was returned")
			case test.Cond(tc.winner == noWinner, !hw):
				t.Errorf("no winner was expected but %v was returned", w)
			case test.Cond(tc.winner != noWinner, tc.winner == w):
				t.Errorf("expected = %v, actual = %v", tc.winner, w)
			case isIndexInList(p, tc.idxs):
				t.Errorf("%v must be in %v", p, tc.idxs)
			}
		})
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

var benchMarks = []m.Marks{
	m.Marks{2, 1, 2, 1, 1, 2, 1, 2, 0},
	m.Marks{2, 1, 2, 1, 1, 0, 1, 2, 0},
	m.Marks{2, 1, 2, 0, 1, 0, 1, 2, 0},
	m.Marks{2, 1, 2, 0, 1, 0, 1, 0, 0},
	m.Marks{2, 0, 2, 0, 1, 0, 1, 0, 0},
	m.Marks{2, 0, 0, 0, 1, 0, 1, 0, 0},
	m.Marks{2, 0, 0, 0, 1, 0, 0, 0, 0},
	m.Marks{0, 0, 0, 0, 1, 0, 0, 0, 0},
	m.Marks{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func bench(b *testing.B, i int, parallel bool) {
	c := Computer{ID: m.Player((i+1)%2 + 1), Players: 2}
	bo := m.Board{Marks: benchMarks[i-1], Size: 3}

	if parallel {
		for n := 0; n < b.N; n++ {
			c.GetMove(bo)
		}
	} else {
		for n := 0; n < b.N; n++ {
			c.getMoveSequential(bo)
		}
	}
}

func BenchmarkComputerGetMove1S(b *testing.B) { bench(b, 1, false) }
func BenchmarkComputerGetMove2S(b *testing.B) { bench(b, 2, false) }
func BenchmarkComputerGetMove3S(b *testing.B) { bench(b, 3, false) }
func BenchmarkComputerGetMove4S(b *testing.B) { bench(b, 4, false) }
func BenchmarkComputerGetMove5S(b *testing.B) { bench(b, 5, false) }
func BenchmarkComputerGetMove6S(b *testing.B) { bench(b, 6, false) }
func BenchmarkComputerGetMove7S(b *testing.B) { bench(b, 7, false) }
func BenchmarkComputerGetMove8S(b *testing.B) { bench(b, 8, false) }
func BenchmarkComputerGetMove9S(b *testing.B) { bench(b, 9, false) }
func BenchmarkComputerGetMove1P(b *testing.B) { bench(b, 1, true) }
func BenchmarkComputerGetMove2P(b *testing.B) { bench(b, 2, true) }
func BenchmarkComputerGetMove3P(b *testing.B) { bench(b, 3, true) }
func BenchmarkComputerGetMove4P(b *testing.B) { bench(b, 4, true) }
func BenchmarkComputerGetMove5P(b *testing.B) { bench(b, 5, true) }
func BenchmarkComputerGetMove6P(b *testing.B) { bench(b, 6, true) }
func BenchmarkComputerGetMove7P(b *testing.B) { bench(b, 7, true) }
func BenchmarkComputerGetMove8P(b *testing.B) { bench(b, 8, true) }
func BenchmarkComputerGetMove9P(b *testing.B) { bench(b, 9, true) }
