package server_test

import (
	"encoding/json"
	"regexp"
	"server"
	"testing"
)

func TestNewGameManager(t *testing.T) {
	gameManager := server.NewGameManager("http://localhost")
	if gameManager.ActiveActions.NewGameAction == nil {
		t.Fatalf(`expected non-nil new game action but got nil`)
	}
	if gameManager.ActiveActions.RollAction != nil {
		t.Fatalf(`expected nil roll action but got non-nil`)
	}
	if gameManager.ActiveActions.NewTurnAction != nil {
		t.Fatalf(`expected nil turn action but got non-nil`)
	}
	if actions := len(gameManager.ActionLinks); actions != 1 {
		t.Fatalf(`expected 1 action but got %v`, actions)
	}
	for _, action := range gameManager.ActionLinks {
		if action.Type != server.NewGameAction {
			t.Fatalf(`expected ActionType %v but got %v`, server.NewGameAction, action.Type)
		}
		if action.Method != `POST` {
			t.Fatalf(`expected POST but got %v`, action.Method)
		}
		expectedUrl := `http://localhost/[0-9]+\?winningScore=\{WinningScore\}`
		if matches, _ := regexp.MatchString(expectedUrl, action.Url); !matches {
			t.Fatalf(`expected '%v' but got '%v'`, expectedUrl, action.Url)
		}
		t.Logf(`url: %v`, action.Url)
	}
	jsonBytes, _ := json.Marshal(gameManager)
	t.Logf(string(jsonBytes))
}

func TestGenerateNewGame(t *testing.T) {
	t.Skip(`skipping to do some refactoring`)
	gameManager := server.NewGameManager(`http://localhost`)
	gameManager.ActiveActions.NewGameAction(10000)
	if gameManager.ActiveGame == nil {
		t.Fatalf(`expected non-nil game but got nil`)
	}
	if gameManager.ActiveActions.NewGameAction == nil {
		t.Fatalf(`expected non-nil new game action but got nil`)
	}
	if gameManager.ActiveActions.NewTurnAction == nil {
		t.Fatalf(`expected non-nil new game action but got nil`)
	}
}
