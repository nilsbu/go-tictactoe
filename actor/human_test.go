package actor

import (
	"errors"
	"go-tictactoe/mechanics"
	"testing"
)

func TestSplitString(t *testing.T) {
	tables := []struct {
		s   string
		pos mechanics.Position
		err error
	}{
		{"0,1", mechanics.Position{0, 1}, nil},
		{"011,1", mechanics.Position{11, 1}, nil},
		{"2, 1", mechanics.Position{2, 1}, nil},
		{"x4,1", mechanics.Position{2, 1}, errors.New("Error")},
		{"4,1,2", mechanics.Position{2, 1}, errors.New("Error")},
		{"1232", mechanics.Position{2, 1}, errors.New("Error")},
		{"", mechanics.Position{2, 1}, errors.New("Error")},
	}

	for _, table := range tables {
		pos, err := splitString(table.s)
		if (table.err == nil) != (err == nil) {
			t.Errorf("unexpected error behavior: expected = \"%v\", actual = \"%v\"",
				table.err, err)
			continue
		}
		if err != nil {
			continue
		}
		if table.pos != pos {
			t.Errorf("position: expected = %v, actual = %v", table.pos, pos)
		}
	}
}

func TestIsInField(t *testing.T) {
	tables := []struct {
		pos  mechanics.Position
		size int
		ok   bool
	}{
		{mechanics.Position{0, 1}, 3, true},
		{mechanics.Position{0, 3}, 3, false},
		{mechanics.Position{-1, 1}, 3, false},
		{mechanics.Position{3, 0}, 4, true},
	}

	for _, table := range tables {
		ok := isInField(table.pos, table.size)
		if table.ok != ok {
			t.Errorf("position = %v with size = %v: expected = %v, actual = %v",
				table.pos, table.size, table.ok, ok)
		}
	}
}
