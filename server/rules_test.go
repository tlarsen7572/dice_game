package server_test

import "testing"
import "server"

func TestRolledDice(t *testing.T) {
	roll := server.RollDice(6)
	if actualCount := len(roll); actualCount != 6 {
		t.Fatalf(`expected 6 dice but got %v`, actualCount)
	}
}
