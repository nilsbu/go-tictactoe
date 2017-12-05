package main

import (
	"go-tictactoe/actor"
	"go-tictactoe/mechanics"
	"fmt"
)

func main() {
	game, err := mechanics.NewGame(3, 2, 2)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	human := actor.Human{0}
	for {
		pos, err := human.GetMove(game.Field)
		if err != nil {
			break
		}

		game.Move(pos, game.NextPlayer)
		fmt.Println(game.Field)
	}
}
