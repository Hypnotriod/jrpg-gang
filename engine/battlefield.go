package engine

import (
	"fmt"
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type BattlefieldUnitDescriptor struct {
	Code     domain.UnitCode `json:"code"`
	Faction  GameUnitFaction `json:"faction"`
	Position domain.Position `json:"position"`
}

type Battlefield struct {
	Matrix         [][]Cell                    `json:"matrix"`
	UnitDescriptor []BattlefieldUnitDescriptor `json:"unitDescriptor,omitempty"`
	Units          []*GameUnit                 `json:"units"`
	Corpses        []*GameUnit                 `json:"corpses"`
}

func (b Battlefield) String() string {
	return fmt.Sprintf(
		"matrix: %v, units: [%v], corpses: [%s]",
		b.Matrix,
		util.AsCommaSeparatedObjectsSlice(b.Units),
		util.AsCommaSeparatedObjectsSlice(b.Corpses),
	)
}

func (b *Battlefield) Clone() *Battlefield {
	r := &Battlefield{}
	r.Matrix = make([][]Cell, len(b.Matrix))
	for x := range b.Matrix {
		for y := range b.Matrix[x] {
			r.Matrix[x] = append(r.Matrix[x], *b.Matrix[x][y].Clone())
		}
	}
	r.UnitDescriptor = append(r.UnitDescriptor, b.UnitDescriptor...)
	for i := range b.Units {
		r.Units = append(r.Units, b.Units[i].Clone())
	}
	for i := range b.Corpses {
		r.Corpses = append(r.Corpses, b.Corpses[i].Clone())
	}
	return r
}

func NewBattlefield(matrix [][]Cell) *Battlefield {
	b := &Battlefield{}
	b.Matrix = matrix
	b.Units = []*GameUnit{}
	b.Corpses = []*GameUnit{}
	return b
}

func (b *Battlefield) Dispose() {
	b.Units = nil
	b.Corpses = nil
}

func (b *Battlefield) PlaceUnit(unit *GameUnit, position domain.Position) *domain.ActionResult {
	result := domain.NewActionResult()
	if !b.checkPositionBounds(position) {
		return result.WithResult(domain.ResultOutOfBounds)
	}
	if !b.checkPositionCanPlaceUnit(position) || !b.checkPositionFaction(position, unit.Faction) {
		return result.WithResult(domain.ResultNotAccomplished)
	}
	unitAtPosition := b.FindUnitByPosition(position)
	if unitAtPosition != nil {
		return result.WithResult(domain.ResultNotEmpty)
	}
	unit.Position = position
	if b.FindUnit(unit.Uid) == nil {
		b.Units = append(b.Units, unit)
	}
	return result.WithResult(domain.ResultAccomplished)
}

func (b *Battlefield) MoveUnit(uid uint, position domain.Position) *domain.ActionResult {
	result := domain.NewActionResult()
	unit := b.FindUnit(uid)
	if unit == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	if !b.checkPositionBounds(position) {
		return result.WithResult(domain.ResultOutOfBounds)
	}
	if !b.checkPositionCanPlaceUnit(position) || !b.checkPositionFaction(position, unit.Faction) {
		return result.WithResult(domain.ResultNotAccomplished)
	}
	unitAtPosition := b.FindUnitByPosition(position)
	if unitAtPosition != nil {
		return result.WithResult(domain.ResultNotEmpty)
	}
	unit.Position = position
	b.UpdateCellsFactions()
	return result.WithResult(domain.ResultAccomplished)
}

func (b *Battlefield) MoveToCorpses(uid uint) {
	unit := b.FindUnit(uid)
	if unit == nil {
		return
	}
	survivedUnits := []*GameUnit{}
	for _, u := range b.Units {
		if u.Uid != uid {
			survivedUnits = append(survivedUnits, u)
		}
	}
	unit.IsDead = true
	b.Units = survivedUnits
	b.Corpses = append(b.Corpses, unit)
}

func (b *Battlefield) FindUnit(uid uint) *GameUnit {
	return util.Findp(b.Units, func(u *GameUnit) bool {
		return u.Uid == uid
	})
}

func (b *Battlefield) FindUnitByPosition(position domain.Position) *GameUnit {
	return util.Findp(b.Units, func(u *GameUnit) bool {
		return u.Position.Equals(position)
	})
}

func (b *Battlefield) CanMoveUnitTo(unit *GameUnit, position domain.Position) bool {
	return b.checkPositionBounds(position) &&
		b.checkPositionFaction(position, unit.Faction) &&
		b.checkPositionCanPlaceUnit(position) &&
		b.FindUnitByPosition(position) == nil
}

func (b *Battlefield) UpdateCellsFactions() {
	leftBound := -1
	rightBound := len(b.Matrix)
	for _, unit := range b.Units {
		if unit.Faction == GameUnitFactionLeft {
			leftBound = util.Max(unit.Position.X, leftBound)
		} else {
			rightBound = util.Min(unit.Position.X, rightBound)
		}
	}
	for x := range b.Matrix {
		for y := range b.Matrix[x] {
			if b.Matrix[x][y].Type == CellTypeObstacle {
				continue
			}
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

func (b *Battlefield) FilterSurvivors() []*GameUnit {
	newCorpses := []*GameUnit{}
	survivedUnits := []*GameUnit{}
	for _, unit := range b.Units {
		if unit.State.Health > 0 {
			survivedUnits = append(survivedUnits, unit)
		} else {
			unit.IsDead = true
			newCorpses = append(newCorpses, unit)
			b.Corpses = append(b.Corpses, unit)
		}
	}
	b.Units = survivedUnits
	return newCorpses
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

func (b *Battlefield) FactionUnitsCount(faction GameUnitFaction) int {
	result := 0
	for _, unit := range b.Units {
		if unit.Faction == faction {
			result++
		}
	}
	return result
}

func (b *Battlefield) GetUnitsByFaction(faction GameUnitFaction) []*GameUnit {
	return util.Filter(b.Units, func(unit *GameUnit) bool {
		return unit.Faction == faction
	})
}

func (b *Battlefield) FindReachableTargets(unit *GameUnit) map[uint]*GameUnit {
	result := map[uint]*GameUnit{}
	for i := range unit.Inventory.Weapon {
		weapon := &unit.Inventory.Weapon[i]
		for _, target := range b.Units {
			if target.Faction != unit.Faction && unit.CanReach(&target.Unit, weapon.Range) {
				result[weapon.Uid] = target
			}
		}
	}
	return result
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

func (b *Battlefield) tryToPlaceUnit(unit *GameUnit, position domain.Position) bool {
	if !b.checkPositionCanPlaceUnit(position) ||
		!b.checkPositionFaction(position, unit.Faction) ||
		b.FindUnitByPosition(position) != nil {
		return false
	}
	unit.Position = position
	if b.FindUnit(unit.Uid) == nil {
		b.Units = append(b.Units, unit)
	}
	return true
}

func (b *Battlefield) placeUnitDefault(unit *GameUnit) {
	var x, y int
	for x = 0; x < len(b.Matrix); x++ {
		for y = 0; y < len(b.Matrix[x]); y++ {
			position := domain.Position{X: x, Y: y}
			if b.tryToPlaceUnit(unit, position) {
				return
			}
		}
	}
}

func (b *Battlefield) placeUnitsDefault(units []*GameUnit) {
	for _, unit := range units {
		b.placeUnitDefault(unit)
	}
}
