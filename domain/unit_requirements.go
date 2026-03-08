package domain

type UnitRequirements struct {
	UnitAttributes
	Level        uint             `json:"level"`
	Class        UnitClass        `json:"class,omitempty"`
	Achievements UnitAchievements `json:"achievements,omitempty"`
	Quests       UnitQuests       `json:"quests,omitempty"`
}

func (r *UnitRequirements) Check(unit *Unit, attributes UnitAttributes) bool {
	return (r.Class == UnitClassEmpty || r.Class == unit.Class) &&
		r.Level <= unit.Stats.Progress.Level &&
		unit.Achievements.Test(r.Achievements) &&
		unit.Quests.Test(r.Quests) &&
		r.Strength <= attributes.Strength &&
		r.Physique <= attributes.Physique &&
		r.Agility <= attributes.Agility &&
		r.Endurance <= attributes.Endurance &&
		r.Intelligence <= attributes.Intelligence &&
		r.Initiative <= attributes.Initiative &&
		r.Luck <= attributes.Luck
}
