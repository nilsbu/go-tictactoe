package actor

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	b "go-tictactoe/board"
)

// Human represents a human player.
type Human struct {
}

// GetMove returns the move the player makes after prompting them for input.
func (h *Human) GetMove(bo b.Board) (b.Position, error) {
	// TODO test, move scanner in struct
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		pos, msg, err := isAcceptableMove(bo, scanner.Text())
		if err != nil || msg == "" {
			return pos, err
		}

		fmt.Println(msg)
	}

	return b.Position{0, 0}, errors.New("gathering input failed unexpectedly")
}

// isAcceptableMove checks if an input string corresponds to an acceptable move.
// A position is returned if the input is acceptable, a message for the user is
// is returned if the input was invalid.
// If the input process should be aborted, an error is returned.
func isAcceptableMove(bo b.Board, s string) (pos b.Position,
	msg string, err error) {

	s = strings.Trim(s, " ")
	if s == "quit" || s == "exit" {
		return b.Position{0, 0}, "", errors.New("quit")
	}

	split := strings.Split(s, ",")

	if len(split) != 2 {
		return b.Position{0, 0}, "input must have contain two ints", nil
	}

	var err2 error
	for i := 0; i < 2; i++ {
		pos[i], err2 = strconv.Atoi(strings.Trim(split[i], " \n"))
		if err2 != nil {
			return b.Position{0, 0}, "both parameters must be numbers", nil
		}
	}

	return pos, "", nil
}
