package mechanics

import (
	"testing"

	"go-tictactoe/test"
)

func TestPlayerCounter_Inc(t *testing.T) {
	for n := 2; n <= 4; n++ {
		pc := PlayerCounter{Next: 0, Total: Player(n)}

		for i := 0; i <= 2*n; i++ {
			if pc.Next != Player(i%n) {
				t.Errorf("after %v increases, counter should be %v but was %v", i, i%n,
					pc.Next)
				break
			}

			pc.Inc()
		}
	}
}

func TestNewGame(t *testing.T) {
	tables := []struct {
		boardSize    int
		humanPlayers int
		players      []PlayerType
		err          test.ErrorAnticipation
	}{
		{3, 2, []PlayerType{Human, Human}, test.NoError},
		{4, 1, []PlayerType{Human, Computer}, test.NoError},
		{5, 0, []PlayerType{Computer, Computer, Computer}, test.NoError},
		{3, 1, []PlayerType{Human}, test.AnyError},
		{2, 1, []PlayerType{Human, Computer}, test.AnyError},
		{3, 3, []PlayerType{Human, Human}, test.AnyError},
	}

	for i, table := range tables {
		switch game, err :=
			NewGame(table.boardSize, len(table.players), table.humanPlayers); false {
		case test.Cond(table.err == test.AnyError, err != nil):
			t.Errorf("expected error in step %v but none was returned", i+1)
		case test.Cond(table.err == test.NoError, err == nil):
			t.Errorf("no error expected in step %v but one was returned", i+1)
		case table.err == test.NoError:
		case len(game.Players) == len(table.players):
			t.Errorf("number of players in step %v: expected = %v, actual = %v", i+1,
				len(table.players), len(game.Players))
		case equals(table.players, game.Players):
			t.Errorf("player setup in step %v: expected = %v, actual %v", i+1,
				table.players, game.Players)
		case len(game.Board.Marks) == table.boardSize*table.boardSize:
			t.Errorf("marks size in step %v: expected = %v, actual = %v", i+1,
				table.boardSize*table.boardSize, len(game.Board.Marks))
		case game.Board.Size == table.boardSize:
			t.Errorf("board size in step %v: expected = %v, actual = %v", i+1,
				table.boardSize, game.Board.Size)
		case game.CurrentPlayer.Next == Player(0):
			t.Errorf("next player in step %v: expected = 0, actual = %v", i+1,
				game.CurrentPlayer.Next)
		}
	}
}

func equals(ps []PlayerType, os []PlayerType) bool {
	// Same length is assumed
	for i := 0; i < len(ps); i++ {
		if ps[i] != os[i] {
			return false
		}
	}
	return true
}

// TODO Test NextPlayer for more than two players
func TestGame_Move2(t *testing.T) {
	tables := []struct {
		pos        Position
		playerPre  Player
		playerPost Player
		post       Marks
		err        test.ErrorAnticipation
	}{
		{[2]int{0, 0}, 0, 1, []Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, test.NoError},
		{[2]int{1, 1}, 1, 0, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, test.NoError},
		{[2]int{2, 2}, 1, 0, nil, test.AnyError}, // False NextPlayer
		{[2]int{1, 1}, 0, 0, nil, test.AnyError}, // Field not empty
		{[2]int{4, 1}, 0, 0, nil, test.AnyError}, // Outside board
	}

	game, err := NewGame(3, 2, 0)
	if err != nil {
		t.Errorf("game creation failed: %v", err)
		return
	}

	for i, table := range tables {
		switch err := game.Move(table.pos, table.playerPre); false {
		case test.Cond(table.err == test.AnyError, err != nil):
			t.Errorf("expected error in step %v but none was returned", i+1)
		case test.Cond(table.err == test.NoError, err == nil):
			t.Errorf("no error expected in step %v but one was returned", i+1)
		case table.err == test.NoError:
		case game.Board.Marks.Equal(table.post):
			t.Errorf("board different in step %v: expected = %v, actual = %v", i+1,
				table.post, game.Board.Marks)
		case game.CurrentPlayer.Next == table.playerPost:
			t.Errorf("next player wrong in step %v: expected = %v, actual = %v", i+1,
				table.playerPost, game.CurrentPlayer)
		}
	}
}
