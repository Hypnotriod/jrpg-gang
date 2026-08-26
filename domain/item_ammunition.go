package domain

type AmmunitionKind string

const (
	NoAmmunition AmmunitionKind = ""
)

type Ammunition struct {
	Item
	Equipped bool           `json:"equipped,omitzero"`
	Kind     AmmunitionKind `json:"kind"`
	Quantity uint           `json:"quantity,omitzero"`
	Damage   []DamageImpact `json:"damage,omitempty"`
}

func (a *Ammunition) EnchanceDamageImpact(damage []DamageImpact, pickDeviation func(deviation float32) float32) []DamageImpact {
	instantDamageEnchanced := false
	instantDamage := DamageImpact{}
	temporalDamage := []DamageImpact{}
	result := []DamageImpact{}
	for _, imp := range a.Damage {
		if imp.Duration == 0 {
			instantDamage.Accumulate(imp.Damage)
			instantDamage.Chance += imp.Chance
			instantDamage.Deviation += imp.Deviation
		} else {
			temporalDamage = append(temporalDamage, imp)
		}
	}
	if instantDamage.Deviation != 0 {
		instantDamage.EnchanceAll(pickDeviation(instantDamage.Deviation))
	}
	for _, imp := range damage {
		if imp.Duration == 0 {
			imp.Damage.Accumulate(instantDamage.Damage)
			imp.EnchanceChance(instantDamage.Chance)
			instantDamageEnchanced = true
		}
		result = append(result, imp)
	}
	result = append(result, temporalDamage...)
	if !instantDamageEnchanced && instantDamage.HasEffect() {
		return append([]DamageImpact{instantDamage}, result...)
	}
	return result
}
