package domain

type UnitRequirements struct {
	UnitAttributes
	Class UnitClass `json:"class,omitempty"`
}

func (r *UnitRequirements) Check(class UnitClass, attributes UnitAttributes) bool {
	return (r.Class == UnitClassEmpty || r.Class == class) &&
		r.Strength <= attributes.Strength &&
		r.Physique <= attributes.Physique &&
		r.Agility <= attributes.Agility &&
		r.Endurance <= attributes.Endurance &&
		r.Intelligence <= attributes.Intelligence &&
		r.Initiative <= attributes.Initiative &&
		r.Luck <= attributes.Luck
}
