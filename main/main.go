package main

import (
	"fmt"

	g "go-tictactoe/game"
	"go-tictactoe/rules"
)

func main() {
	game, err := g.NewGame(3, 3, 0)
	if err != nil {
		fmt.Println(err)
	}

	for {
		player := game.Players[game.CurrentPlayer.Next-1]
		pos, err := player.GetMove(game.Board)
		if err != nil {
			break
		}

		game.Move(pos, game.CurrentPlayer.Next)
		fmt.Println(game.Board)

		if w, hw := rules.GetWinner(game.Board); hw {
			fmt.Printf("Player %v won, congrats.\n", w)
			fmt.Println("Congrats.")
			break
		}

		if rules.IsFull(game.Board) {
			fmt.Println("It's a draw.")
			break
		}
	}
}
