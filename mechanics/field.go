package mechanics

import (
	"fmt"
	"go-tictactoe/util"
	"strings"
)

var symbols = []string{" ", "x", "o", "8", "v", "^"}

type Field struct {
	Marks Marks
	Size  int
}
type Marks []Player
type Position [2]int
type Player int

func (field Field) String() (s string) {

	s += strings.Repeat("-", 2*field.Size+1) + "\n"

	for y := 0; y < field.Size; y++ {
		s += "|"

		for x := 0; x < field.Size; x++ {
			s += fmt.Sprintf("%v|", symbols[field.Marks[y*field.Size+x]])
		}

		s += "\n"
		s += strings.Repeat("-", 2*field.Size+1) + "\n"
	}

	return
}

func (field Field) Put(pos Position, player Player) error {
	if pos[0] < 0 || pos[0] >= field.Size {
		return util.NewError("x-position out of range, required: 0 <= %v < %v", pos[0], field.Size)
	}
	if pos[1] < 0 || pos[1] >= field.Size {
		return util.NewError("y-position out of range, required: 0 <= %v < %v", pos[1], field.Size)
	}

	p := pos[1]*field.Size + pos[0]
	if field.Marks[p] != 0 {
		return util.NewError("Field already written: position = %v, value = %v", pos, field.Marks[p])
	}

	field.Marks[p] = Player(player + 1)

	return nil
}
