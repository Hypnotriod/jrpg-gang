package util

import (
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

var rng *rand.Rand = rand.New(mt19937.New())

func CheckRandomChance(percents float32) bool {
	rnd := rng.Float32() * 100
	return percents > rnd
}

func ApplyRandomSeed() {
	rng.Seed(time.Now().UnixNano())
}
