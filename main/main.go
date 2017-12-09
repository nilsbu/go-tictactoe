package main

import (
	"fmt"

	a "go-tictactoe/actor"
	g "go-tictactoe/game"
)

func main() {
	game, err := g.NewGame(3, 2, 1)
	if err != nil {
		fmt.Println(err)
	}

	players := []a.Actor{
		&a.Human{ID: 1},
		&a.Computer{ID: 2, Players: 2},
	}

	for {
		pos, err := players[game.CurrentPlayer.Next-1].GetMove(game.Board)
		if err != nil {
			break
		}

		game.Move(pos, game.CurrentPlayer.Next)
		fmt.Println(game.Board)
	}
}
