package domain

type Disposable struct {
	Item
	Quantity     uint                     `json:"quantity,omitzero"`
	Range        ActionRange              `json:"range"`
	Spread       []Position               `json:"spread,omitempty"`
	UseCost      UnitBaseAttributes       `json:"useCost"`
	Damage       []DamageImpact           `json:"damage,omitempty"`
	Modification []UnitModificationImpact `json:"modification,omitempty"`
}
