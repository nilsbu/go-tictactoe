package main

import (
	"fmt"

	"go-tictactoe/actor"
	"go-tictactoe/mechanics"
)

func main() {
	game, err := mechanics.NewGame(3, 2, 1)
	if err != nil {
		fmt.Println(err)
	}

	players := []actor.Actor{
		&actor.Human{ID: 1},
		&actor.Computer{ID: 2, Players: 2},
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
