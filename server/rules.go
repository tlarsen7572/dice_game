package server

import "math/rand"

func RollDice(number int) []int {
	roll := make([]int, number)
	for _, index := range roll {
		roll[index] = rand.Intn(6) + 1
	}
	return roll
}
