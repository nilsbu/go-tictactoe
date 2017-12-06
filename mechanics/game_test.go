package mechanics

import (
	"errors"
	"testing"
)

var errAny = errors.New("Some Error")

func TestNewGame(t *testing.T) {
	tables := []struct {
		boardSize    int
		humanPlayers int
		players      []PlayerType
		err          error
	}{
		{3, 2, []PlayerType{human, human}, nil},
		{4, 1, []PlayerType{human, computer}, nil},
		{5, 0, []PlayerType{computer, computer, computer}, nil},
		{3, 1, []PlayerType{human}, errAny},
		{2, 1, []PlayerType{human, computer}, errAny},
		{3, 3, []PlayerType{human, human}, errAny},
	}

	for _, table := range tables {
		game, err := NewGame(table.boardSize, len(table.players), table.humanPlayers)
		if (err == nil) != (table.err == nil) {
			t.Errorf("unexpected error behavior: expected = \"%v\", actual = \"%v\"",
				table.err, err)
			continue
		}
		if err != nil {
			continue
		}
		if len(game.Players) != len(table.players) {
			t.Errorf("number of players: expected = \"%v\", actual = \"%v\"",
				len(table.players), len(game.Players))
		}
		for i := 0; i < len(game.Players); i++ {
			if game.Players[i] != table.players[i] {
				t.Errorf("player setup: expected = \"%v\", actual \"%v\"",
					table.players, game.Players)
				break
			}
		}
		if len(game.Board.Marks) != table.boardSize*table.boardSize {
			t.Errorf("marks size: expected = \"%v\", actual = \"%v\"",
				table.boardSize*table.boardSize, len(game.Board.Marks))
		}
		if game.Board.Size != table.boardSize {
			t.Errorf("board size: expected = \"%v\", actual = \"%v\"",
				table.boardSize, game.Board.Size)
		}
		if game.NextPlayer != 0 {
			t.Errorf("next player: expected = \"0\", actual = \"%v\"",
				game.NextPlayer)
		}
	}
}

func TestGame_Move2(t *testing.T) {
	tables := []struct {
		pos        Position
		playerPre  Player
		playerPost Player
		post       Marks
		err        error
	}{
		{[2]int{0, 0}, 0, 1, []Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, nil},
		{[2]int{1, 1}, 1, 0, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, nil},
		{[2]int{4, 2}, 1, 1, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, errAny},
	}

	game, err := NewGame(3, 2, 0)
	if err != nil {
		t.Errorf("game creation failed: %v", err)
		return
	}

	for i, table := range tables {
		err := game.Move(table.pos, table.playerPre)
		if (err == nil) != (table.err == nil) {
			t.Errorf("unexpected error behavior in step %v: expected = \"%v\", actual = \"%v\"",
				i+1, table.err, err)
			continue
		}
		if err != nil {
			continue
		}
		if !game.Board.Marks.Equal(table.post) {
			t.Errorf("board different in step %v: expected = %v, actual = %v", i+1,
				table.post, game.Board.Marks)
		}
		if game.NextPlayer != table.playerPost {
			t.Errorf("next player wrong in step %v: expected = %v, actual = %v", i+1,
				table.playerPost, game.NextPlayer)
		}
	}
}
