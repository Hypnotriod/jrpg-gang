package engine

type Battlefield struct {
	Size  Size        `json:"size"`
	Units []*GameUnit `json:"units"`
}

func NewBattlefield(size Size) *Battlefield {
	b := &Battlefield{}
	b.Size = size
	return b
}

func (b *Battlefield) PlaceUnit(unit *GameUnit) bool {
	unitAtPosition := b.FindUnitByPosition(unit.Position)
	if unitAtPosition == nil {
		b.Units = append(b.Units, unit)
		return true
	}
	return false
}

func (b *Battlefield) MoveUnit(uid uint, toPosition Position) bool {
	unit := b.FindUnitById(uid)
	unitAtPosition := b.FindUnitByPosition(toPosition)
	if unit != nil && unitAtPosition == nil {
		unit.Position = toPosition
		return true
	}
	return false
}

func (b *Battlefield) FindUnitById(uid uint) *GameUnit {
	for i := 0; i < len(b.Units); i++ {
		if b.Units[i].Uid == uid {
			return b.Units[i]
		}
	}
	return nil
}

func (b *Battlefield) FindUnitByPosition(position Position) *GameUnit {
	for i := 0; i < len(b.Units); i++ {
		if b.Units[i].Position.Equals(&position) {
			return b.Units[i]
		}
	}
	return nil
}
