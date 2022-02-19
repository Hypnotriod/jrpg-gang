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
	if !b.checkPositionCanPlaceUnit(position) || !b.checkPositionFaction(position, unit.Faction) {
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
	unit := b.FindUnitById(uid)
	if unit == nil {
		return result.WithResultType(domain.ResultNotFound)
	}
	if !b.checkPositionBounds(position) {
		return result.WithResultType(domain.ResultOutOfBounds)
	}
	if !b.checkPositionCanPlaceUnit(position) || !b.checkPositionFaction(position, unit.Faction) {
		return result.WithResultType(domain.ResultNotAccomplished)
	}
	unitAtPosition := b.FindUnitByPosition(position)
	if unitAtPosition != nil {
		return result.WithResultType(domain.ResultNotEmpty)
	}
	unit.Position = position
	b.UpdateCellsFactions()
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

func (b *Battlefield) CanMoveUnitTo(unit *GameUnit, position domain.Position) bool {
	return b.checkPositionBounds(position) &&
		b.checkPositionFaction(position, unit.Faction) &&
		b.checkPositionCanPlaceUnit(position)
}

func (b *Battlefield) checkPositionBounds(position domain.Position) bool {
	return position.X >= 0 && position.Y >= 0 && position.X < len(b.Matrix) && position.Y < len(b.Matrix[0])
}

func (b *Battlefield) checkPositionFaction(position domain.Position, faction GameUnitFaction) bool {
	return b.Matrix[position.X][position.Y].ContainsFaction(faction)
}

func (b *Battlefield) checkPositionCanPlaceUnit(position domain.Position) bool {
	return b.Matrix[position.X][position.Y].CanPlaceUnit()
}

func (b *Battlefield) UpdateCellsFactions() {
	leftBound := -1
	rightBound := len(b.Matrix[0])
	for _, unit := range b.Units {
		if unit.Faction == GameUnitFactionLeft {
			leftBound = unit.Position.X
		} else {
			rightBound = unit.Position.X
		}
	}
	for x := range b.Matrix {
		for y := range b.Matrix[x] {
			if x <= leftBound {
				b.Matrix[x][y].Factions = []GameUnitFaction{GameUnitFactionLeft}
			} else if x >= rightBound {
				b.Matrix[x][y].Factions = []GameUnitFaction{GameUnitFactionRight}
			} else {
				b.Matrix[x][y].Factions = []GameUnitFaction{GameUnitFactionLeft, GameUnitFactionRight}
			}
		}
	}
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

func (b *Battlefield) FactionsCount() int {
	factions := map[GameUnitFaction]struct{}{}
	for _, unit := range b.Units {
		factions[unit.Faction] = struct{}{}
	}
	return len(factions)
}

func (b *Battlefield) FindReachableTargets(unit *GameUnit) map[uint]*GameUnit {
	result := map[uint]*GameUnit{}
	for i := range unit.Inventory.Weapon {
		weapon := &unit.Inventory.Weapon[i]
		for _, target := range b.Units {
			if target.Faction != unit.Faction && unit.CanReachWithWeapon(&target.Unit, weapon) {
				result[weapon.Uid] = target
			}
		}
	}
	return result
}
