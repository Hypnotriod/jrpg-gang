package engine

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type Battlefield struct {
	Matrix [][]Cell    `json:"matrix"`
	Units  []*GameUnit `json:"units"`
}

func (b Battlefield) String() string {
	return fmt.Sprintf(
		"matrix: %v, units: [%v]",
		b.Matrix,
		util.AsCommaSeparatedSlice(b.Units),
	)
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

func (b *Battlefield) MoveUnit(uid uint, position domain.Position) domain.ActionResult {
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

func (b *Battlefield) FindUnitByPosition(position domain.Position) *GameUnit {
	for i := 0; i < len(b.Units); i++ {
		if b.Units[i].Position.Equals(position) {
			return b.Units[i]
		}
	}
	return nil
}

func (b *Battlefield) checkPositionBounds(position domain.Position) bool {
	return position.X >= 0 && position.Y >= 0 && position.X < len(b.Matrix) && position.Y < len(b.Matrix[0])
}

func (b *Battlefield) checkPositionFraction(position domain.Position, fractionId uint) bool {
	return b.Matrix[position.X][position.Y].ContainsFractionId(fractionId)
}

func (b *Battlefield) checkPositionCanPlaceUnit(position domain.Position) bool {
	return b.Matrix[position.X][position.Y].CanPlaceUnit()
}

func (b *Battlefield) FilterSurvivors() {
	survivedUnits := []*GameUnit{}
	for _, unit := range b.Units {
		if unit.State.Health > 0 {
			survivedUnits = append(survivedUnits, unit)
		}
	}
	b.Units = survivedUnits
}
