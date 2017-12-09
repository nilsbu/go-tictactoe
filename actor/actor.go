package actor

import m "go-tictactoe/mechanics"

// Actor is an interface that represents a player in a game.
// It can be a human or computer.
//
// GetMove provides the next move the player makes.
// When it returns an error, the game is supposed to be aborted.
type Actor interface {
	GetMove(b m.Board) (m.Position, error)
}
