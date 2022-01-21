package engine

type Battlefield struct {
	Cells [][]Cell `json:"cells"`
}

func NewBattlefield(size Size) *Battlefield {
	b := &Battlefield{}
	for x := uint(0); x < size.Width; x++ {
		b.Cells[x] = make([]Cell, size.Height)
		for y := uint(0); y < size.Height; y++ {
			b.Cells[x][y].Position.X = int(x)
			b.Cells[x][y].Position.Y = int(y)
		}
	}
	return b
}

func (b *Battlefield) PlaceUnitOnCell(config CellConfiguration, unit *GameUnit) {
	cell := b.GetCell(config.Position)
	cell.Unit = unit
	cell.CellConfiguration = config
}

func (b *Battlefield) MoveUnit(fromCell *Cell, toCell *Cell) {
	toCell.CellConfiguration = fromCell.CellConfiguration
	toCell.Unit = fromCell.Unit
	fromCell.CellConfiguration = CellConfiguration{}
	fromCell.Unit = nil
}

func (b *Battlefield) FindUnit(uid uint) *GameUnit {
	for _, cells := range b.Cells {
		for _, cell := range cells {
			if cell.Unit != nil && cell.Unit.Uid == uid {
				return cell.Unit
			}
		}
	}
	return nil
}

func (b *Battlefield) GetCell(position Position) *Cell {
	return &b.Cells[position.X][position.Y]
}
