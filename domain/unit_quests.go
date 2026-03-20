package domain

type UnitQuestStatus string

const (
	UnitQuestStatusEmpty     UnitQuestStatus = ""
	UnitQuestStatusInactive  UnitQuestStatus = "inactive"
	UnitQuestStatusActive    UnitQuestStatus = "active"
	UnitQuestStatusCompleted UnitQuestStatus = "completed"
	UnitQuestStatusFailed    UnitQuestStatus = "failed"
)

type UnitQuests map[QuestCode]UnitQuestStatus

func (a UnitQuests) Deactivate(questId QuestCode) {
	delete(a, questId)
}

func (a UnitQuests) Activate(questId QuestCode) {
	a[questId] = UnitQuestStatusActive
}

func (a UnitQuests) Complete(questId QuestCode) {
	a[questId] = UnitQuestStatusCompleted
}

func (a UnitQuests) Set(quests UnitQuests) {
	for questId, value := range quests {
		a[questId] = value
	}
}

func (a UnitQuests) Test(quests UnitQuests) bool {
	for questId, value := range quests {
		if a[questId] == UnitQuestStatusEmpty && value != UnitQuestStatusInactive {
			return false
		}
		if a[questId] != value {
			return false
		}
	}
	return true
}
