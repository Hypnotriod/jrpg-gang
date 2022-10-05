package domain

import (
	"fmt"
	"jrpg-gang/util"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

type UnitSlots map[EquipmentSlot]uint

type UnitCode string

type Unit struct {
	rng          *rand.Rand
	Uid          uint                     `json:"uid,omitempty"`
	Name         string                   `json:"name"`
	Code         UnitCode                 `json:"code,omitempty"`
	Booty        UnitBooty                `json:"booty"`
	State        UnitState                `json:"state"`
	Stats        UnitStats                `json:"stats"`
	Damage       []DamageImpact           `json:"damage"`
	Modification []UnitModificationImpact `json:"modification"`
	Inventory    UnitInventory            `json:"inventory"`
	Slots        UnitSlots                `json:"slots"`
	Position     Position                 `json:"position"`
}

func (u Unit) String() string {
	return fmt.Sprintf(
		"%s, code:%s, booty: {%v}, state: {%v}, stats: {%v}, damage: [%s], modification: [%s], inventory: {%v}, slots: %v, position: {%v}, uid: %d",
		u.Name,
		u.Code,
		u.Booty,
		u.State,
		u.Stats,
		util.AsCommaSeparatedObjectsSlice(u.Damage),
		util.AsCommaSeparatedObjectsSlice(u.Modification),
		u.Inventory,
		u.Slots,
		u.Position,
		u.Uid,
	)
}

func (u *Unit) Clone() *Unit {
	r := &Unit{}
	r.rng = u.rng
	r.Uid = u.Uid
	r.Name = u.Name
	r.Code = u.Code
	r.Booty = u.Booty
	r.State = u.State
	r.Stats = u.Stats
	r.Damage = []DamageImpact{}
	r.Damage = append(r.Damage, u.Damage...)
	r.Modification = []UnitModificationImpact{}
	r.Modification = append(r.Modification, u.Modification...)
	r.Inventory = *u.Inventory.Clone()
	r.Slots = util.CloneMap(u.Slots)
	r.Position = u.Position
	return r
}

func (u *Unit) ClearImpact() {
	u.Damage = []DamageImpact{}
	u.Modification = []UnitModificationImpact{}
}

func (u *Unit) TotalAgility() float32 {
	var agility float32 = u.Stats.Attributes.Agility
	for _, ench := range u.Modification {
		agility += ench.Attributes.Agility
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			agility += ench.Attributes.Agility
		}
	}
	return util.Max(agility, 0)
}

func (u *Unit) TotalPhysique() float32 {
	var physique float32 = u.Stats.Attributes.Physique
	for _, ench := range u.Modification {
		physique += ench.Attributes.Physique
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			physique += ench.Attributes.Physique
		}
	}
	return util.Max(physique, 0)
}

func (u *Unit) TotalIntelligence() float32 {
	var intelligence float32 = u.Stats.Attributes.Intelligence
	for _, ench := range u.Modification {
		intelligence += ench.Attributes.Intelligence
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			intelligence += ench.Attributes.Intelligence
		}
	}
	return util.Max(intelligence, 0)
}

func (u *Unit) TotalLuck() float32 {
	var luck float32 = u.Stats.Attributes.Luck
	for _, ench := range u.Modification {
		luck += ench.Attributes.Luck
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			luck += ench.Attributes.Luck
		}
	}
	return util.Max(luck, 0)
}

func (u *Unit) TotalInitiative() float32 {
	if u.State.IsStunned {
		return -1
	}
	var initiative float32 = u.Stats.Attributes.Initiative
	for _, ench := range u.Modification {
		initiative += ench.Attributes.Initiative
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			initiative += ench.Attributes.Initiative
		}
	}
	return util.Max(initiative, 0)
}

func (u *Unit) TotalModification() *UnitModification {
	var modification *UnitModification = &UnitModification{}
	for _, ench := range u.Modification {
		modification.Accumulate(ench.UnitModification)
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			modification.Accumulate(ench)
		}
	}
	return modification
}

func (u *Unit) CheckRequirements(requirements UnitAttributes) bool {
	attributes := u.TotalModification().Attributes
	attributes.Accumulate(u.Stats.Attributes)
	attributes.Normalize()
	return attributes.Strength >= requirements.Strength &&
		attributes.Physique >= requirements.Physique &&
		attributes.Agility >= requirements.Agility &&
		attributes.Endurance >= requirements.Endurance &&
		attributes.Intelligence >= requirements.Intelligence &&
		attributes.Initiative >= requirements.Initiative &&
		attributes.Luck >= requirements.Luck
}

func (u *Unit) CheckUseCost(useCost UnitBaseAttributes) bool {
	return u.State.Health >= useCost.Health &&
		u.State.Mana >= useCost.Mana &&
		u.State.Stamina >= useCost.Stamina
}

func (u *Unit) CheckRandomChance(percents float32) bool {
	if u.rng == nil {
		u.rng = rand.New(mt19937.New())
		u.rng.Seed(time.Now().UnixNano())
	}
	rnd := u.rng.Float32() * MAXIMUM_CHANCE
	return percents > rnd
}

func (u *Unit) SetRng(rng *rand.Rand) {
	u.rng = rng
}

func (u *Unit) CanReachWithWeapon(target *Unit, weapon *Weapon) bool {
	return weapon.Range.Check(u.Position, target.Position)
}

func (u *Unit) CalculateCritilalAttackChance(target *Unit) float32 {
	chance := (u.TotalLuck() - u.State.Curse) - (target.TotalLuck() - target.State.Curse)
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateAttackChance(target *Unit, damage DamageImpact) float32 {
	chance := (u.TotalAgility() - u.State.Curse) - (target.TotalAgility() - target.State.Curse) + damage.Chance
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateModificationChance(modification UnitModificationImpact) float32 {
	chance := (u.TotalIntelligence() - u.State.Curse) + modification.Chance
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateStunChance(target *Unit, damage Damage) float32 {
	chance := (damage.PhysicalDamage() - target.State.Curse) - (u.TotalPhysique() - u.State.Curse)
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateRetreatChance() float32 {
	chance := u.State.Fear
	return util.Max(chance, 0)
}
