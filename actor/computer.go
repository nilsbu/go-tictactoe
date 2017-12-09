package actor

import (
	m "go-tictactoe/mechanics"
	r "go-tictactoe/rules"
)

// Computer represents a computer player.
// It implements a function to make a move.
type Computer struct {
	ID      m.Player
	Players m.Player
}

// GetMove makes the next move for the computer player calling it.
func (c *Computer) GetMove(b m.Board) (m.Position, error) {
	p, _, _ := computeOptimalMoveSeq(b, c.ID, c.Players)
	return m.NewPosition(p, b.Size), nil
}

// computeOptimalMoveSeq finds the optimal move for the player.
func computeOptimalMoveSeq(b m.Board, current m.Player, numPlayers m.Player) (
	pos int, winner m.Player, hasWinner bool) {

	winner = 0
	hasWinner = true
	for p := 0; p < len(b.Marks); p++ {
		if b.Marks[p] > 0 {
			continue
		}

		tmpWinner, tmpHas := attempt(b, p, current, numPlayers)

		switch {
		case tmpWinner == current:
			return p, current, true
		case !tmpHas:
			pos = p
			winner = 0
			hasWinner = false
		case hasWinner:
			pos = p
			winner = tmpWinner
		}
	}

	return
}

func attempt(b m.Board, p int, current m.Player, numPlayers m.Player) (
	winner m.Player, hasWinner bool) {

	defer func() { b.Marks[p] = 0 }()
	b.Marks[p] = current

	if winner, hasWinner = r.GetWinner(b); hasWinner {
		return winner, hasWinner
	}

	if r.IsFull(b) {
		return m.Player(0), false
	}

	nextPlayer := m.Player((current % numPlayers) + 1)
	_, winner, hasWinner = computeOptimalMoveSeq(b, nextPlayer, numPlayers)

	return
}
