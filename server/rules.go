package server

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func RollDice(number int) []int {
	roll := make([]int, number)
	for _, index := range roll {
		roll[index] = r.Intn(6) + 1
	}
	sort.Ints(roll)
	return roll
}

func Score(roll []int) ScoredRoll {
	calculator := &scoreCalculator{DiceValues: map[int]int{}}
	calculator.calculateDice(roll)
	return calculator.scoredRoll()
}

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

const longGameScore = 10000
const allDice = 6

func rollEquals(roll []int, expected []int) bool {
	if len(roll) != len(expected) {
		return false
	}
	for index, value := range roll {
		if value != expected[index] {
			return false
		}
	}
	return true
}

func rollIsThreeStraights(roll []int) bool {
	if len(roll) != allDice {
		return false
	}
	if roll[0] == roll[1] &&
		roll[2] == roll[3] &&
		roll[4] == roll[5] &&
		roll[0] != roll[2] &&
		roll[0] != roll[4] {
		return true
	}
	return false
}

type ScoredRoll struct {
	Score       int
	ScoringDice []int
}

type scoreCalculator struct {
	DiceValues  map[int]int
	Score       int
	ScoringDice []int
}

func (c *scoreCalculator) calculateDice(roll []int) {
	if rollEquals(roll, []int{1, 2, 3, 4, 5, 6}) {
		c.ScoringDice = []int{0, 1, 2, 3, 4, 5}
		c.Score = 1000
		return
	}
	if rollIsThreeStraights(roll) {
		c.ScoringDice = []int{0, 1, 2, 3, 4, 5}
		c.Score = 1000
		return
	}

	for index, die := range roll {
		switch die {
		case 1:
			c.processOne(index)
		case 5:
			c.processFive(index)
		case 2, 3, 4, 6:
			c.processOther(index, die)
		default:
			panic(fmt.Sprintf(`invalid dice value %v`, die))
		}
	}
}

func (c *scoreCalculator) processOne(index int) {
	c.DiceValues[1]++
	ones := c.DiceValues[1]
	c.Score += 100
	c.ScoringDice = append(c.ScoringDice, index)
	if ones%3 == 0 {
		c.Score += 700 // 700 + previous 3 values of 100 = 1000, the value of a triplet of ones
	}
	if ones%6 == 0 {
		c.Score = longGameScore // make score maximum possible to force the game to end
		return
	}
}

func (c *scoreCalculator) processFive(index int) {
	c.DiceValues[5]++
	fives := c.DiceValues[5]
	c.Score += 50
	c.ScoringDice = append(c.ScoringDice, index)
	if fives%3 == 0 {
		c.Score += 350 // 350 + previous 3 values of 150 = 500, the value of a triplet of fives
	}
}

func (c *scoreCalculator) processOther(index int, diceValue int) {
	c.DiceValues[diceValue]++
	totalCount := c.DiceValues[diceValue]
	if totalCount%3 == 0 {
		c.ScoringDice = append(c.ScoringDice, index-2, index-1, index)
		c.Score += diceValue * 100
	}
}

func (c *scoreCalculator) scoredRoll() ScoredRoll {
	return ScoredRoll{
		Score:       c.Score,
		ScoringDice: c.ScoringDice,
	}
}
