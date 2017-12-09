package test

import (
	"fmt"
	"testing"
)

func TestCond(t *testing.T) {
	testCases := []struct {
		a, b, c bool
	}{
		{false, false, true},
		{false, true, true},
		{true, false, false},
		{true, true, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(" %v -> %v", tc.a, tc.b), func(t *testing.T) {
			switch c := Cond(tc.a, tc.b); false {
			case tc.c == c:
				t.Errorf("false result of conditional: expected =  %v, actual = %v",
					tc.c, c)
			}
		})
	}
}
