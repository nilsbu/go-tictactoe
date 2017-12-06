package mechanics

import (
	"fmt"
	"strings"
)

// Symbols stores the marks that players make on the board.
// The first one is the mark of an empty board, the subsequent ones belong to
// the players.
var symbols = []string{" ", "x", "o", "8", "v", "^"}

// Board is the board the game is played on.
// The size is customizable but the board is always quadratic.
// Marks stores the moves the players have made.
// 0 means the board is empty, all other values represent the player numbers.
type Board struct {
	Marks Marks
	Size  int
}

// Marks stores the players moves on the board.
// The (sequential) array slice is interpreted as a line-wise traversal of the
// board.
type Marks []Player

// Position is a position on the board.
// The two values are the x- and y-coordinate.
// The top-left corner is (0, 0).
type Position [2]int

// Player is the ID of a player.
type Player int

func (b Board) String() string {

	s := strings.Repeat("-", 2*b.Size+1) + "\n"

	for y := 0; y < b.Size; y++ {
		s += "|"

		for x := 0; x < b.Size; x++ {
			s += fmt.Sprintf("%v|", symbols[b.Marks[y*b.Size+x]])
		}

		s += "\n"
		s += strings.Repeat("-", 2*b.Size+1) + "\n"
	}

	return s
}

// Put makes a mark for a player on the board.
// If the position is not on the board or a mark has already been made at the
// specified position, Put returns an error.
func (b Board) Put(pos Position, player Player) error {
	if pos[0] < 0 || pos[0] >= b.Size {
		return fmt.Errorf("x-position out of range, required: 0 <= %v < %v",
			pos[0], b.Size)
	}
	if pos[1] < 0 || pos[1] >= b.Size {
		return fmt.Errorf("y-position out of range, required: 0 <= %v < %v",
			pos[1], b.Size)
	}

	p := pos[1]*b.Size + pos[0]
	if b.Marks[p] != 0 {
		return fmt.Errorf("board already written: position = %v, value = %v",
			pos, b.Marks[p])
	}

	b.Marks[p] = Player(player + 1)

	return nil
}
