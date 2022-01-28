package util

import (
	"encoding/hex"
)

type UidGen struct {
	uidCounter uint
}

func (g *UidGen) Next() uint {
	g.uidCounter++
	return g.uidCounter
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
