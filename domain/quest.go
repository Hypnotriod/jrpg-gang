package domain

import "jrpg-gang/util"

type QuestCode string

type QuestTrigger struct {
	Requirements *UnitRequirements `json:"requirements,omitempty"`
	Achievements *UnitAchievements `json:"achievements,omitempty"`
}

type QuestReward struct {
	UnitBooty
	Achievements *[]UnitAchievement `json:"achievements,omitempty"`
	Items        *UnitInventory     `json:"items,omitempty"`
}

func (r *QuestReward) Clone() *QuestReward {
	res := &QuestReward{}
	res.UnitBooty = r.UnitBooty
	if r.Achievements != nil {
		res.Achievements = util.ClonePtr(r.Achievements)
	}
	if r.Items != nil {
		res.Items = r.Items.Clone()
	}
	return res
}

type Quest struct {
	Name        string       `json:"name"`
	Code        QuestCode    `json:"code"`
	Reward      QuestReward  `json:"reward"`
	Activation  QuestTrigger `json:"activation"`
	Completion  QuestTrigger `json:"completion"`
	Description string       `json:"description,omitempty"`
}

func (q *Quest) Clone() *Quest {
	res := &Quest{}
	res.Name = q.Name
	res.Code = q.Code
	res.Reward = *q.Reward.Clone()
	res.Activation = q.Activation
	res.Completion = q.Completion
	res.Description = q.Description
	return res
}
