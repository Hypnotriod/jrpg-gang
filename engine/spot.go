package engine

import (
	"jrpg-gang/domain"
)

type Spot struct {
	Name         string                   `json:"name"`
	Code         string                   `json:"code"`
	Battlefield  Battlefield              `json:"battlefield"`
	Booty        []domain.UnitBooty       `json:"booty"`
	Experience   uint                     `json:"experience"`
	Achievements []domain.UnitAchievement `json:"achievements"`
}

func (s *Spot) Clone() *Spot {
	r := &Spot{}
	r.Name = s.Name
	r.Code = s.Code
	r.Battlefield = *s.Battlefield.Clone()
	r.Experience = s.Experience
	r.Booty = s.Booty
	r.Achievements = s.Achievements
	return r
}
