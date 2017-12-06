package main

import (
	"fmt"

	"go-tictactoe/actor"
	"go-tictactoe/mechanics"
)

func main() {
	game, err := mechanics.NewGame(3, 2, 2)
	if err != nil {
		fmt.Println(err)
	}

	human := actor.Human{ID: 0}
	for {
		pos, err := human.GetMove(game.Field)
		if err != nil {
			break
		}

		game.Move(pos, game.NextPlayer)
		fmt.Println(game.Field)
	}
}
