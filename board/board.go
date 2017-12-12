package board

import (
	"fmt"
	"strings"
)

// Symbols stores the marks that players make on the board.
// The first one is the mark of an empty board, the subsequent ones belong to
// the players.
var symbols = []string{" ", "x", "o", "8", "v", "^"}

// FIXME symbols should not be limited without checks

// Board provides access to the board that the game is played on.
// Put makes marks at the requested position.
// IsWritable finds out if a position is stil free.
// Get returns a copy of the data.
type Board interface {
	Put(p Position, player Player) (ok bool, reason string)
	IsWritable(p Position) (ok bool, reason string)
	Get() Data
}

// Data is the board the game is played on.
// The size is customizable but the board is always quadratic.
// Marks stores the moves the players have made.
// 0 means the board is empty, all other values represent the player numbers.
type Data struct {
	Marks Marks
	Size  int
}

// Marks stores the players moves on the board.
// The (sequential) array slice is interpreted as a line-wise traversal of the
// board.
type Marks []Player

// Player is the ID of a player or the number of players.
type Player int

// Position is a position on the board.
// The two values are the x- and y-coordinate.
// The top-left corner is (0, 0).
type Position [2]int

// NewPosition creates a Position from an index in a Marks array and the size
// (edge length) of the board.
func NewPosition(index int, size int) Position {
	return Position{index % size, index / size}
}

// ToIndex returns the array index of a Position.
func (p Position) ToIndex(s int) int {
	return p[1]*s + p[0]
}

func (bo Data) String() string {
	// TODO move this to another place
	s := strings.Repeat("-", 2*bo.Size+1) + "\n"

	for y := 0; y < bo.Size; y++ {
		s += "|"

		for x := 0; x < bo.Size; x++ {
			s += fmt.Sprintf("%v|", symbols[bo.Marks[y*bo.Size+x]])
		}

		s += "\n"
		s += strings.Repeat("-", 2*bo.Size+1) + "\n"
	}

	return s
}

// Put makes a mark for a player on the board.
// If the position is not on the board or a mark has already been made at the
// specified position, Put returns an error.
func (bo Data) Put(p Position, player Player) (ok bool, reason string) {
	if ok, reason := bo.IsWritable(p); !ok {
		return ok, reason
	}

	bo.Marks[p.ToIndex(bo.Size)] = player
	return true, ""
}

// IsWritable checks is a position can be written in.
// It has to be within the limits of the board and empty.
// If the position is not writable a reason is given.
func (bo Data) IsWritable(p Position) (ok bool, reason string) {
	if p[0] < 0 || p[0] >= bo.Size || p[1] < 0 || p[1] >= bo.Size {
		return false, fmt.Sprintf("position out of range, board has size %vx%v",
			bo.Size, bo.Size)
	}

	if bo.Marks[p.ToIndex(bo.Size)] != 0 {
		return false, fmt.Sprintf("position is not empty")
	}

	return true, ""
}

// Get returns a copy of the board.
func (bo Data) Get() Data {
	// TODO test
	result := Data{Marks: make(Marks, bo.Size*bo.Size), Size: bo.Size}

	for i, v := range bo.Marks {
		result.Marks[i] = v
	}

	return result
}

// Equal returns true when the Marks in the parameter board are equal to the
// one of the receiver.
func (marks Marks) Equal(other Marks) bool {
	// TODO this function is kinda ugly, since it is only used in tests
	if len(marks) != len(other) {
		return false
	}
	for i := 0; i < len(marks); i++ {
		if marks[i] != other[i] {
			return false
		}
	}
	return true
}
