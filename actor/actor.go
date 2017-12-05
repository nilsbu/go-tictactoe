package actor

import "go-tictactoe/mechanics"

type Actor interface {
	GetMove(field mechanics.Field) (mechanics.Position, error)
}
