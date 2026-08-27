package domain

type Magic struct {
	Item
	Requirements *UnitRequirements        `json:"requirements,omitempty"`
	Range        ActionRange              `json:"range"`
	Spread       []Position               `json:"spread,omitempty"`
	UseCost      UnitBaseAttributes       `json:"useCost"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}
