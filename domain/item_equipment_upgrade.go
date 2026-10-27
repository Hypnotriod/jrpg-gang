package domain

type EquipmentUpgrade struct {
	Item
	Modification *UnitModification   `json:"modification,omitempty"`
	Damage       []DamageImpact      `json:"damage,omitempty"`
	UseCost      *UnitBaseAttributes `json:"useCost,omitempty"`
}
