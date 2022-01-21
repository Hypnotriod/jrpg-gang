package engine

type Battlefield struct {
	Cells [][]Cell `json:"cells"`
}

func NewBattlefield(width uint, height uint) *Battlefield {
	bf := &Battlefield{}
	for h := uint(0); h < height; h++ {
		bf.Cells[h] = make([]Cell, width)
	}
	return bf
}
