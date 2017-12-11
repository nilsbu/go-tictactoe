package actor

import (
	"fmt"
	"math"
	"testing"

	b "go-tictactoe/board"
	"go-tictactoe/test"
)

const noWinner b.Player = -1

var testCases = []struct {
	players b.Player
	id      b.Player
	marks   b.Marks
	idxs    []int
	winner  b.Player
}{
	{2, 2, b.Marks{2, 0, 2, 0, 1, 0, 1, 1, 0}, []int{1}, 2},
	{2, 1, b.Marks{0, 0, 0, 2, 1, 0, 0, 0, 0}, []int{0, 1, 2, 6, 7, 8}, 1},
	{2, 1, b.Marks{0, 0, 0, 1, 2, 0, 0, 0, 0}, []int{0, 2, 6, 8}, noWinner},
	{2, 2, b.Marks{0, 0, 0, 1, 2, 0, 1, 0, 0}, []int{0}, noWinner},
	{2, 1, b.Marks{2, 0, 0, 1, 2, 0, 1, 0, 0}, []int{8}, noWinner},
	{2, 2, b.Marks{2, 0, 0, 1, 2, 0, 1, 0, 1}, []int{7}, noWinner},
	{2, 1, b.Marks{2, 0, 0, 1, 2, 0, 1, 2, 1}, []int{1}, noWinner},
	{2, 2, b.Marks{0, 0, 1, 1, 2, 0, 0, 0, 0}, []int{0, 1, 6, 7}, noWinner},
	{2, 1, b.Marks{0, 0, 1, 1, 2, 0, 0, 2, 0}, []int{1}, noWinner},
	{2, 2, b.Marks{1, 0, 1, 1, 2, 0, 0, 2, 0}, []int{1}, 2},
	{2, 2, b.Marks{1, 0, 0, 2, 1, 0, 0, 0, 0}, []int{1, 2, 5, 6, 7, 8}, 1},
	{2, 2, b.Marks{2, 0, 0, 0, 1, 0, 1, 0, 0}, []int{2}, noWinner},
}

func TestComputerGetMove(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("#%v: %v", i, tc.marks), func(t *testing.T) {
			c := Computer{ID: tc.id, Players: tc.players}
			s := int(math.Sqrt(float64(len(tc.marks))))
			bo := b.Board{Marks: tc.marks, Size: s}
			switch pos, err := c.GetMove(bo); false {
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
			marks := make(b.Marks, len(tc.marks))
			copy(marks, tc.marks)
			bo := b.Board{Marks: marks, Size: s}
			switch p := computeOptimalMovePar(bo, tc.id, tc.players); false {
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
			marks := make(b.Marks, len(tc.marks))
			copy(marks, tc.marks)
			bo := b.Board{Marks: marks, Size: s}
			p, w, hw := computeOptimalMoveSeq(bo, tc.id, tc.players)
			switch false {
			case isBoardUnchanged(tc.marks, bo.Marks):
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

func isBoardUnchanged(a, bo b.Marks) bool {
	for i := range a {
		if a[i] != bo[i] {
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

var benchMarks = []b.Marks{
	b.Marks{2, 1, 2, 1, 1, 2, 1, 2, 0},
	b.Marks{2, 1, 2, 1, 1, 0, 1, 2, 0},
	b.Marks{2, 1, 2, 0, 1, 0, 1, 2, 0},
	b.Marks{2, 1, 2, 0, 1, 0, 1, 0, 0},
	b.Marks{2, 0, 2, 0, 1, 0, 1, 0, 0},
	b.Marks{2, 0, 0, 0, 1, 0, 1, 0, 0},
	b.Marks{2, 0, 0, 0, 1, 0, 0, 0, 0},
	b.Marks{0, 0, 0, 0, 1, 0, 0, 0, 0},
	b.Marks{0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func bench(bm *testing.B, i int, parallel bool) {
	c := Computer{ID: b.Player((i+1)%2 + 1), Players: 2}
	bo := b.Board{Marks: benchMarks[i-1], Size: 3}

	if parallel {
		for n := 0; n < bm.N; n++ {
			c.GetMove(bo)
		}
	} else {
		for n := 0; n < bm.N; n++ {
			c.getMoveSequential(bo)
		}
	}
}

func (c *Computer) getMoveSequential(bo b.Board) (b.Position, error) {
	p, _, _ := computeOptimalMoveSeq(bo, c.ID, c.Players)
	return b.NewPosition(p, bo.Size), nil
}

// TODO put in one function
func BenchmarkComputerGetMove1S(bm *testing.B) { bench(bm, 1, false) }
func BenchmarkComputerGetMove2S(bm *testing.B) { bench(bm, 2, false) }
func BenchmarkComputerGetMove3S(bm *testing.B) { bench(bm, 3, false) }
func BenchmarkComputerGetMove4S(bm *testing.B) { bench(bm, 4, false) }
func BenchmarkComputerGetMove5S(bm *testing.B) { bench(bm, 5, false) }
func BenchmarkComputerGetMove6S(bm *testing.B) { bench(bm, 6, false) }
func BenchmarkComputerGetMove7S(bm *testing.B) { bench(bm, 7, false) }
func BenchmarkComputerGetMove8S(bm *testing.B) { bench(bm, 8, false) }
func BenchmarkComputerGetMove9S(bm *testing.B) { bench(bm, 9, false) }
func BenchmarkComputerGetMove1P(bm *testing.B) { bench(bm, 1, true) }
func BenchmarkComputerGetMove2P(bm *testing.B) { bench(bm, 2, true) }
func BenchmarkComputerGetMove3P(bm *testing.B) { bench(bm, 3, true) }
func BenchmarkComputerGetMove4P(bm *testing.B) { bench(bm, 4, true) }
func BenchmarkComputerGetMove5P(bm *testing.B) { bench(bm, 5, true) }
func BenchmarkComputerGetMove6P(bm *testing.B) { bench(bm, 6, true) }
func BenchmarkComputerGetMove7P(bm *testing.B) { bench(bm, 7, true) }
func BenchmarkComputerGetMove8P(bm *testing.B) { bench(bm, 8, true) }
func BenchmarkComputerGetMove9P(bm *testing.B) { bench(bm, 9, true) }
