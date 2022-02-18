package util

import (
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

type RndGen struct {
	uidCounter uint
	rng        *rand.Rand
}

func NewRndGen() *RndGen {
	g := &RndGen{}
	g.rng = rand.New(mt19937.New())
	g.rng.Seed(time.Now().UnixNano())
	return g
}

func (g *RndGen) NextUid() uint {
	g.uidCounter++
	return g.uidCounter
}

func (g *RndGen) Hash() string {
	data := make([]byte, 16)
	for n := 0; n < len(data); n += 4 {
		u := g.rng.Uint32()
		data[n+0] = byte(u >> 24)
		data[n+1] = byte(u >> 16)
		data[n+2] = byte(u >> 8)
		data[n+3] = byte(u >> 0)
	}
	return hex.EncodeToString(data[:])
}

func (g *RndGen) PickIndex(n int) int {
	return g.rng.Int() % n
}

func (g *RndGen) PickInt(values []int) int {
	index := g.rng.Int() % len(values)
	return values[index]
}

func (g *RndGen) PickIntByWeight(valueWight map[int]int) int {
	var sum int = 0
	var any int
	for _, w := range valueWight {
		sum += w
	}
	limit := g.rng.Int() % sum
	sum = 0
	for v, w := range valueWight {
		any = v
		sum += w
		if limit < sum {
			return v
		}
	}
	return any
}
