package util

import (
	"encoding/hex"
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

func RandomId() string {
	data := make([]byte, 16)
	for n := 0; n < len(data); n += 4 {
		u := rng.Uint32()
		data[n+0] = byte(u >> 24)
		data[n+1] = byte(u >> 16)
		data[n+2] = byte(u >> 8)
		data[n+3] = byte(u >> 0)
	}
	return hex.EncodeToString(data[:])
}
