package internalsensors

import "testing"

func TestHullIntegrity_Value(t *testing.T) {
	sec := NewHullIntegrity(49)
	num := sec.Value()
	if num != 49 {
		t.Errorf("expected section number to be 49, received: %d", num)
	}
}
