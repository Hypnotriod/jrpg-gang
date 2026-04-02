package engine

import (
	"jrpg-gang/domain"
	"jrpg-gang/util"
)

type GameQuestStatus struct {
	domain.Quest
	Status domain.UnitQuestStatus `json:"status"`
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
		questStatus := GameQuestStatus{Quest: *quest.Clone(), Status: status}
		r.Quests = append(r.Quests, questStatus)
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
	if unit.Quests[code] != domain.UnitQuestStatusEmpty && unit.Quests[code] != domain.UnitQuestStatusInactive {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if quest.Activation.Requirements != nil && !unit.CheckRequirements(*quest.Activation.Requirements) {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if quest.Activation.Quests != nil {
		unit.Quests.Set(quest.Activation.Quests)
		result.Quests[unit.Uid] = quest.Activation.Quests
	}
	if quest.Activation.Achievements != nil {
		if _, ok := result.Achievements[unit.Uid]; !ok {
			result.Achievements[unit.Uid] = domain.UnitAchievements{}
		}
		unit.Achievements.Set(quest.Activation.Achievements)
		result.Achievements[unit.Uid].Set(quest.Activation.Achievements)
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
	if unit.Quests[code] != domain.UnitQuestStatusActive {
		return result.WithResult(domain.ResultNotAllowed)
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
	if unit.Quests[code] != domain.UnitQuestStatusActive {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if quest.Completion.Requirements != nil && !unit.CheckRequirements(*quest.Completion.Requirements) {
		return result.WithResult(domain.ResultNotAllowed)
	}
	if quest.Completion.Quests != nil {
		unit.Quests.Set(quest.Completion.Quests)
		result.Quests[unit.Uid] = quest.Completion.Quests
	}
	if quest.Completion.Achievements != nil {
		if _, ok := result.Achievements[unit.Uid]; !ok {
			result.Achievements[unit.Uid] = domain.UnitAchievements{}
		}
		unit.Achievements.Set(quest.Completion.Achievements)
		result.Achievements[unit.Uid].Set(quest.Completion.Achievements)
	}
	if quest.Reward.Achievements != nil {
		if _, ok := result.Achievements[unit.Uid]; !ok {
			result.Achievements[unit.Uid] = domain.UnitAchievements{}
		}
		unit.Achievements.Merge(quest.Reward.Achievements)
		result.Achievements[unit.Uid].Merge(quest.Reward.Achievements)
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
	unit.Quests[code] = domain.UnitQuestStatusCompleted
	return result.WithResult(domain.ResultAccomplished)
}
