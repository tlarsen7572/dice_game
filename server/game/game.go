package game

import (
	"server/rules"
	"server/turn"
)

func NewGame(winningScore int) *Game {
	return &Game{
		WinningScore: winningScore,
		CurrentScore: 0,
		Roller:       rules.RollDice,
		Turns:        []int{},
		ActiveTurn: &turn.Turn{
			Score:           0,
			LastRoll:        []int{},
			LastScoringDice: []int{},
		},
	}
}

type Game struct {
	WinningScore int
	CurrentScore int
	Roller       func(int) []int `json:"-"`
	Turns        []int
	ActiveTurn   *turn.Turn
}

func (g *Game) NewTurn() {
	g.CurrentScore += g.ActiveTurn.Score
	g.Turns = append(g.Turns, g.ActiveTurn.Score)
	g.ActiveTurn.Reset()
}

func (g *Game) Roll() {
	g.createTurnIfNil()
	diceToRoll := len(g.ActiveTurn.LastRoll)
	if diceToRoll == 0 {
		diceToRoll = 6
	}
	diceToRoll = diceToRoll - len(g.ActiveTurn.LastScoringDice)
	roll := g.Roller(diceToRoll)
	g.ActiveTurn.Roll(roll)
}

func (g *Game) createTurnIfNil() {
	if g.ActiveTurn == nil {
		g.ActiveTurn = &turn.Turn{}
	}
}
