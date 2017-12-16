package game

import (
	"errors"
	"fmt"

	a "go-tictactoe/actor"
	b "go-tictactoe/board"
	"go-tictactoe/io"
)

// Flow provides the method Loop in which the main loop is contained.
// It requests moves from all players after another and ends when an outcom has
// been achieved.
type Flow interface {
	Loop()
}

// MinPlayers is the minimal number of players, human or not, that are needed
// in a game.
const MinPlayers = 2

// MaxPlayers is the maximum number of players, human or not, that are allowed
// in a game.
const MaxPlayers = 5

// MinBoardSize is the minimal board size required.
// There are this many columns and rows required.
const MinBoardSize = 3

// Game holds the information about the current state of the game.
type Game struct {
	Players []a.Actor
	Board   b.Board

	CurrentPlayer PlayerCounter

	io io.IO
}

// PlayerType differentiates between human and computer players.
type PlayerType int

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

	playerArr := make([]a.Actor, players)

	var i b.Player
	for ; i < b.Player(humanPlayers); i++ {
		playerArr[i] = a.Actor(a.NewHuman())
	}
	for ; i < b.Player(players); i++ {
		playerArr[i] = a.Actor(&a.Computer{ID: i + 1, Players: b.Player(players)})
	}

	return &Game{
		playerArr,
		b.NewData(boardSize),
		PlayerCounter{Next: 1, Total: b.Player(players)},
		io.NewConsole(),
	}, nil
}

// Move adds a players move to the board.
// It is checked if the next move belongs to the player.
func (g *Game) Move(pos b.Position, player b.Player) error {
	if player != g.CurrentPlayer.Next {
		return fmt.Errorf("next move belongs to player %v", g.CurrentPlayer.Next)
	}

	if ok, reason := g.Board.Put(pos, player); !ok {
		return errors.New(reason)
	}

	g.CurrentPlayer.Inc()

	return nil
}

// Loop is the main loop.
func (g *Game) Loop() {
	for {
		g.io.Outf("Your move, player %v:\n", g.CurrentPlayer.Next)
		player := g.Players[g.CurrentPlayer.Next-1]
		pos, err := player.GetMove(g.Board)
		if err != nil {
			break
		}

		err = g.Move(pos, g.CurrentPlayer.Next)
		g.io.OutBoard(g.Board)
		if err != nil {
			g.io.Outln(err)
			continue
		}

		if w, hw := g.Board.(b.Outcome).GetWinner(); hw {
			g.io.Outf("Player %v won, congrats.\n", w)
			g.io.Outln("Congrats.")
			break
		}

		if g.Board.(b.Outcome).IsFull() {
			g.io.Outln("It's a draw.")
			break
		}
	}
}
