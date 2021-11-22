package util

import "math/rand"

func CheckRandomChance(percents float32) bool {
	rnd := rand.Float32() * 100
	return percents > rnd
}
