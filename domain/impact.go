package domain

type Impact struct {
	Duration int     `json:"duration,omitempty"`
	Chance   float32 `json:"chance,omitempty"`
}

type DamageImpact struct {
	Impact
	Damage
}

type UnitModificationImpact struct {
	Impact
	UnitModification
}
