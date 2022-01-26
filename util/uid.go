package util

type UidGen struct {
	uidCounter uint
}

func (g *UidGen) Next() uint {
	g.uidCounter++
	return g.uidCounter
}
