package actor

import (
	b "tictactoe/board"
)

const maxDepth = 6

// Computer represents a computer player.
// It implements a function to make a move.
type Computer struct {
	ID      b.Player
	Players b.Player
}

// GetMove makes the next move for the computer player calling it.
func (c *Computer) GetMove(bo b.Board) (b.Position, error) {
	d := bo.Get()
	p := computeOptimalMovePar(d, c.ID, c.Players)
	return b.NewPosition(p, d.Size), nil
}

// computeOptimalMoveSeq finds the optimal move for the player.
func computeOptimalMoveSeq(bo b.Data, current b.Player, numPlayers b.Player,
	rec int) (
	pos int, winner b.Player, hasWinner bool) {

	winner = 0
	hasWinner = true
	for p := 0; p < len(bo.Marks); p++ {
		if bo.Marks[p] > 0 {
			continue
		}

		tmpWinner, tmpHas := attempt(bo, p, current, numPlayers, rec)

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

func attempt(bo b.Data, p int, current b.Player, numPlayers b.Player, rec int) (
	winner b.Player, hasWinner bool) {

	defer func() { bo.Marks[p] = 0 }()
	bo.Marks[p] = current

	finished, draw, winner := bo.IsFinished()
	if finished {
		return winner, !draw
	} else if rec == 0 {
		// Don't continue here, report draw
		return b.Player(0), false
	}

	nextPlayer := b.Player((current % numPlayers) + 1)
	_, winner, hasWinner = computeOptimalMoveSeq(bo, nextPlayer, numPlayers,
		rec-1)

	return
}

const (
	nop = iota
	blocked
	loss
	draw
	win
)

func computeOptimalMovePar(bo b.Data, current b.Player, numPlayers b.Player) (
	pos int) {

	type answer struct {
		v int
		p int
	}

	wait := make(chan answer, len(bo.Marks))

	for p := 0; p < len(bo.Marks); p++ {
		if bo.Marks[p] > 0 {
			wait <- answer{v: blocked, p: p}
			continue
		}

		go func(i int, cur b.Player, numP b.Player) {
			mcop := make(b.Marks, len(bo.Marks))
			copy(mcop, bo.Marks)
			bcop := b.Data{Marks: mcop, Size: bo.Size}
			tmpWinner, tmpHas := attempt(bcop, i, cur, numP, maxDepth)

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
	for p := 0; p < len(bo.Marks); p++ {
		a := <-wait

		if res < a.v {
			pos = a.p
			res = a.v
		}
	}
	return
}
