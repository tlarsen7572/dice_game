package server

func NewGame(winningScore int) *Game {
	return &Game{
		WinningScore: winningScore,
		CurrentScore: 0,
		Roller:       RollDice,
		Turns:        []int{},
		ActiveTurn: &Turn{
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
