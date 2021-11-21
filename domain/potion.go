package domain

type Potion struct {
	Item
	Impact      []DamageImpact          `json:"impact,omitempty"`
	Enhancement []UnitEnhancementImpact `json:"enhancement,omitempty"`
}
