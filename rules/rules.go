package rules

import "go-tictactoe/mechanics"

// NoWinner is a return value for when there is no winner.
const NoWinner = -1

// GetWinner determines if there a player has won the game.
// NoWinner is returned if this is not the case, otherwise the player's ID is
// returned.
func GetWinner(f mechanics.Field) mechanics.Player {
	winner := getRowWinner(f)
	// TODO A second return variable should be used to check if there is a winner.
	if winner != NoWinner {
		return winner
	}

	winner = getColumnWinner(f)
	if winner != NoWinner {
		return winner
	}

	return getDiagonalWinner(f)
}

func getRowWinner(field mechanics.Field) (winner mechanics.Player) {
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
			return field.Marks[y*field.Size]
		}
	}

	return NoWinner
}

func getColumnWinner(field mechanics.Field) (winner mechanics.Player) {
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
			return field.Marks[x]
		}
	}

	return NoWinner
}

func getDiagonalWinner(field mechanics.Field) (winner mechanics.Player) {
	if field.Marks[0] != 0 {
		for xy := 1; xy < field.Size; xy++ {
			if field.Marks[xy*field.Size+xy] != field.Marks[0] {
				break
			}
			if xy == field.Size-1 {
				return field.Marks[0]
			}
		}
	}

	if field.Marks[field.Size-1] != 0 {
		for xy := 1; xy < field.Size; xy++ {
			if field.Marks[xy*field.Size-xy+field.Size-1] != field.Marks[field.Size-1] {
				break
			}
			if xy == field.Size-1 {
				return field.Marks[field.Size-1]
			}
		}
	}

	return NoWinner
}
