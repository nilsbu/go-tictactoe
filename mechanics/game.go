package mechanics

import "fmt"

const (
	human PlayerType = iota
	computer
)

const MinPlayers = 2
const MinFieldSize = 3

// Game holds the information about the current state of the game.
type Game struct {
	Players []PlayerType
	Field   Field

	NextPlayer Player
}

// PlayerType differentiates between human and computer players.
type PlayerType int

// NewGame initializes a Game.
// An error is thrown when fewer than MinPlayers players are requested, when the
// board is smaller MinFieldSize fields across or when the number of human
// players is larger than the total number.
func NewGame(fieldSize, players, humanPlayers int) (*Game, error) {
	if players < MinPlayers {
		return nil, fmt.Errorf("Too few players: %v < %v", players, MinPlayers)
	}
	if fieldSize < MinFieldSize {
		return nil, fmt.Errorf("Field too small: %v < %v", fieldSize, MinFieldSize)
	}
	if players < humanPlayers {
		return nil, fmt.Errorf("More humans than players: %v > %v", humanPlayers, players)
	}

	playerArr := make([]PlayerType, players)

	i := 0
	for ; i < humanPlayers; i++ {
		playerArr[i] = human
	}
	for ; i < players; i++ {
		playerArr[i] = computer
	}

	return &Game{
		playerArr,
		Field{make(Marks, fieldSize*fieldSize), fieldSize},
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
	g.Field.Put(pos, player)
	g.NextPlayer = Player(int(g.NextPlayer+1) % len(g.Players))

	return nil
}
