package server

type Roller interface {
	Roll(int) []int
}

type options struct {
	overrideRollWith Roller
}

type OptionSetter func(*options) *options

func OverrideRollWith(roller Roller) OptionSetter {
	return func(o *options) *options {
		o.overrideRollWith = roller
		return o
	}
}

func NewGame(winningScore int, optionSetters ...OptionSetter) *Game {
	o := &options{}
	for _, setter := range optionSetters {
		o = setter(o)
	}

	g := &Game{}
	if o.overrideRollWith != nil {
		g.roller = o.overrideRollWith.Roll
	} else {
		g.roller = RollDice
	}

	g.winningScore = winningScore
	g.Roll = g.generateRollAction(6)
	return g
}

type RollAndScore struct {
	Roll        []int
	Score       int
	ScoringDice []int
}

type Game struct {
	winningScore     int
	scoredTurns      []int
	currentTurnScore int
	lastRoll         []int
	lastScore        ScoredRoll
	roller           func(int) []int

	Roll    func()
	NewTurn func()
}

func (g *Game) NewGame() {

}

func (g *Game) newTurn() {

}

func (g *Game) LastRolledScore() RollAndScore {
	return RollAndScore{
		Roll:        g.lastRoll,
		Score:       g.lastScore.Score,
		ScoringDice: g.lastScore.ScoringDice,
	}
}

func (g *Game) CurrentTurnScore() int {
	return g.currentTurnScore
}

func (g *Game) generateRollAction(totalDice int) func() {
	return func() {
		g.lastRoll = g.roller(totalDice)
		g.lastScore = Score(g.lastRoll)
		g.currentTurnScore += g.lastScore.Score
		if g.lastScore.Score == 0 {
			g.currentTurnScore = 0
			g.Roll = nil
			return
		}

		g.NewTurn = g.newTurn
		diceRemaining := totalDice - len(g.lastScore.ScoringDice)
		if diceRemaining == 0 {
			diceRemaining = 6
		}
		g.Roll = g.generateRollAction(diceRemaining)
	}
}
