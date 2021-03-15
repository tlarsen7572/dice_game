package game_test

import (
	"encoding/json"
	game2 "server/game"
	"server/mock_roller"
	"server/rules"
	"server/turn"
	"testing"
)

func TestGameFirstScoringTurn(t *testing.T) {
	mockRoller := &mock_roller.MockRoller{
		RollOverride: []int{1, 2, 2, 3, 3, 4},
	}
	game := &game2.Game{
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
	game := &game2.Game{
		WinningScore: 10000,
		CurrentScore: 500,
		Roller:       rules.RollDice,
		Turns:        []int{100, 200, 0, 200},
		ActiveTurn: &turn.Turn{
			Score:           100,
			LastRoll:        []int{1, 2, 2, 3, 3, 4},
			LastScoringDice: []int{0},
		},
	}

	actualJsonBytes, _ := json.Marshal(game)
	expectedJsonStr := `{"WinningScore":10000,"CurrentScore":500,"Turns":[100,200,0,200],"ActiveTurn":{"Score":100,"LastRoll":[1,2,2,3,3,4],"LastScoringDice":[0]}}`
	if actualJsonStr := string(actualJsonBytes); actualJsonStr != expectedJsonStr {
		t.Fatalf("expected\n'%v'\nbut got\n'%v'", expectedJsonStr, actualJsonStr)
	}
}

func TestNewGameToJson(t *testing.T) {
	game := game2.NewGame(10000)

	actualJsonBytes, _ := json.Marshal(game)
	expectedJsonStr := `{"WinningScore":10000,"CurrentScore":0,"Turns":[],"ActiveTurn":{"Score":0,"LastRoll":[],"LastScoringDice":[]}}`
	if actualJsonStr := string(actualJsonBytes); actualJsonStr != expectedJsonStr {
		t.Fatalf("expected\n'%v'\nbut got\n'%v'", expectedJsonStr, actualJsonStr)
	}
}
