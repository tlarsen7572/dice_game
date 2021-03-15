package turn

import "server/rules"

type Turn struct {
	Score           int
	LastRoll        []int
	LastScoringDice []int
}

func (t *Turn) Roll(roll []int) {
	t.LastRoll = roll
	score := rules.Score(roll)
	t.LastScoringDice = score.ScoringDice
	t.Score += score.Score
	if score.Score == 0 {
		t.Score = 0
	}
}

func (t *Turn) Reset() {
	t.Score = 0
	t.LastScoringDice = []int{}
	t.LastRoll = []int{}
}
