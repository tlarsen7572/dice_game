package server_test

import (
	"server"
	"testing"
)

func TestStartGame(t *testing.T) {
	game := server.NewGame(10000)
	if game.Roll == nil {
		t.Fatalf(`expected non-nil Roll action but got nil`)
	}
	if game.NewTurn != nil {
		t.Fatalf(`expected nil NewTurn action but got non-nil`)
	}
}

type MockRoller struct {
	RollOverride  []int
	RequestedDice int
}

func (r *MockRoller) Roll(totalDice int) []int {
	r.RequestedDice = totalDice
	if r.RollOverride == nil {
		return server.RollDice(totalDice)
	}
	return r.RollOverride
}

func TestTurnEndingWithZeroPoints(t *testing.T) {
	mockRoll := &MockRoller{}
	game := server.NewGame(10000, server.OverrideRollWith(mockRoll))
	mockRoll.RollOverride = []int{1, 2, 2, 3, 3, 4}
	game.Roll()
	score := game.LastRolledScore()
	if mockRoll.RequestedDice != 6 {
		t.Fatalf(`expected 6 requested dice but got %v`, mockRoll.RequestedDice)
	}
	if score.Score != 100 {
		t.Fatalf(`expected score of 100 but got %v`, score.Score)
	}
	if game.CurrentTurnScore() != 100 {
		t.Fatalf(`expected 100 but got %v`, game.CurrentTurnScore())
	}
	if game.Roll == nil {
		t.Fatalf(`expected nil Roll action but got non-nil`)
	}

	mockRoll.RollOverride = []int{1, 2, 3, 3, 4}
	game.Roll()
	score = game.LastRolledScore()
	if mockRoll.RequestedDice != 5 {
		t.Fatalf(`expected 5 requested dice but got %v`, mockRoll.RequestedDice)
	}
	if score.Score != 100 {
		t.Fatalf(`expected score of 100 but got %v`, score.Score)
	}
	if game.CurrentTurnScore() != 200 {
		t.Fatalf(`expected 200 but got %v`, game.CurrentTurnScore())
	}
	if game.Roll == nil {
		t.Fatalf(`expected nil Roll action but got non-nil`)
	}

	mockRoll.RollOverride = []int{2, 3, 3, 4}
	game.Roll()
	score = game.LastRolledScore()
	if mockRoll.RequestedDice != 4 {
		t.Fatalf(`expected 4 requested dice but got %v`, mockRoll.RequestedDice)
	}
	if score.Score != 0 {
		t.Fatalf(`expected score of 0 but got %v`, score.Score)
	}
	if game.CurrentTurnScore() != 0 {
		t.Fatalf(`expected 0 but got %v`, game.CurrentTurnScore())
	}
	if game.Roll != nil {
		t.Fatalf(`expected non-nil Roll action but got nil`)
	}
	if game.NewTurn == nil {
		t.Fatalf(`expected non-nil NewTurn action but got nil`)
	}
}
