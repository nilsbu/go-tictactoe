package actor

import (
	"go-tictactoe/mechanics"
	"go-tictactoe/util"
)

// Computer represents a computer player.
// It implements a function to make a move.
type Computer struct {
	Players int
}

// GetMove makes the next move for the computer player calling it.
func (c *Computer) GetMove(field mechanics.Field) (mechanics.Position, error) {
	return mechanics.Position{0, 0}, util.NewError("Not implemented")
}

// computeOptimalMoveSeq finds the optimal move for the player.
func computeOptimalMoveSeq(marks []mechanics.Player,
	// TODO unfinished, doesn't work
	current mechanics.Player,
	numPlayers int) (pos int, winner mechanics.Player) {

	winner = -1
	for p := 0; p < len(marks); p++ {
		if marks[p] > 0 {
			continue
		}

		marks[p] = current
		nextPlayer := mechanics.Player((int(current) + 1) % numPlayers)
		_, res := computeOptimalMoveSeq(marks, nextPlayer, numPlayers)
		marks[p] = 0

		if winner < res {
			winner = res
			pos = p
		}
	}

	return
}
