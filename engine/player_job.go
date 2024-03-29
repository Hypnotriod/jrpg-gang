package engine

import (
	"jrpg-gang/domain"
	"time"
)

type PlayerJobCode string

const (
	JobCodeEmpty PlayerJobCode = ""
)

type PlayerJob struct {
	Name         string                   `json:"name"`
	Code         PlayerJobCode            `json:"code"`
	Reward       domain.UnitBooty         `json:"reward"`
	Duration     float32                  `json:"duration"`
	Countdown    float32                  `json:"countdown"`
	Requirements *domain.UnitRequirements `json:"requirements,omitempty"`
	Description  string                   `json:"description,omitempty"`
}

type PlayerJobStatus struct {
	IsInProgress   bool                        `json:"isInProgress,omitempty"`
	IsComplete     bool                        `json:"isComplete,omitempty"`
	CompletionTime time.Time                   `json:"completionTime,omitempty"`
	Code           PlayerJobCode               `json:"code,omitempty"`
	Countdown      map[PlayerJobCode]time.Time `json:"countdown"`
}

func NewPlayerJobStatus() *PlayerJobStatus {
	status := &PlayerJobStatus{}
	status.Countdown = make(map[PlayerJobCode]time.Time)
	return status
}

func (s *PlayerJobStatus) Update() {
	timeNow := time.Now()
	for k, u := range s.Countdown {
		if timeNow.Compare(u) >= 0 {
			delete(s.Countdown, k)
		}
	}
	if s.IsInProgress && timeNow.Compare(s.CompletionTime) >= 0 {
		s.IsInProgress = false
		s.IsComplete = true
		s.CompletionTime = time.Time{}
	}
}

func (s *PlayerJobStatus) Clone() *PlayerJobStatus {
	status := &PlayerJobStatus{
		IsInProgress:   s.IsInProgress,
		IsComplete:     s.IsComplete,
		Code:           s.Code,
		CompletionTime: s.CompletionTime,
		Countdown:      make(map[PlayerJobCode]time.Time),
	}
	for k, v := range s.Countdown {
		status.Countdown[k] = v
	}
	return status
}

func (s *PlayerJobStatus) Apply(config PlayerJob) {
	timeNow := time.Now()
	s.IsInProgress = true
	s.IsComplete = false
	s.CompletionTime = timeNow.Add(time.Duration(config.Duration) * time.Second)
	s.Code = config.Code
	s.Countdown[config.Code] = s.CompletionTime.Add(time.Duration(config.Countdown) * time.Second)
}

func (s *PlayerJobStatus) Reset() {
	s.IsInProgress = false
	s.IsComplete = false
	s.CompletionTime = time.Time{}
	s.Code = JobCodeEmpty
}
