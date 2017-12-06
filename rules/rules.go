package rules

import "go-tictactoe/mechanics"

// GetWinner determines if there a player has won the game.
// NoWinner is returned if this is not the case, otherwise the player's ID is
// returned.
func GetWinner(f mechanics.Field) (id mechanics.Player, hasWinner bool) {
	id, hasWinner = getRowWinner(f)
	if hasWinner {
		return
	}

	id, hasWinner = getColumnWinner(f)
	if hasWinner {
		return
	}

	return getDiagonalWinner(f)
}

func getRowWinner(field mechanics.Field) (id mechanics.Player, hasWinner bool) {
	for y := 0; y < field.Size; y++ {
		if field.Marks[y*field.Size] == 0 {
			continue
		}

		x := 1
		for ; x < field.Size; x++ {
			if field.Marks[y*field.Size+x] != field.Marks[y*field.Size] {
				break
			}
		}
		if x == field.Size {
			return field.Marks[y*field.Size], true
		}
	}

	return 0, false
}

func getColumnWinner(field mechanics.Field) (id mechanics.Player, hasWinner bool) {
	for x := 0; x < field.Size; x++ {
		if field.Marks[x] == 0 {
			continue
		}

		y := 1
		for ; y < field.Size; y++ {
			if field.Marks[y*field.Size+x] != field.Marks[x] {
				break
			}
		}
		if y == field.Size {
			return field.Marks[x], true
		}
	}

	return 0, false
}

func getDiagonalWinner(field mechanics.Field) (id mechanics.Player, hasWinner bool) {
	if field.Marks[0] != 0 {
		for xy := 1; xy < field.Size; xy++ {
			if field.Marks[xy*field.Size+xy] != field.Marks[0] {
				break
			}
			if xy == field.Size-1 {
				return field.Marks[0], true
			}
		}
	}

	if field.Marks[field.Size-1] != 0 {
		for xy := 1; xy < field.Size; xy++ {
			if field.Marks[xy*field.Size-xy+field.Size-1] != field.Marks[field.Size-1] {
				break
			}
			if xy == field.Size-1 {
				return field.Marks[field.Size-1], true
			}
		}
	}

	return 0, false
}
