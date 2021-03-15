package mock_roller

import "server/rules"

type MockRoller struct {
	RollOverride  []int
	RequestedDice int
}

func (r *MockRoller) Roll(totalDice int) []int {
	r.RequestedDice = totalDice
	if r.RollOverride == nil {
		return rules.RollDice(totalDice)
	}
	return r.RollOverride
}
