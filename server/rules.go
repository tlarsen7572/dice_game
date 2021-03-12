package server

import (
	"math/rand"
	"sort"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

const longGameScore = 10000

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
	twos := 0
	score := 0
	scoringDice := []int{}
	for index, die := range roll {
		if die == 1 {
			ones++
			scoringDice = append(scoringDice, index)
		}
		if die == 2 {
			twos++
			if twos%3 == 0 {
				scoringDice = append(scoringDice, index-2, index-1, index)
				score += 200
			}
		}
	}
	if ones == 6 {
		return longGameScore, scoringDice
	}
	setsOfThree := ones / 3
	onesRemainder := ones % 3
	score += 1000 * setsOfThree
	score += onesRemainder * 100

	return score, scoringDice
}
