package domain

type Weapon struct {
	Equipment
	AmmunitionKind AmmunitionKind     `json:"ammunitionKind,omitempty"`
	Range          ActionRange        `json:"range"`
	UseCost        UnitBaseAttributes `json:"useCost"`
	Damage         []DamageImpact     `json:"damage"`
}

func (w Weapon) RequiresAmmunition() bool {
	return w.AmmunitionKind != NoAmmunition
}
