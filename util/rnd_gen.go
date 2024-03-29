package util

import (
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/lithammer/shortuuid/v4"
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

func (g *RndGen) MakeUUID() string {
	return shortuuid.New()
}

func (g *RndGen) MakeUUIDWithUniquenessCheck(isUnique func(value string) bool) string {
	value := shortuuid.New()
	for !isUnique(value) {
		value = shortuuid.New()
	}
	return value
}

func (g *RndGen) NextUid() uint {
	g.uidCounter++
	return g.uidCounter
}

func (g *RndGen) ResetUidGen() {
	g.uidCounter = 0
}

func (g *RndGen) MakeHex16() string {
	return g.makeHex(16)
}

func (g *RndGen) MakeHex32() string {
	return g.makeHex(32)
}

func (g *RndGen) makeHex(size int) string {
	data := make([]byte, size/2)
	for n := 0; n < len(data); n += 4 {
		u := g.rng.Uint32()
		data[n+0] = byte(u >> 24)
		data[n+1] = byte(u >> 16)
		data[n+2] = byte(u >> 8)
		data[n+3] = byte(u >> 0)
	}
	return hex.EncodeToString(data)
}

func (g *RndGen) PickIndex(n int) int {
	return g.rng.Int() % n
}

func (g *RndGen) PickInt(values []int) int {
	index := g.rng.Int() % len(values)
	return values[index]
}
