package domain

type Impact struct {
	Duration  int     `json:"duration,omitempty"`
	Chance    float32 `json:"chance,omitempty"`
	Deviation float32 `json:"deviation,omitempty"` // adds random 0 to n to all parameters
}

func (i *Impact) EnchanceChance(chance float32) {
	if i.Chance != 0 {
		i.Chance += chance
		if i.Chance < MINIMUM_CHANCE {
			i.Chance = MINIMUM_CHANCE
		}
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
