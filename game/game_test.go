package game

import (
	"fmt"
	"testing"

	a "go-tictactoe/actor"
	b "go-tictactoe/board"
	"go-tictactoe/test"
)

func TestPlayerCounter_Inc(t *testing.T) {
	for n := 2; n <= 4; n++ {
		pc := PlayerCounter{Next: 1, Total: b.Player(n)}

		t.Run(fmt.Sprintf("n=%v", n), func(t *testing.T) {
			for i := 0; i <= 2*n; i++ {
				if pc.Next != b.Player(i%n+1) {
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
		players      []a.Actor
		err          test.ErrorAnticipation
	}{
		{3, 2, []a.Actor{&a.Human{}, &a.Human{}}, test.NoError},
		{4, 1, []a.Actor{&a.Human{}, &a.Computer{ID: 2, Players: 2}},
			test.NoError},
		{5, 0, []a.Actor{&a.Computer{ID: 1, Players: 3}, &a.Computer{ID: 2,
			Players: 3}, &a.Computer{ID: 3, Players: 3}}, test.NoError},
		{3, 1, []a.Actor{&a.Human{}}, test.AnyError},
		{2, 1, []a.Actor{&a.Human{}, &a.Computer{ID: 2, Players: 2}},
			test.AnyError},
		{3, 3, []a.Actor{&a.Human{}, &a.Human{}}, test.AnyError},
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
			case equalsActors(tc.players, game.Players):
				t.Errorf("player setup: expected = %v, actual %v",
					tc.players, game.Players)
			case len(game.Board.(b.Data).Marks) == tc.size*tc.size:
				t.Errorf("marks size: expected = %v, actual = %v",
					tc.size*tc.size, len(game.Board.(b.Data).Marks))
			case game.Board.(b.Data).Size == tc.size:
				t.Errorf("board size: expected = %v, actual = %v",
					tc.size, game.Board.(b.Data).Size)
			case game.CurrentPlayer.Next == 1:
				t.Errorf("next player: expected = 1, actual = %v",
					game.CurrentPlayer.Next)
			}
		})
	}
}

func equalsActors(ps []a.Actor, os []a.Actor) bool {
	// Same length is assumed
	for i := 0; i < len(ps); i++ {
		if !equals(ps[i], os[i]) {
			return false
		}
	}
	return true
}

func equals(p a.Actor, o a.Actor) bool {
	if _, aok := p.(*a.Human); aok {
		if _, bok := o.(*a.Human); bok {
			return true
		}
		return false
	}
	if pv, aok := p.(*a.Computer); aok {
		if ov, bok := o.(*a.Computer); bok {
			return pv.ID == ov.ID && pv.Players == ov.Players
		}
		return false
	}
	return false
}

// TODO Test NextPlayer for more than two players
func TestGame_Move2(t *testing.T) {
	testCases := []struct {
		pos      b.Position
		plyrPre  b.Player
		plyrPost b.Player
		post     b.Marks
		err      test.ErrorAnticipation
	}{
		{[2]int{0, 0}, 1, 2, []b.Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, test.NoError},
		{[2]int{1, 1}, 2, 1, []b.Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, test.NoError},
		{[2]int{2, 2}, 2, 1, nil, test.AnyError}, // False NextPlayer
		{[2]int{1, 1}, 1, 1, nil, test.AnyError}, // Field not empty
		{[2]int{4, 1}, 1, 1, nil, test.AnyError}, // Outside board
	}

	g, err := NewGame(3, 2, 0)
	if err != nil {
		t.Errorf("game creation failed: %v", err)
		return
	}

	for i, tc := range testCases {
		s := fmt.Sprintf("#%v: %v at %v", i, tc.plyrPre, tc.pos)
		t.Run(s, func(t *testing.T) {
			switch err := g.Move(tc.pos, tc.plyrPre); false {
			case test.Cond(tc.err == test.AnyError, err != nil):
				t.Errorf("expected error but none was returned")
			case test.Cond(tc.err == test.NoError, err == nil):
				t.Errorf("no error expected but one was returned")
			case tc.err == test.NoError:
			case g.Board.(b.Data).Marks.Equal(tc.post):
				t.Errorf("board different: expected = %v, actual = %v",
					tc.post, g.Board.(b.Data).Marks)
			case g.CurrentPlayer.Next == tc.plyrPost:
				t.Errorf("next player wrong: expected = %v, actual = %v",
					tc.plyrPost, g.CurrentPlayer)
			}
		})
	}
}
