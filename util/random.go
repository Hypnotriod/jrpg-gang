package util

import (
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

const MINIMUM_CHANCE float32 = 1
const MAXIMUM_CHANCE float32 = 100

var rng *rand.Rand = rand.New(mt19937.New())

func CheckRandomChance(percents float32) bool {
	rnd := rng.Float32() * MAXIMUM_CHANCE
	return percents > rnd
}

func ApplyRandomSeed() {
	rng.Seed(time.Now().UnixNano())
}
