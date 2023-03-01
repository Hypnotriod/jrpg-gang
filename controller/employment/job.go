package employment

import (
	"jrpg-gang/domain"
	"time"
)

type JobCode string

const (
	JobCodeEmpty JobCode = ""
)

type Job struct {
	Code        JobCode          `json:"code"`
	Reward      domain.UnitBooty `json:"reward"`
	Duration    float32          `json:"duration"`
	Countdown   float32          `json:"countdown"`
	Description string           `json:"description,omitempty"`
}

type JobStatus struct {
	IsInProgress   bool                  `json:"isInProgress,omitempty"`
	IsComplete     bool                  `json:"isComplete,omitempty"`
	CompletionTime time.Time             `json:"completionTime,omitempty"`
	Code           JobCode               `json:"code,omitempty"`
	Countdown      map[JobCode]time.Time `json:"countdown"`
}

func NewJobStatus() *JobStatus {
	status := &JobStatus{}
	status.Countdown = make(map[JobCode]time.Time)
	return status
}

func (s *JobStatus) Update() {
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
