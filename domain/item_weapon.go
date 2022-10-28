package domain

type Weapon struct {
	Equipment
	AmmunitionKind string             `json:"ammunitionKind,omitempty"`
	Range          ActionRange        `json:"range"`
	UseCost        UnitBaseAttributes `json:"useCost"`
	Damage         []DamageImpact     `json:"damage"`
}

func (w Weapon) RequiresAmmunition() bool {
	return len(w.AmmunitionKind) != 0
}
