package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameQuestProgress struct {
	Target uint `json:"target"`
	Goal   uint `json:"goal"`
}

type GameQuestStatus struct {
	domain.Quest
	Status   domain.UnitQuestStatus                       `json:"status"`
	Progress map[domain.UnitAchievement]GameQuestProgress `json:"progress,omitempty"`
}

type GameQuestsStatus struct {
	Quests []GameQuestStatus `json:"quests"`
}

type GameQuests struct {
	quests *[]domain.Quest
	rndGen *util.RndGen
}

func NewGameQuests(quests *[]domain.Quest, populateFromDescriptor func(inventory *domain.UnitInventory)) *GameQuests {
	r := &GameQuests{}
	r.quests = quests
	r.rndGen = util.NewRndGen()
	for i := range *quests {
		items := (*quests)[i].Reward.Items
		if items != nil {
			populateFromDescriptor(items)
			items.PopulateUids(r.rndGen)
			items.UnequipAmmunition()
		}
	}
	return r
}

func (q *GameQuests) GetStatus(unit *domain.Unit) *GameQuestsStatus {
	r := &GameQuestsStatus{}
	quests := *q.quests
	r.Quests = make([]GameQuestStatus, 0, len(quests))
	for i := range quests {
		quest := &quests[i]
		if quest.Activation.Requirements != nil && !unit.Quests.Test(quest.Activation.Requirements.Quests) {
			continue
		}
		status := unit.Quests[quest.Code]
		if status == domain.UnitQuestStatusEmpty {
			status = domain.UnitQuestStatusInactive
		}
		r.Quests[i].Quest = *quest.Clone()
		r.Quests[i].Status = status
		if status != domain.UnitQuestStatusActive || quest.Completion.Requirements == nil {
			continue
		}
		r.Quests[i].Progress = map[domain.UnitAchievement]GameQuestProgress{}
		for achievement, goal := range quest.Completion.Requirements.Achievements {
			target := unit.Achievements[achievement]
			r.Quests[i].Progress[achievement] = GameQuestProgress{
				Target: target,
				Goal:   goal,
			}
		}
	}
	return r
}

func (q *GameQuests) ExecuteAction(action domain.Action, unit *domain.Unit, rndGen *util.RndGen) *domain.ActionResult {
	switch action.Action {
	case domain.ActionActivate:
		return q.activate(unit, action.QuestCode)
	case domain.ActionDeactivate:
		return q.deactivate(unit, action.QuestCode)
	case domain.ActionComplete:
		return q.complete(unit, action.QuestCode, rndGen)
	}
	return domain.NewActionResult().WithResult(domain.ResultNotAccomplished)
}

func (q *GameQuests) activate(unit *domain.Unit, code domain.QuestCode) *domain.ActionResult {
	result := domain.NewActionResult()
	quest := util.Find(*q.quests, func(quest domain.Quest) bool {
		return quest.Code == code
	})
	if quest == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	if quest.Activation.Requirements != nil && !unit.CheckRequirements(*quest.Activation.Requirements) {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if quest.Activation.Achievements != nil {
		unit.Achievements.Set(quest.Activation.Achievements)
	}
	unit.Quests[code] = domain.UnitQuestStatusActive
	return result.WithResult(domain.ResultAccomplished)
}

func (q *GameQuests) deactivate(unit *domain.Unit, code domain.QuestCode) *domain.ActionResult {
	result := domain.NewActionResult()
	quest := util.Find(*q.quests, func(quest domain.Quest) bool {
		return quest.Code == code
	})
	if quest == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	unit.Quests[code] = domain.UnitQuestStatusInactive
	return result.WithResult(domain.ResultAccomplished)
}

func (q *GameQuests) complete(unit *domain.Unit, code domain.QuestCode, rndGen *util.RndGen) *domain.ActionResult {
	result := domain.NewActionResult()
	quest := util.Find(*q.quests, func(quest domain.Quest) bool {
		return quest.Code == code
	})
	if quest == nil {
		return result.WithResult(domain.ResultNotFound)
	}
	if quest.Completion.Requirements != nil && !unit.CheckRequirements(*quest.Completion.Requirements) {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if quest.Completion.Achievements != nil {
		unit.Achievements.Set(quest.Completion.Achievements)
	}
	if !quest.Reward.UnitBooty.IsEmpty() {
		unit.Booty.Accumulate(quest.Reward.UnitBooty)
		result.Booty = &domain.UnitBooty{}
		result.Booty.Accumulate(quest.Reward.UnitBooty)
	}
	if quest.Reward.Items != nil {
		unit.Inventory.Merge(quest.Reward.Items, rndGen)
		result.Items = quest.Reward.Items.Clone()
	}
	if quest.Reward.Experience != 0 {
		unit.Stats.Progress.Experience += quest.Reward.Experience
		result.Experience[unit.Uid] = quest.Reward.Experience
	}
	unit.Achievements.Merge(quest.Reward.Achievements)
	result.Achievements.Merge(quest.Reward.Achievements)
	unit.Quests[code] = domain.UnitQuestStatusCompleted
	return result.WithResult(domain.ResultAccomplished)
}
