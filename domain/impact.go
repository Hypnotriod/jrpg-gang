package domain

type Impact struct {
	Duration  int     `json:"duration,omitzero"`
	Chance    float32 `json:"chance,omitzero"`
	Deviation float32 `json:"deviation,omitzero"`
}

func (i *Impact) EnchanceChance(chance float32) {
	if i.Chance == 0 && chance < 0 {
		i.Chance = max(FULL_CHANCE+chance, MINIMUM_CHANCE)
	} else if i.Chance != 0 {
		i.Chance = max(i.Chance+chance, MINIMUM_CHANCE)
	}
}

type DamageImpact struct {
	Impact
	Damage
}

type UnitModificationImpact struct {
	Impact
	UnitModification
}
