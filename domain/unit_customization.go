package domain

type UnitSkinCode string

type UnitCustomization struct {
	Skin UnitSkinCode `json:"skin,omitempty"`
}

func (c *UnitCustomization) Clone() *UnitCustomization {
	if c == nil {
		return nil
	}
	return &UnitCustomization{
		Skin: c.Skin,
	}
}
