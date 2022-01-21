package engine

import "jrpg-gang/domain"

type Battlefield struct {
	Cells [][]Cell `json:"cells"`
}

func NewBattlefield(size Size) *Battlefield {
	bf := &Battlefield{}
	for x := uint(0); x < size.Width; x++ {
		bf.Cells[x] = make([]Cell, size.Height)
		for y := uint(0); y < size.Height; y++ {
			bf.Cells[x][y].Position.X = int(x)
			bf.Cells[x][y].Position.Y = int(y)
		}
	}
	return bf
}

func (bf *Battlefield) PlaceUnitOnCell(config CellConfiguration, unit *domain.Unit) {
	cell := bf.GetCell(config.Position)
	cell.Unit = unit
	cell.CellConfiguration = config
}

func (bf *Battlefield) MoveUnit(fromCell *Cell, toCell *Cell) {
	toCell.CellConfiguration = fromCell.CellConfiguration
	toCell.Unit = fromCell.Unit
	fromCell.CellConfiguration = CellConfiguration{}
	fromCell.Unit = nil
}

func (bf *Battlefield) GetCell(position Position) *Cell {
	return &bf.Cells[position.X][position.Y]
}
