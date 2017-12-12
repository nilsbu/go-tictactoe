package main

import (
	"fmt"

	b "go-tictactoe/board"
	g "go-tictactoe/game"
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

		if w, hw := game.Board.(b.Outcome).GetWinner(); hw {
			fmt.Printf("Player %v won, congrats.\n", w)
			fmt.Println("Congrats.")
			break
		}

		if game.Board.(b.Outcome).IsFull() {
			fmt.Println("It's a draw.")
			break
		}
	}
}
