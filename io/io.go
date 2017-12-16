package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	b "github.com/nilsbu/go-tictactoe/board"
)

// IO bundles calls to input and output.
type IO interface {
	Out(a ...interface{})
	Outf(s string, a ...interface{})
	Outln(a ...interface{})
	OutBoard(b.Board)
	In() (string, error)
}

// Console is the standard input and output on the console.
type Console struct {
	in  *bufio.Scanner
	out *bufio.Writer
}

// Symbols stores the marks that players make on the board.
// The first one is the mark of an empty board, the subsequent ones belong to
// the players.
var Symbols = []string{" ", "x", "o", "^", "#", "v"}

// NewConsole is the constructor for Console.
func NewConsole() Console {
	return Console{
		in:  bufio.NewScanner(os.Stdin),
		out: bufio.NewWriter(os.Stdout),
	}
}

// Out prints a with the formatting done by fmt.Sprint.
func (io Console) Out(a ...interface{}) {
	io.out.WriteString(fmt.Sprint(a...))
	io.out.Flush()
}

// Outf formats "s" with "a" as paramters and prints them.
func (io Console) Outf(s string, a ...interface{}) {
	io.Out(fmt.Sprintf(s, a...))
}

// Outln prints a with newline appended.
func (io Console) Outln(a ...interface{}) {
	s := fmt.Sprint(a...)
	io.Out(s + "\n")
}

//OutBoard prints the board.
func (io Console) OutBoard(bo b.Board) {
	io.Out(formatBoard(bo))
}

func formatBoard(bo b.Board) string {
	d := bo.Get()
	s := strings.Repeat("-", 2*d.Size+1) + "\n"

	for y := 0; y < d.Size; y++ {
		s += "|"

		for x := 0; x < d.Size; x++ {
			s += fmt.Sprintf("%v|", Symbols[d.Marks[y*d.Size+x]])
		}

		s += "\n"
		s += strings.Repeat("-", 2*d.Size+1) + "\n"
	}

	return s
}

// In reads the console input that ends with newline.
// If it fails an error is returned.
func (io Console) In() (string, error) {
	for io.in.Scan() {
		return io.in.Text(), nil
	}

	return "", io.in.Err()
}
