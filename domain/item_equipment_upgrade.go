package domain

type UpgradeKind string

const (
	UpgradeKindWeapon UpgradeKind = "weapon"
	UpgradeKindArmor  UpgradeKind = "armor"
)

type Upgrade struct {
	Item
	Kind         UpgradeKind         `json:"kind"`
	Modification *UnitModification   `json:"modification,omitempty"`
	Damage       []DamageImpact      `json:"damage,omitempty"`
	UseCost      *UnitBaseAttributes `json:"useCost,omitempty"`
	Requirements *UnitRequirements   `json:"requirements,omitempty"`
}
