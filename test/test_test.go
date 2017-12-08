package test

import "testing"

func TestCond(t *testing.T) {
	tables := []struct {
		a, b, c bool
	}{
		{false, false, true},
		{false, true, true},
		{true, false, false},
		{true, true, true},
	}

	for _, table := range tables {
		switch c := Cond(table.a, table.b); false {
		case table.c == c:
			t.Errorf("false result of conditional %v -> %v: expected =  %v, actual "+
				"= %v", table.a, table.b, table.c, c)
		}
	}
}
