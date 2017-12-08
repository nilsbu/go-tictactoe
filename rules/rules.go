package rules

import "go-tictactoe/mechanics"

// IsFull checks if a board is full.
func IsFull(b mechanics.Board) bool {
	for _, v := range b.Marks {
		if v == mechanics.Player(0) {
			return false
		}
	}

	return true
}

// GetWinner determines if there a player has won the game.
// NoWinner is returned if this is not the case, otherwise the player's ID is
// returned.
func GetWinner(b mechanics.Board) (id mechanics.Player, hasWinner bool) {
	id, hasWinner = getRowWinner(b)
	if hasWinner {
		return
	}

	id, hasWinner = getColumnWinner(b)
	if hasWinner {
		return
	}

	return getDiagonalWinner(b)
}

func getRowWinner(b mechanics.Board) (id mechanics.Player, hasWinner bool) {
	for y := 0; y < b.Size; y++ {
		if b.Marks[y*b.Size] == 0 {
			continue
		}

		x := 1
		for ; x < b.Size; x++ {
			if b.Marks[y*b.Size+x] != b.Marks[y*b.Size] {
				break
			}
		}
		if x == b.Size {
			return b.Marks[y*b.Size], true
		}
	}

	return 0, false
}

func getColumnWinner(b mechanics.Board) (id mechanics.Player, hasWinner bool) {
	for x := 0; x < b.Size; x++ {
		if b.Marks[x] == 0 {
			continue
		}

		y := 1
		for ; y < b.Size; y++ {
			if b.Marks[y*b.Size+x] != b.Marks[x] {
				break
			}
		}
		if y == b.Size {
			return b.Marks[x], true
		}
	}

	return 0, false
}

func getDiagonalWinner(b mechanics.Board) (id mechanics.Player, hasWinner bool) {
	if b.Marks[0] != 0 {
		for xy := 1; xy < b.Size; xy++ {
			if b.Marks[xy*b.Size+xy] != b.Marks[0] {
				break
			}
			if xy == b.Size-1 {
				return b.Marks[0], true
			}
		}
	}

	if b.Marks[b.Size-1] != 0 {
		for xy := 1; xy < b.Size; xy++ {
			if b.Marks[xy*b.Size-xy+b.Size-1] != b.Marks[b.Size-1] {
				break
			}
			if xy == b.Size-1 {
				return b.Marks[b.Size-1], true
			}
		}
	}

	return 0, false
}
