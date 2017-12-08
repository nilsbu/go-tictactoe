package mechanics

import "fmt"

// MinPlayers is the minimal number of players, human or not, that are needed
// for a game.
const MinPlayers = 2

// MinBoardSize is the minimal board size required.
// There are this many columns and rows required.
const MinBoardSize = 3

// Game holds the information about the current state of the game.
type Game struct {
	Players []PlayerType
	Board   Board

	NextPlayer Player
}

// PlayerType differentiates between human and computer players.
type PlayerType int

// Human and Computer are PlayerTypes that donote human and computer players
// respectively.
const (
	Human PlayerType = iota
	Computer
)

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
		Board{make(Marks, boardSize*boardSize), boardSize},
		0,
	}, nil
}

// Move adds a players move to the board.
// It is checked if the next move belongs to the player.
func (g *Game) Move(pos Position, player Player) error {
	// TODO What about false moves?
	if player != g.NextPlayer {
		return fmt.Errorf("Next move belongs to player %v", g.NextPlayer)
	}

	// FIXME not checked for errors
	g.Board.Put(pos, player)
	g.NextPlayer = Player(int(g.NextPlayer+1) % len(g.Players))

	return nil
}
