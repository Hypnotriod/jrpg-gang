package domain

type AmmunitionKind string

const (
	NoAmmunition AmmunitionKind = ""
)

type Ammunition struct {
	Item
	Equipped bool           `json:"equipped,omitempty"`
	Kind     AmmunitionKind `json:"kind"`
	Quantity uint           `json:"quantity,omitempty"`
	Damage   []DamageImpact `json:"damage,omitempty"`
}

func (a *Ammunition) EnchanceDamageImpact(damage []DamageImpact) []DamageImpact {
	instantDamageEnchanced := false
	instantDamage := DamageImpact{}
	temporalDamage := []DamageImpact{}
	result := []DamageImpact{}
	for _, imp := range a.Damage {
		if imp.Duration == 0 {
			instantDamage.Accumulate(imp.Damage)
			instantDamage.Chance += imp.Chance
		} else {
			temporalDamage = append(temporalDamage, imp)
		}
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
