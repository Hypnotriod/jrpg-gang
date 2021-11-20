package domain

type Potion struct {
	Item
	Damage      []DamageImpact          `json:"damage,omitempty"`
	Enhancement []UnitEnhancementImpact `json:"enhancement,omitempty"`
}
