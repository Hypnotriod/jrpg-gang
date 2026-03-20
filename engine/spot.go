package engine

import (
	"jrpg-gang/domain"
)

type Spot struct {
	Name          string                `json:"name"`
	Code          string                `json:"code"`
	Battlefield   Battlefield           `json:"battlefield"`
	Booty         []domain.UnitBooty    `json:"booty"`
	Experience    uint                  `json:"experience"`
	QuestTriggers []domain.QuestTrigger `json:"questTriggers,omitempty"`
}

func (s *Spot) Clone() *Spot {
	r := &Spot{}
	r.Name = s.Name
	r.Code = s.Code
	r.Battlefield = *s.Battlefield.Clone()
	r.Experience = s.Experience
	r.Booty = s.Booty
	r.QuestTriggers = s.QuestTriggers
	return r
}
