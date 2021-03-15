package turn_test

import (
	"reflect"
	turn2 "server/turn"
	"testing"
)

func TestNormalTurn(t *testing.T) {
	turn := &turn2.Turn{}
	turn.Roll([]int{1, 2, 2, 3, 3, 4})
	if turn.Score != 100 {
		t.Fatalf(`expected 100 but got %v`, turn.Score)
	}
	if !reflect.DeepEqual(turn.LastRoll, []int{1, 2, 2, 3, 3, 4}) {
		t.Fatalf(`expected %v but got %v`, []int{1, 2, 2, 3, 3, 4}, turn.LastRoll)
	}
	if !reflect.DeepEqual(turn.LastScoringDice, []int{0}) {
		t.Fatalf(`expected %v but got %v`, []int{0}, turn.LastScoringDice)
	}
}

func TestScoringZeroReducesTotalScoreToZero(t *testing.T) {
	turn := &turn2.Turn{}
	turn.Roll([]int{1, 2, 2, 3, 3, 4})
	turn.Roll([]int{2, 2, 3, 3, 4})
	if turn.Score != 0 {
		t.Fatalf(`expected 0 but got %v`, turn.Score)
	}
}

func TestResetTurn(t *testing.T) {
	turn := &turn2.Turn{}
	turn.Roll([]int{1, 2, 2, 3, 3, 4})
	turn.Reset()
	if turn.Score != 0 {
		t.Fatalf(`expected 0 but got %v`, turn.Score)
	}
	if !reflect.DeepEqual(turn.LastScoringDice, []int{}) {
		t.Fatalf(`expected an empty slice but got %v`, turn.LastScoringDice)
	}
	if !reflect.DeepEqual(turn.LastRoll, []int{}) {
		t.Fatalf(`expected an empty slice but got %v`, turn.LastRoll)
	}
}
