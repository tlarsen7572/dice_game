package server

import (
	"fmt"
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
		c.Score += 700
	}
	if ones%6 == 0 {
		c.Score = longGameScore
		return
	}
}

func (c *scoreCalculator) processFive(index int) {
	c.DiceValues[5]++
	fives := c.DiceValues[5]
	c.Score += 50
	c.ScoringDice = append(c.ScoringDice, index)
	if fives%3 == 0 {
		c.Score += 350
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

func Score(roll []int) ScoredRoll {
	calculator := &scoreCalculator{DiceValues: map[int]int{}}
	calculator.calculateDice(roll)
	return calculator.scoredRoll()
}
