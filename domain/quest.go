package domain

type QuestCode string

type QuestTrigger struct {
	Requirements *UnitRequirements `json:"requirements,omitempty"`
	Achievements *UnitAchievements `json:"achievements,omitempty"`
}

type QuestReward struct {
	UnitBooty
	Achievements *[]UnitAchievement         `json:"achievements,omitempty"`
	Items        *[]UnitInventoryDescriptor `json:"items,omitempty"`
}

type Quest struct {
	Name        string       `json:"name"`
	Code        QuestCode    `json:"code"`
	Reward      QuestReward  `json:"reward"`
	Activation  QuestTrigger `json:"activation"`
	Completion  QuestTrigger `json:"completion"`
	Description string       `json:"description,omitempty"`
}
