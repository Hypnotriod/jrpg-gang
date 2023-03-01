package employment

import (
	"jrpg-gang/controller/users"
	"sync"
	"time"
)

type EmploymentConfig struct {
	Jobs []Job `json:"jobs"`
}

type Employment struct {
	mu           sync.RWMutex
	config       EmploymentConfig
	jobCodeToJob map[JobCode]*Job
	jobsStatus   map[string]*JobStatus
}

func NewEmployment(config EmploymentConfig) *Employment {
	e := &Employment{}
	e.config = config
	e.jobCodeToJob = make(map[JobCode]*Job)
	e.jobsStatus = make(map[string]*JobStatus)
	e.prepare()
	return e
}

func (e *Employment) prepare() {
	for i := range e.config.Jobs {
		job := &e.config.Jobs[i]
		e.jobCodeToJob[job.Code] = job
	}
}

func (e *Employment) ApplyForAJob(user *users.User, code JobCode) bool {
	defer e.mu.Unlock()
	e.mu.Lock()
	status, ok := e.jobsStatus[user.Email]
	if !ok {
		status = NewJobStatus()
		e.jobsStatus[user.Email] = status
		return false
	}
	status.Update()
	if status.IsInProgress || status.IsComplete {
		return false
	}
	config, ok := e.jobCodeToJob[code]
	if !ok {
		return false
	}
	if _, ok := status.Countdown[config.Code]; ok {
		return false
	}
	timeNow := time.Now()
	status.CompletionTime = timeNow.Add(time.Duration(config.Duration) * time.Second)
	status.Countdown[config.Code] = status.CompletionTime.Add(time.Duration(config.Countdown) * time.Second)
	return true
}
