package server

import (
	"encoding/json"
)

type Game struct {
	WinningScore int
	CurrentScore int
	Roller       func(int) []int `json:"-"`
	Turns        []int
	ActiveTurn   *Turn
}

func (g *Game) NewTurn() {
	g.CurrentScore += g.ActiveTurn.Score
	g.Turns = append(g.Turns, g.ActiveTurn.Score)
	g.ActiveTurn.Reset()
}

func (g *Game) Roll() {
	g.createTurnIfNil()
	diceToRoll := 6 - len(g.ActiveTurn.LastRoll)
	roll := g.Roller(diceToRoll)
	g.ActiveTurn.Roll(roll)
}

func (g *Game) createTurnIfNil() {
	if g.ActiveTurn == nil {
		g.ActiveTurn = &Turn{}
	}
}

func (g *Game) ToJson() string {
	jsonBytes, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	}
	return string(jsonBytes)
}
