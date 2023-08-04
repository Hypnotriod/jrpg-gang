package domain

import (
	"jrpg-gang/util"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

type UnitCode string

type UnitClass string

const (
	UnitClassEmpty UnitClass = ""
)

type Unit struct {
	rng          *rand.Rand
	Uid          uint                     `json:"uid,omitempty" bson:"-"`
	Name         string                   `json:"name" bson:"name"`
	Code         UnitCode                 `json:"code,omitempty" bson:"code,omitempty"`
	Class        UnitClass                `json:"class,omitempty" bson:"class,omitempty"`
	Booty        UnitBooty                `json:"booty" bson:"booty"`
	State        UnitState                `json:"state" bson:"-"`
	Stats        UnitStats                `json:"stats" bson:"stats"`
	Damage       []DamageImpact           `json:"damage,omitempty" bson:"-"`
	Modification []UnitModificationImpact `json:"modification,omitempty" bson:"-"`
	Inventory    UnitInventory            `json:"inventory" bson:"inventory"`
	Slots        map[EquipmentSlot]uint   `json:"slots" bson:"slots"`
	Position     Position                 `json:"position" bson:"-"`
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

func (u *Unit) TotalActionPoints() float32 {
	var actionPoints float32 = u.Stats.BaseAttributes.ActionPoints
	actionPoints += float32(u.Stats.Attributes.Initiative / 10)
	for _, ench := range u.Modification {
		actionPoints += ench.BaseAttributes.ActionPoints
	}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			actionPoints += ench.BaseAttributes.ActionPoints
		}
	}
	return util.Max(actionPoints, MINIMUM_BASE_ATTRIBUTE_ACTION_POINTS)
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

func (u *Unit) TotalUnitModification() *UnitModification {
	var modification *UnitModification = &UnitModification{}
	for _, ench := range u.Modification {
		modification.Accumulate(ench.UnitModification)
	}
	return modification
}

func (u *Unit) TotalEquipmentModification() *UnitModification {
	var modification *UnitModification = &UnitModification{}
	for _, item := range u.Inventory.GetEquipment(true) {
		for _, ench := range item.Modification {
			modification.Accumulate(ench)
		}
	}
	return modification
}

func (u *Unit) TotalModification() *UnitModification {
	modification := u.TotalUnitModification()
	modification.Accumulate(*u.TotalEquipmentModification())
	return modification
}

func (u *Unit) ReduceDamageImpact(damage Damage) {
	if len(u.Damage) == 0 || !damage.HasEffect() {
		return
	}
	for i := range u.Damage {
		d := u.Damage[i].Damage
		u.Damage[i].Reduce(damage)
		u.Damage[i].Normalize()
		damage.Reduce(d)
		damage.Normalize()
	}
	u.FilterDamageImpact()
}

func (u *Unit) FilterDamageImpact() {
	filteredDamage := []DamageImpact{}
	for i := range u.Damage {
		if u.Damage[i].HasEffect() {
			filteredDamage = append(filteredDamage, u.Damage[i])
		}
	}
	u.Damage = filteredDamage
}

func (u *Unit) CheckRequirements(requirements UnitRequirements) bool {
	attributes := u.TotalModification().Attributes
	attributes.Accumulate(u.Stats.Attributes)
	attributes.Normalize()
	return requirements.Check(u.Class, attributes)
}

func (u *Unit) CheckUseCost(useCost UnitBaseAttributes) bool {
	return u.State.Health >= useCost.Health &&
		u.State.Mana >= useCost.Mana &&
		u.State.Stamina >= useCost.Stamina &&
		u.State.ActionPoints >= useCost.ActionPoints
}

func (u *Unit) CheckRandomChance(percents float32) bool {
	if u.rng == nil {
		u.rng = rand.New(mt19937.New())
		u.rng.Seed(time.Now().UnixNano())
	}
	rnd := u.rng.Int() % int(MAXIMUM_CHANCE)
	return float32(rnd) < percents
}

func (u *Unit) PickDeviation(deviation float32) float32 {
	if u.rng == nil {
		u.rng = rand.New(mt19937.New())
		u.rng.Seed(time.Now().UnixNano())
	}
	rnd := u.rng.Int() % int(deviation+1)
	if deviation < 0 {
		return float32(rnd * -1)
	}
	return float32(rnd)
}

func (u *Unit) SetRng(rng *rand.Rand) {
	u.rng = rng
}

func (u *Unit) CanReach(target *Unit, actionRange ActionRange) bool {
	return !actionRange.Has() || actionRange.Check(u.Position, target.Position)
}

func (u *Unit) CanUseWeapon(weapon *Weapon, checkUseCost bool) bool {
	if checkUseCost && !u.CheckUseCost(weapon.UseCost) {
		return false
	}
	if weapon.RequiresAmmunition() {
		ammunition := u.Inventory.FindEquippedAmmunition()
		quantity := 1 + len(weapon.Spread)
		if ammunition == nil || ammunition.Quantity < uint(quantity) {
			return false
		}
	}
	if weapon.IsBroken() {
		return false
	}
	return true
}

func (u *Unit) ReduceActionPoints(points float32) {
	u.State.ActionPoints -= points
	if u.State.ActionPoints < 0 {
		u.State.ActionPoints = 0
	}
}

func (u *Unit) ClearActionPoints() {
	u.State.ActionPoints = 0
}

func (u *Unit) CalculateCritilalAttackChance(target *Unit) float32 {
	chance := (u.TotalLuck() - u.State.Stress) - (target.TotalLuck() - target.State.Stress)
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateAttackChance(target *Unit, damage DamageImpact) float32 {
	if target.State.IsStunned {
		chance := (u.TotalAgility() - u.State.Stress) + target.State.Stress + damage.Chance
		return util.Max(chance, MINIMUM_CHANCE)
	}
	chance := (u.TotalAgility() - u.State.Stress) - (target.TotalAgility() - target.State.Stress) + damage.Chance
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateModificationChance(modification UnitModificationImpact) float32 {
	chance := (u.TotalIntelligence() - u.State.Stress) + modification.Chance
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateStunChance(target *Unit, damage Damage) float32 {
	chance := (damage.PhysicalDamage() - u.State.Stress) - (target.TotalPhysique() - target.State.Stress)
	return util.Max(chance, MINIMUM_CHANCE)
}

func (u *Unit) CalculateRetreatChance() float32 {
	chance := u.State.Stress
	return util.Max(chance, 0)
}

func (u *Unit) CalculateCriticalMissChance() float32 {
	chance := u.State.Stress
	return util.Max(chance, 0)
}
