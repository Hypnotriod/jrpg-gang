package util

import (
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

type UidGen struct {
	uidCounter uint
	rng        *rand.Rand
}

func NewUidGen() *UidGen {
	g := &UidGen{}
	g.rng = rand.New(mt19937.New())
	g.rng.Seed(time.Now().UnixNano())
	return g
}

func (g *UidGen) Next() uint {
	g.uidCounter++
	return g.uidCounter
}

func (g *UidGen) Hash() string {
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