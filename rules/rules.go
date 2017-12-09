package rules

import b "go-tictactoe/board"

// IsFull checks if a board is full.
func IsFull(bo b.Board) bool {
	for _, v := range bo.Marks {
		if v == b.Player(0) {
			return false
		}
	}

	return true
}

// GetWinner determines if there a player has won the game.
// NoWinner is returned if this is not the case, otherwise the player's ID is
// returned.
func GetWinner(bo b.Board) (id b.Player, hasWinner bool) {
	id, hasWinner = getRowWinner(bo)
	if hasWinner {
		return
	}

	id, hasWinner = getColumnWinner(bo)
	if hasWinner {
		return
	}

	return getDiagonalWinner(bo)
}

func getRowWinner(bo b.Board) (id b.Player, hasWinner bool) {
	for y := 0; y < bo.Size; y++ {
		if bo.Marks[y*bo.Size] == 0 {
			continue
		}

		x := 1
		for ; x < bo.Size; x++ {
			if bo.Marks[y*bo.Size+x] != bo.Marks[y*bo.Size] {
				break
			}
		}
		if x == bo.Size {
			return bo.Marks[y*bo.Size], true
		}
	}

	return 0, false
}

func getColumnWinner(bo b.Board) (id b.Player, hasWinner bool) {
	for x := 0; x < bo.Size; x++ {
		if bo.Marks[x] == 0 {
			continue
		}

		y := 1
		for ; y < bo.Size; y++ {
			if bo.Marks[y*bo.Size+x] != bo.Marks[x] {
				break
			}
		}
		if y == bo.Size {
			return bo.Marks[x], true
		}
	}

	return 0, false
}

func getDiagonalWinner(bo b.Board) (id b.Player, hasWinner bool) {
	if bo.Marks[0] != 0 {
		for xy := 1; xy < bo.Size; xy++ {
			if bo.Marks[xy*bo.Size+xy] != bo.Marks[0] {
				break
			}
			if xy == bo.Size-1 {
				return bo.Marks[0], true
			}
		}
	}

	if bo.Marks[bo.Size-1] != 0 {
		for xy := 1; xy < bo.Size; xy++ {
			if bo.Marks[xy*bo.Size-xy+bo.Size-1] != bo.Marks[bo.Size-1] {
				break
			}
			if xy == bo.Size-1 {
				return bo.Marks[bo.Size-1], true
			}
		}
	}

	return 0, false
}
