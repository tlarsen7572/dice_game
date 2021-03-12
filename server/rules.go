package server

import (
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
	score := 0
	scoringDice := []int{}
	for index, die := range roll {
		if die == 1 {
			score += 100
			scoringDice = append(scoringDice, index)
		}
	}
	return score, scoringDice
}
