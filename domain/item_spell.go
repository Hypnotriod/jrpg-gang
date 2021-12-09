package domain

type Magic struct {
	Item
	Requirements UnitAttributes          `json:"requirements"`
	Damage       []DamageImpact          `json:"damage,omitempty"`
	Enhancement  []UnitEnhancementImpact `json:"enhancement,omitempty"`
}
