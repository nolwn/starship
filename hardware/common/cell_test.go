package common

import "testing"

func TestCell_Level(t *testing.T) {
	c := NewCell(1_000)

	if c.Level() != 1_000 {
		t.Error("Level should return the amount of energy available")
	}

	c.level = 345
	if c.Level() != 345 {
		t.Error("Level should return the amount of energy available")
	}
}
