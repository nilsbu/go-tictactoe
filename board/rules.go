package board

// Outcome is an interface that figures out the outcome of the game based on the
// game's rules.
// IsFinished tells whether the game is finished and if yes, whether it is a
// draw or which player won.
type Outcome interface {
	IsFinished() (finished bool, draw bool, winner Player)
}

type xyAccessor struct {
	major  int
	minor  int
	access func(int, int) int
}

// IsFinished returns if the game is finished.
// A game is finished, if there is no room on the board for a player to win.
// If it is finished, the winner is returned.
// If none exists, draw is true.
func (bo Data) IsFinished() (finished bool, draw bool, winner Player) {
	finished = true

	ds := []xyAccessor{
		// Rows
		{bo.Size, bo.Size, func(maj, min int) int { return maj*bo.Size + min }},
		//Columns
		{bo.Size, bo.Size, func(maj, min int) int { return min*bo.Size + maj }},
		// Main diagonal
		{1, bo.Size, func(maj, min int) int { return min * (bo.Size + 1) }},
		// Antidiagonal
		{1, bo.Size, func(maj, min int) int { return (min+1)*bo.Size - min - 1 }},
	}

	for _, d := range ds {
		for maj := 0; maj < d.major; maj++ {
			blocked, hasWinner, id := isMinorFinished(bo, d, maj)
			if hasWinner {
				return true, false, id
			}
			finished = finished && blocked
		}
	}

	return finished, finished, 0
}

func isMinorFinished(bo Data, d xyAccessor, maj int) (blocked bool, hasWinner bool,
	id Player) {
	hasWinner = true

	for min := 0; min < d.minor; min++ {
		c := bo.Marks[d.access(maj, min)]

		if id == 0 {
			id = c
		} else if c != id && c != 0 {
			return true, false, 0
		}

		if c != id || c == 0 {
			hasWinner = false
		}
	}

	return
}
