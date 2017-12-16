package actor

import (
	"errors"
	"strconv"
	"strings"

	b "tictactoe/board"
	"tictactoe/io"
)

// Human represents a human player.
type Human struct {
	io io.IO
}

// NewHuman creates a human plauer with console as IO.
func NewHuman() *Human {
	return &Human{io: io.NewConsole()}
}

// GetMove returns the move the player makes after prompting them for input.
func (h *Human) GetMove(bo b.Board) (b.Position, error) {
	// TODO test
	for {
		s, err1 := h.io.In()
		if err1 != nil {
			return b.Position{0, 0}, err1
		}

		pos, msg, err2 := isAcceptableMove(bo, s)
		if err2 != nil || msg == "" {
			return pos, err2
		}

		h.io.Outln(msg)
	}
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
