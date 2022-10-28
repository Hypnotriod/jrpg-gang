package domain

type Magic struct {
	Item
	Requirements UnitAttributes           `json:"requirements"`
	Range        ActionRange              `json:"range"`
	UseCost      UnitBaseAttributes       `json:"useCost"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}
