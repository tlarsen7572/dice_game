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

func TestRolledDiceAreSorted(t *testing.T) {
	for i := 0; i < 1000; i++ {
		roll := server.RollDice(2)
		if roll[0] > roll[1] {
			t.Fatalf(`roll %v was not sorted in ascending order`, roll)
		}
	}
}

func TestScoreOne(t *testing.T) {
	roll := []int{1}
	score, scoringDice := server.Score(roll)
	if score != 100 || len(scoringDice) != 1 || scoringDice[0] != 0 {
		t.Fatalf(`expected score of 100 and scoring dice of [0] but got %v and %v`, score, scoringDice)
	}
}

func TestScoreOneWithTwoDice(t *testing.T) {
	roll := []int{2, 1}
	score, scoringDice := server.Score(roll)
	if score != 100 || len(scoringDice) != 1 || scoringDice[0] != 1 {
		t.Fatalf(`expected score of 100 and scoring dice of [1] but got %v and %v`, score, scoringDice)
	}
}

func TestScoreMultipleOnes(t *testing.T) {
	roll := []int{1, 1}
	score, scoringDice := server.Score(roll)
	if score != 200 || len(scoringDice) != 2 || scoringDice[0] != 0 || scoringDice[1] != 1 {
		t.Fatalf(`expected score of 200 and scoring dice of [0 1] but got %v and %v`, score, scoringDice)
	}
}

func TestScoreThreeOnes(t *testing.T) {
	roll := []int{1, 1, 1}
	score, scoringDice := server.Score(roll)
	if score != 1000 || len(scoringDice) != 3 {
		t.Fatalf(`expected score of 1000 and scoring dice of [0 1 2] but got %v and %v`, score, scoringDice)
	}
}

func TestScoreFourOnes(t *testing.T) {
	roll := []int{1, 1, 1, 1}
	score, scoringDice := server.Score(roll)
	if score != 1100 || len(scoringDice) != 4 {
		t.Fatalf(`expected score of 1100 and scoring dice of [0 1 2 3] but got %v and %v`, score, scoringDice)
	}
}

func TestScoreSixOnes(t *testing.T) {
	roll := []int{1, 1, 1, 1, 1, 1}
	score, scoringDice := server.Score(roll)
	if score != 10000 || len(scoringDice) != 6 {
		t.Fatalf(`expected score of max int64 and scoring dice of [0 1 2 3 4 5] but got %v and %v`, score, scoringDice)
	}
}
