package server_test

import (
	"server"
	"testing"
)

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

func TestGameFirstScoringTurn(t *testing.T) {
	mockRoller := &MockRoller{
		RollOverride: []int{1, 2, 2, 3, 3, 4},
	}
	game := &server.Game{
		WinningScore: 10000,
		Roller:       mockRoller.Roll,
	}
	game.Roll()
	if game.CurrentScore != 0 {
		t.Fatalf(`expected 0 but got %v`, game.CurrentScore)
	}
	game.NewTurn()
	if game.CurrentScore != 100 {
		t.Fatalf(`expected 100 but got %v`, game.CurrentScore)
	}
	if len(game.Turns) != 1 {
		t.Fatalf(`expected 1 turn but got %v`, len(game.Turns))
	}
}

func TestGameToJson(t *testing.T) {
	game := &server.Game{
		WinningScore: 10000,
		CurrentScore: 500,
		Roller:       server.RollDice,
		Turns:        []int{100, 200, 0, 200},
		ActiveTurn: &server.Turn{
			Score:           100,
			LastRoll:        []int{1, 2, 2, 3, 3, 4},
			LastScoringDice: []int{0},
		},
	}

	actualJsonStr := game.ToJson()
	expectedJsonStr := `{"WinningScore":10000,"CurrentScore":500,"Turns":[100,200,0,200],"ActiveTurn":{"Score":100,"LastRoll":[1,2,2,3,3,4],"LastScoringDice":[0]}}`
	if actualJsonStr != expectedJsonStr {
		t.Fatalf("expected\n'%v'\nbut got\n'%v'", expectedJsonStr, actualJsonStr)
	}
}
