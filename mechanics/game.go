package mechanics

import "go-tictactoe/util"

const (
	human PlayerType = iota
	computer
)

const minPlayers = 2
const minFieldSize = 3

type Game struct {
	Players []PlayerType
	Field   Field

	NextPlayer Player
}

type PlayerType int

func NewGame(fieldSize, players, humanPlayers int) (game *Game, err error) {
	if players < minPlayers {
		return nil, util.NewError("Too few players: %v < %v", players, minPlayers)
	}
	if fieldSize < minFieldSize {
		return nil, util.NewError("Field too small: %v < %v", fieldSize, minFieldSize)
	}
	if players < humanPlayers {
		return nil, util.NewError("More humans than players: %v > %v", humanPlayers, players)
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

func (game *Game) Move(pos Position, player Player) error {
	if player != game.NextPlayer {
		return util.NewError("Next move belongs to player %v", game.NextPlayer)
	}

	game.Field.Put(pos, player)
	game.NextPlayer = Player(int(game.NextPlayer+1) % len(game.Players))

	return nil
}
