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
}

func NewGameQuests(quests *[]domain.Quest) *GameQuests {
	r := &GameQuests{
		quests: quests,
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
		r.Quests[i].Quest = *quest
		r.Quests[i].Status = unit.Quests[quest.Code]
	}
	return r
}

func (q *GameQuests) ExecuteAction(action domain.Action, unit *domain.Unit, code domain.QuestCode) *domain.ActionResult {
	switch action.Action {
	case domain.ActionActivate:
		return q.activate(unit, code)
	case domain.ActionDeactivate:
		return q.deactivate(unit, code)
	case domain.ActionComplete:
		return q.complete(unit, code)
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
		unit.Achievements.Set(*quest.Activation.Achievements)
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

func (q *GameQuests) complete(unit *domain.Unit, code domain.QuestCode) *domain.ActionResult {
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
		unit.Achievements.Set(*quest.Completion.Achievements)
	}
	if !quest.Reward.UnitBooty.IsEmpty() {
		unit.Booty.Accumulate(quest.Reward.UnitBooty)
		result.Booty = &domain.UnitBooty{}
		result.Booty.Accumulate(quest.Reward.UnitBooty)
	}
	unit.Achievements.Accumulate(*quest.Reward.Achievements)
	result.Achievements.Accumulate(*quest.Reward.Achievements)
	result.Items = util.ClonePtr(quest.Reward.Items)
	unit.Quests[code] = domain.UnitQuestStatusDone
	return result.WithResult(domain.ResultAccomplished)
}
