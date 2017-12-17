package game

import (
	"fmt"
	"testing"

	"github.com/nilsbu/fastest"

	a "github.com/nilsbu/go-tictactoe/actor"
	b "github.com/nilsbu/go-tictactoe/board"
	"github.com/nilsbu/go-tictactoe/io"
)

func TestSymbols(t *testing.T) {
	ft := fastest.T{T: t}

	ft.Equals(MaxPlayers, len(io.Symbols)-1)
}

func TestPlayerCounter_Inc(t *testing.T) {
	ft := fastest.T{T: t}

	for n := 2; n <= 4; n++ {
		pc := PlayerCounter{Next: 1, Total: b.Player(n)}

		ft.Seq(fmt.Sprintf("n=%v", n), func(ft fastest.T) {
			for i := 0; i <= 2*n; i++ {
				ft.Equals(pc.Next, b.Player(i%n+1), "after %v increases, counter should be %v but was %v", i, i%n+1, pc.Next)

				pc.Inc()
			}
		})
	}
}

func TestNewGame(t *testing.T) {
	ft := fastest.T{T: t}

	testCases := []struct {
		size         int
		humanPlayers int
		players      []a.Actor
		err          bool
	}{
		{3, 2, []a.Actor{&a.Human{}, &a.Human{}}, false},
		{4, 1, []a.Actor{&a.Human{}, &a.Computer{ID: 2, Players: 2}}, false},
		{5, 0, []a.Actor{&a.Computer{ID: 1, Players: 3}, &a.Computer{ID: 2,
			Players: 3}, &a.Computer{ID: 3, Players: 3}}, false},
		{3, 1, []a.Actor{&a.Human{}}, true},
		{2, 1, []a.Actor{&a.Human{}, &a.Computer{ID: 2, Players: 2}}, true},
		{3, 3, []a.Actor{&a.Human{}, &a.Human{}}, true},
	}

	for i, tc := range testCases {
		s := fmt.Sprintf("#%v: %vx%v with %v", i, tc.size, tc.size, tc.players)
		ft.Seq(s, func(ft fastest.T) {
			game, err := NewGame(tc.size, len(tc.players), tc.humanPlayers)

			ft.Implies(tc.err == true, err != nil, "expected error but none was returned")
			ft.Implies(tc.err == false, err == nil, "no error expected but one was returned")
			ft.Only(tc.err == false)
			ft.Equals(len(game.Players), len(tc.players))
			ft.True(equalsActors(tc.players, game.Players), "player setup: expected = %v, actual %v", tc.players, game.Players)
			ft.Equals(len(game.Board.(b.Data).Marks), tc.size*tc.size)
			ft.Equals(game.Board.(b.Data).Size, tc.size)
			ft.Equals(b.Player(1), game.CurrentPlayer.Next)
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
	ft := fastest.T{T: t}

	testCases := []struct {
		pos      b.Position
		plyrPre  b.Player
		plyrPost b.Player
		post     b.Marks
		err      bool
	}{
		{[2]int{0, 0}, 1, 2, []b.Player{1, 0, 0, 0, 0, 0, 0, 0, 0}, false},
		{[2]int{1, 1}, 2, 1, []b.Player{1, 0, 0, 0, 2, 0, 0, 0, 0}, false},
		{[2]int{2, 2}, 2, 1, nil, true}, // False NextPlayer
		{[2]int{1, 1}, 1, 1, nil, true}, // Field not empty
		{[2]int{4, 1}, 1, 1, nil, true}, // Outside board
	}

	g, err := NewGame(3, 2, 0)
	ft.Nil(err, "game creation failed: %v", err)

	for i, tc := range testCases {
		s := fmt.Sprintf("#%v: %v at %v", i, tc.plyrPre, tc.pos)
		ft.Seq(s, func(ft fastest.T) {
			err := g.Move(tc.pos, tc.plyrPre)

			ft.Implies(tc.err == true, err != nil, "expected error but none was returned")
			ft.Implies(tc.err == false, err == nil, "no error expected but one was returned")
			ft.Only(tc.err == false)
			ft.True(g.Board.(b.Data).Marks.Equal(tc.post), "board different: expected = %v, actual = %v", tc.post, g.Board.(b.Data).Marks)
			ft.Equals(g.CurrentPlayer.Next, tc.plyrPost)
		})
	}
}
