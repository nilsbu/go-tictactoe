package game

import (
	"errors"
	"fmt"

	b "go-tictactoe/board"
)

// MinPlayers is the minimal number of players, human or not, that are needed
// for a game.
const MinPlayers = 2

// MinBoardSize is the minimal board size required.
// There are this many columns and rows required.
const MinBoardSize = 3

// Game holds the information about the current state of the game.
type Game struct {
	Players []PlayerType
	Board   b.Board

	CurrentPlayer PlayerCounter
}

// PlayerType differentiates between human and computer players.
type PlayerType int

// Human and Computer are PlayerTypes that donote human and computer players
// respectively.
const (
	Human PlayerType = iota
	Computer
)

// PlayerCounter is a counter that provides the ID of the next plauer.
type PlayerCounter struct {
	// Next is the ID of the next player.
	Next b.Player
	// Total is the number of players.
	Total b.Player
}

// Inc moves the counter on to the next player.
// If the last player was reached it starts with the first player.
func (pc *PlayerCounter) Inc() {
	pc.Next = pc.Next%pc.Total + 1
}

// NewGame initializes a Game.
// An error is thrown when fewer than MinPlayers players are requested, when the
// board is smaller MinBoardSize boards across or when the number of human
// players is larger than the total number.
func NewGame(boardSize, players, humanPlayers int) (*Game, error) {
	if players < MinPlayers {
		return nil, fmt.Errorf("too few players: %v < %v", players, MinPlayers)
	}
	if boardSize < MinBoardSize {
		return nil, fmt.Errorf("board too small: %v < %v", boardSize, MinBoardSize)
	}
	if players < humanPlayers {
		return nil, fmt.Errorf("more humans than players: %v > %v",
			humanPlayers, players)
	}

	playerArr := make([]PlayerType, players)

	i := 0
	for ; i < humanPlayers; i++ {
		playerArr[i] = Human
	}
	for ; i < players; i++ {
		playerArr[i] = Computer
	}

	return &Game{
		playerArr,
		b.Board{Marks: make(b.Marks, boardSize*boardSize), Size: boardSize},
		PlayerCounter{Next: 1, Total: b.Player(players)},
	}, nil
}

// Move adds a players move to the board.
// It is checked if the next move belongs to the player.
func (g *Game) Move(pos b.Position, player b.Player) error {
	// TODO What about false moves?
	if player != g.CurrentPlayer.Next {
		return fmt.Errorf("next move belongs to player %v", g.CurrentPlayer.Next)
	}

	if ok, reason := g.Board.IsWritable(pos); !ok {
		return errors.New(reason)
	}

	g.Board.Put(pos, player)
	g.CurrentPlayer.Inc()

	return nil
}
