package domain

type UnitQuestStatus uint

const (
	UnitQuestStatusInactive UnitQuestStatus = iota
	UnitQuestStatusActive
	UnitQuestStatusDone
)

func (s UnitQuestStatus) String() string {
	switch s {
	case UnitQuestStatusActive:
		return "active"
	case UnitQuestStatusDone:
		return "done"
	default:
		return "inactive"
	}
}

type UnitQuests map[QuestCode]UnitQuestStatus

func (a UnitQuests) Deactivate(questId QuestCode) {
	delete(a, questId)
}

func (a UnitQuests) Activate(questId QuestCode) {
	a[questId] = UnitQuestStatusActive
}

func (a UnitQuests) Complete(questId QuestCode) {
	a[questId] = UnitQuestStatusDone
}

func (a UnitQuests) Merge(quests UnitQuests) {
	for questId, value := range quests {
		a[questId] = max(a[questId], value)
	}
}

func (a UnitQuests) Test(quests UnitQuests) bool {
	for questId, value := range quests {
		if a[questId] != value {
			return false
		}
	}
	return true
}
