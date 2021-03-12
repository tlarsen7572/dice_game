package server

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func RollDice(number int) []int {
	roll := make([]int, number)
	for _, index := range roll {
		roll[index] = r.Intn(6) + 1
	}
	sort.Ints(roll)
	return roll
}

func Score(roll []int) (int, []int) {
	ones := 0
	score := 0
	scoringDice := []int{}
	for index, die := range roll {
		if die == 1 {
			ones++
			scoringDice = append(scoringDice, index)
		}
	}
	if ones >= 3 {
		if ones == 6 {
			return math.MaxInt64, scoringDice
		}
		score += 1000
		score += (ones - 3) * 100
	} else {
		score += ones * 100
	}

	return score, scoringDice
}
