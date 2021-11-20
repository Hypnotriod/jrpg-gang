package domain

type Weapon struct {
	Equipment
	Condition   float32        `json:"condition"`
	Damage      []DamageImpact `json:"damage"`
	Enhancement []DamageImpact `json:"enhancement"`
}
