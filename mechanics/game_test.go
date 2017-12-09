package mechanics

import (
	"fmt"
	"testing"

	"go-tictactoe/test"
)

func TestPlayerCounter_Inc(t *testing.T) {
	for n := 2; n <= 4; n++ {
		pc := PlayerCounter{Next: 1, Total: Player(n)}

		t.Run(fmt.Sprintf("n=%v", n), func(t *testing.T) {
			for i := 0; i <= 2*n; i++ {
				if pc.Next != Player(i%n+1) {
					t.Errorf("after %v increases, counter should be %v but was %v",
						i, i%n+1, pc.Next)
					break
				}
				pc.Inc()
			}
		})
	}
}

func TestNewGame(t *testing.T) {
	testCases := []struct {
		size         int
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

	for i, tc := range testCases {
		s := fmt.Sprintf("#%v: %vx%v with %v", i, tc.size, tc.size, tc.players)
		t.Run(s, func(t *testing.T) {
			switch game, err :=
				NewGame(tc.size, len(tc.players), tc.humanPlayers); false {
			case test.Cond(tc.err == test.AnyError, err != nil):
				t.Errorf("expected error but none was returned")
			case test.Cond(tc.err == test.NoError, err == nil):
				t.Errorf("no error expected but one was returned")
			case tc.err == test.NoError:
			case len(game.Players) == len(tc.players):
				t.Errorf("number of players: expected = %v, actual = %v",
					len(tc.players), len(game.Players))
			case equals(tc.players, game.Players):
				t.Errorf("player setup: expected = %v, actual %v",
					tc.players, game.Players)
			case len(game.Board.Marks) == tc.size*tc.size:
				t.Errorf("marks size: expected = %v, actual = %v",
					tc.size*tc.size, len(game.Board.Marks))
			case game.Board.Size == tc.size:
				t.Errorf("board size: expected = %v, actual = %v",
					tc.size, game.Board.Size)
			case game.CurrentPlayer.Next == 1:
				t.Errorf("next player: expected = 1, actual = %v",
					game.CurrentPlayer.Next)
			}
		})
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
	testCases := []struct {
		pos      Position
		plyrPre  Player
		plyrPost Player
		post     Marks
		err      test.ErrorAnticipation
	}{
		{[2]int{0, 0}, 1, 2, []Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, test.NoError},
		{[2]int{1, 1}, 2, 1, []Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, test.NoError},
		{[2]int{2, 2}, 2, 1, nil, test.AnyError}, // False NextPlayer
		{[2]int{1, 1}, 1, 1, nil, test.AnyError}, // Field not empty
		{[2]int{4, 1}, 1, 1, nil, test.AnyError}, // Outside board
	}

	game, err := NewGame(3, 2, 0)
	if err != nil {
		t.Errorf("game creation failed: %v", err)
		return
	}

	for i, tc := range testCases {
		s := fmt.Sprintf("#%v: %v at %v", i, tc.plyrPre, tc.pos)
		t.Run(s, func(t *testing.T) {
			switch err := game.Move(tc.pos, tc.plyrPre); false {
			case test.Cond(tc.err == test.AnyError, err != nil):
				t.Errorf("expected error but none was returned")
			case test.Cond(tc.err == test.NoError, err == nil):
				t.Errorf("no error expected but one was returned")
			case tc.err == test.NoError:
			case game.Board.Marks.Equal(tc.post):
				t.Errorf("board different: expected = %v, actual = %v",
					tc.post, game.Board.Marks)
			case game.CurrentPlayer.Next == tc.plyrPost:
				t.Errorf("next player wrong: expected = %v, actual = %v",
					tc.plyrPost, game.CurrentPlayer)
			}
		})
	}
}
