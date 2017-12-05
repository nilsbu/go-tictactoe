package actor

import (
	"go-tictactoe/mechanics"
	"go-tictactoe/util"
)

type Computer struct {
	Players int
}

func (c *Computer) GetMove(field mechanics.Field) (pos mechanics.Position, err error) {
	return mechanics.Position{0, 0}, util.NewError("Not implemented")
}

func computeOptimalMoveSeq(marks []mechanics.Player,
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
