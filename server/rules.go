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
	return 100, []int{0}
}
