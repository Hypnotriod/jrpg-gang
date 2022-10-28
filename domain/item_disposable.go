package domain

type Disposable struct {
	Item
	Quantity     uint                     `json:"quantity,omitempty"`
	Range        ActionRange              `json:"range"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}
