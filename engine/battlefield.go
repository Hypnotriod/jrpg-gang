package engine

import "jrpg-gang/domain"

type Battlefield struct {
	Matrix [][]Cell    `json:"matrix"`
	Units  []*GameUnit `json:"units"`
}

func NewBattlefield(matrix [][]Cell) *Battlefield {
	b := &Battlefield{}
	b.Matrix = matrix
	return b
}

func (b *Battlefield) PlaceUnit(unit *GameUnit) domain.ActionResult {
	result := domain.ActionResult{}
	if !b.checkPositionBounds(unit.Position) {
		return *result.WithResultType(domain.ResultOutOfBounds)
	}
	if !b.checkPositionCanPlaceUnit(unit.Position) || !b.checkPositionFraction(unit.Position, unit.FractionId) {
		return *result.WithResultType(domain.ResultNotAccomplished)
	}
	unitAtPosition := b.FindUnitByPosition(unit.Position)
	if unitAtPosition != nil {
		return *result.WithResultType(domain.ResultNotEmpty)
	}
	b.Units = append(b.Units, unit)
	return *result.WithResultType(domain.ResultAccomplished)
}

func (b *Battlefield) MoveUnit(uid uint, position Position) domain.ActionResult {
	result := domain.ActionResult{}
	if !b.checkPositionBounds(position) {
		return *result.WithResultType(domain.ResultOutOfBounds)
	}
	unit := b.FindUnitById(uid)
	if !b.checkPositionCanPlaceUnit(unit.Position) || !b.checkPositionFraction(unit.Position, unit.FractionId) {
		return *result.WithResultType(domain.ResultNotAccomplished)
	}
	unitAtPosition := b.FindUnitByPosition(position)
	if unit == nil {
		return *result.WithResultType(domain.ResultNotFound)
	}
	if unitAtPosition != nil {
		return *result.WithResultType(domain.ResultNotEmpty)
	}
	unit.Position = position
	return *result.WithResultType(domain.ResultAccomplished)
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

func (b *Battlefield) checkPositionBounds(position Position) bool {
	return position.X >= 0 && position.Y >= 0 && position.X < len(b.Matrix) && position.Y < len(b.Matrix[0])
}

func (b *Battlefield) checkPositionFraction(position Position, fractionId uint) bool {
	return b.Matrix[position.X][position.Y].FractionId == fractionId
}

func (b *Battlefield) checkPositionCanPlaceUnit(position Position) bool {
	cell := b.Matrix[position.X][position.Y]
	return cell.Type == CellTypeSpace
}
