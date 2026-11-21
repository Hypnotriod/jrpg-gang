package domain

import "maps"

type UnitAI struct {
	WeaponPreference map[ItemCode]uint `json:"weaponPreference,omitempty"`
}

func (a *UnitAI) Clone() *UnitAI {
	if a == nil {
		return nil
	}
	return &UnitAI{
		WeaponPreference: maps.Clone(a.WeaponPreference),
	}
}
