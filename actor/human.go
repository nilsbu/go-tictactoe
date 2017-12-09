package actor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	m "go-tictactoe/mechanics"
)

// Human represents a human player.
type Human struct {
	ID m.Player
}

// GetMove returns the move the player makes after prompting them for input.
func (h *Human) GetMove(b m.Board) (m.Position, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Your move, player %v:\n", h.ID)

	for scanner.Scan() {
		pos, msg, err := isAcceptableMove(b, scanner.Text())
		if err != nil || msg == "" {
			return pos, err
		}

		fmt.Println(msg)
	}
	if err := scanner.Err(); err != nil {
		return m.Position{0, 0}, err
	}
	return m.Position{0, 0}, errors.New("gathering input failed unexpectedly")
}

// isAcceptableMove checks if an input string corresponds to an acceptable move.
// A position is returned if the input is acceptable, a message for the user is
// is returned if the input was invalid.
// If the input process should be aborted, an error is returned.
func isAcceptableMove(b m.Board, s string) (pos m.Position,
	msg string, err error) {

	s = strings.Trim(s, " ")
	if s == "quit" || s == "exit" {
		return m.Position{0, 0}, "", errors.New("quit")
	}

	split := strings.Split(s, ",")

	if len(split) != 2 {
		return m.Position{0, 0}, "input must have contain two ints", nil
	}

	var err2 error
	for i := 0; i < 2; i++ {
		pos[i], err2 = strconv.Atoi(strings.Trim(split[i], " \n"))
		if err2 != nil {
			return m.Position{0, 0}, "both parameters must be numbers", nil
		}
	}

	if ok, reason := b.IsWritable(pos); ok == false {
		return m.Position{0, 0}, reason, nil
	}

	return pos, "", nil
}
