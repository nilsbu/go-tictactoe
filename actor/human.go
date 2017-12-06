package actor

import (
	"bufio"
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
func (h *Human) GetMove(field mechanics.Field) (pos mechanics.Position, err error) {
	// TODO There should be a method to quit the game here.
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Next move, player %v: ", h.ID+1)

		var text string
		text, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		pos, err = splitString(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !isInField(pos, field.Size) {
			fmt.Println("Entered position is not in field.")
			continue
		}

		break
	}

	return pos, err
}

func splitString(s string) (pos mechanics.Position, err error) {
	split := strings.Split(s, ",")

	if len(split) != 2 {
		err = fmt.Errorf("Input must have contain two ints.")
		return
	}

	for i := 0; i < 2; i++ {
		pos[i], err = strconv.Atoi(strings.Trim(split[i], " \n"))
		if err != nil {
			return
		}
	}

	return
}

func isInField(pos mechanics.Position, size int) bool {
	for i := 0; i < 2; i++ {
		if pos[i] < 0 || pos[i] >= size {
			return false
		}
	}

	return true
}
