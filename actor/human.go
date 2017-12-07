package actor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-tictactoe/mechanics"
)

// Human represents a human player.
type Human struct {
	ID mechanics.Player
}

// GetMove returns the move the player makes after prompting them for input.
func (h *Human) GetMove(b mechanics.Board) (mechanics.Position, error) {
	// TODO input in chess format (e.g. a1)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Your move, player %v:\n", h.ID+1)

	for scanner.Scan() {
		pos, msg, err := isAcceptableMove(b, scanner.Text())
		if err != nil || msg == "" {
			return pos, err
		}

		fmt.Println(msg)
	}
	if err := scanner.Err(); err != nil {
		return mechanics.Position{0, 0}, err
	}
	return mechanics.Position{0, 0}, errors.New("gathering input failed unexpectedly")
}

func isAcceptableMove(b mechanics.Board, s string) (pos mechanics.Position, msg string, err error) {
	s = strings.Trim(s, " ")
	if s == "quit" || s == "exit" {
		return mechanics.Position{0, 0}, "", errors.New("quit")
	}

	split := strings.Split(s, ",")

	if len(split) != 2 {
		return mechanics.Position{0, 0}, "input must have contain two ints", nil
	}

	var err2 error
	for i := 0; i < 2; i++ {
		pos[i], err2 = strconv.Atoi(strings.Trim(split[i], " \n"))
		if err2 != nil {
			return mechanics.Position{0, 0}, "both parameters must be numbers", nil
		}
	}

	if ok, reason := b.IsWritable(pos); ok == false {
		return mechanics.Position{0, 0}, reason, nil
	}

	return pos, "", nil
}
