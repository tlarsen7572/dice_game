package server_test

import (
	"server"
	"testing"
)

func TestNewGameActions(t *testing.T) {
	game := server.NewGame(10000)
	actions := server.GenerateActions(game)
	print(actions)
}
