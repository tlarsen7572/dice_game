package server_test

import "testing"
import "server"

func TestRollDice(t *testing.T) {
	roll := server.RollDice(6)
	if actualCount := len(roll); actualCount != 6 {
		t.Fatalf(`expected 6 dice but got %v`, actualCount)
	}

	roll = server.RollDice(3)
	if actualCount := len(roll); actualCount != 3 {
		t.Fatalf(`expected 3 dice but got %v`, actualCount)
	}
}
