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
	b.Units = []*GameUnit{}
	return b
}

func (b *Battlefield) PlaceUnit(unit *GameUnit, position domain.Position) *domain.ActionResult {
	result := domain.ActionResult{}
	if b.FindUnitById(unit.Uid) != nil {
		return result.WithResultType(domain.ResultNotAllowed)
	}
	if !b.checkPositionBounds(position) {
		return result.WithResultType(domain.ResultOutOfBounds)
	}
	if !b.checkPositionCanPlaceUnit(position) || !b.checkPositionFraction(position, unit.FractionId) {
		return result.WithResultType(domain.ResultNotAccomplished)
	}
	unitAtPosition := b.FindUnitByPosition(position)
	if unitAtPosition != nil {
		return result.WithResultType(domain.ResultNotEmpty)
	}
	unit.Position = position
	b.Units = append(b.Units, unit)
	return result.WithResultType(domain.ResultAccomplished)
}

func (b *Battlefield) MoveUnit(uid uint, position domain.Position) *domain.ActionResult {
	result := domain.ActionResult{}
	if !b.checkPositionBounds(position) {
		return result.WithResultType(domain.ResultOutOfBounds)
	}
	unit := b.FindUnitById(uid)
	if !b.checkPositionCanPlaceUnit(position) || !b.checkPositionFraction(position, unit.FractionId) {
		return result.WithResultType(domain.ResultNotAccomplished)
	}
	unitAtPosition := b.FindUnitByPosition(position)
	if unit == nil {
		return result.WithResultType(domain.ResultNotFound)
	}
	if unitAtPosition != nil {
		return result.WithResultType(domain.ResultNotEmpty)
	}
	unit.Position = position
	return result.WithResultType(domain.ResultAccomplished)
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

func (b *Battlefield) ContainsUnits(units []*GameUnit) int {
	var count int = 0
	for _, u1 := range units {
		for _, u2 := range b.Units {
			if u1.Uid == u2.Uid {
				count++
				break
			}
		}
	}
	return count
}

func (b *Battlefield) FractionsLeft() int {
	var fractions map[uint]struct{} = map[uint]struct{}{}
	for _, unit := range b.Units {
		fractions[unit.FractionId] = struct{}{}
	}
	return len(fractions)
}

func (b *Battlefield) FindReachableTargets(unit *GameUnit) map[uint]*GameUnit {
	result := make(map[uint]*GameUnit)
	for i := range unit.Inventory.Weapon {
		weapon := &unit.Inventory.Weapon[i]
		for _, target := range b.Units {
			if target.FractionId != unit.FractionId && unit.CanReachWithWeapon(&target.Unit, weapon) {
				result[weapon.Uid] = target
			}
		}
	}
	return result
}
