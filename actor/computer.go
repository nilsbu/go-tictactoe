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

func (c *Computer) getMoveSequential(b m.Board) (m.Position, error) {
	p, _, _ := computeOptimalMoveSeq(b, c.ID, c.Players)
	return m.NewPosition(p, b.Size), nil
}

// GetMove makes the next move for the computer player calling it.
func (c *Computer) GetMove(b m.Board) (m.Position, error) {
	p := computeOptimalMovePar(b, c.ID, c.Players)
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

const (
	nop = iota
	blocked
	loss
	draw
	win
)

func computeOptimalMovePar(b m.Board, current m.Player, numPlayers m.Player) (
	pos int) {

	type answer struct {
		v int
		p int
	}

	wait := make(chan answer, len(b.Marks))

	for p := 0; p < len(b.Marks); p++ {
		if b.Marks[p] > 0 {
			wait <- answer{v: blocked, p: p}
			continue
		}

		go func(i int, cur m.Player, numP m.Player) {
			mcop := make(m.Marks, len(b.Marks))
			copy(mcop, b.Marks)
			bcop := m.Board{Marks: mcop, Size: b.Size}
			tmpWinner, tmpHas := attempt(bcop, i, cur, numP)

			switch {
			case tmpWinner == current:
				wait <- answer{v: win, p: i}
			case !tmpHas:
				wait <- answer{v: draw, p: i}
			default:
				wait <- answer{v: loss, p: i}
			}
		}(p, current, numPlayers)
	}

	var res = blocked
	for p := 0; p < len(b.Marks); p++ {
		a := <-wait

		if res < a.v {
			pos = a.p
			res = a.v
		}
	}
	return
}
