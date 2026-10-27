package domain

type Weapon struct {
	Equipment
	AmmunitionKind AmmunitionKind     `json:"ammunitionKind,omitempty"`
	Range          ActionRange        `json:"range"`
	Spread         []Position         `json:"spread,omitempty"`
	UseCost        UnitBaseAttributes `json:"useCost"`
	Damage         []DamageImpact     `json:"damage"`
}

func (w *Weapon) TotalDamage(pickDeviation func(deviation float32) float32) []DamageImpact {
	var extraDamage []DamageImpact
	for i := range w.Upgrades {
		extraDamage = append(extraDamage, w.Upgrades[i].Damage...)
	}
	if len(extraDamage) == 0 {
		return w.Damage
	}
	instantDamageEnchanced := false
	instantDamage := DamageImpact{}
	temporalDamage := []DamageImpact{}
	result := []DamageImpact{}
	for _, imp := range extraDamage {
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
		instantDamage.Deviation = 0
	}
	for _, imp := range w.Damage {
		if imp.Duration == 0 {
			if imp.Deviation != 0 {
				imp.EnchanceAll(pickDeviation(imp.Deviation))
				imp.Deviation = 0
			}
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

func (w *Weapon) TotalUseCost() UnitBaseAttributes {
	useCost := w.UseCost
	for i := range w.Upgrades {
		if w.Upgrades[i].UseCost != nil {
			useCost.Accumulate(*w.Upgrades[i].UseCost)
		}
	}
	return useCost
}

func (w Weapon) RequiresAmmunition() bool {
	return w.AmmunitionKind != NoAmmunition
}
