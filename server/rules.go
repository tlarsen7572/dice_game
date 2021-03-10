package server

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func RollDice(number int) []int {
	roll := make([]int, number)
	for _, index := range roll {
		roll[index] = r.Intn(6) + 1
	}
	return roll
}
