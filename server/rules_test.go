package server_test

import (
	"testing"
)
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

func TestRolledDiceAreInRange(t *testing.T) {
	roll := server.RollDice(1)
	if roll[0] < 1 || roll[0] > 6 {
		t.Fatalf(`expected a dice roll between 1 and 6`)
	}
}

func TestRolledDiceAreDistributedInRange(t *testing.T) {
	results := make([]int, 6)
	for i := 0; i < 60000; i++ {
		roll := server.RollDice(1)
		resultIndex := roll[0] - 1
		results[resultIndex]++
	}
	t.Logf(`results: %v`, results)
	for diceNumber, actualCount := range results {
		if actualCount < 9000 || actualCount > 11000 {
			t.Fatalf(`an unexpected number of %vs were rolled. Expected between 9000 and 11000 but got %v`, diceNumber+1, actualCount)
		}
	}
}
