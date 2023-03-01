package employment

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
	"time"
)

type EmploymentStatus struct {
	CurrentJobCode engine.PlayerJobCode `json:"currentJobCode,omitempty"`
	IsInProgress   bool                 `json:"isInProgress,omitempty"`
	IsComplete     bool                 `json:"isComplete,omitempty"`
	TimeLeft       float32              `json:"timeLeft,omitempty"`
	AvailableJobs  []engine.PlayerJob   `json:"availableJobs"`
}

type Employment struct {
	mu           sync.RWMutex
	jobs         []engine.PlayerJob
	jobCodeToJob map[engine.PlayerJobCode]*engine.PlayerJob
	jobsStatus   map[string]*engine.PlayerJobStatus
}

func NewEmployment() *Employment {
	e := &Employment{}
	e.jobCodeToJob = make(map[engine.PlayerJobCode]*engine.PlayerJob)
	e.jobsStatus = make(map[string]*engine.PlayerJobStatus)
	return e
}

func (e *Employment) Load(path string) error {
	jobs, err := util.ReadJsonFile(&[]engine.PlayerJob{}, path)
	if err != nil {
		return err
	}
	defer e.mu.Unlock()
	e.mu.Lock()
	e.prepare(jobs)
	return nil
}

func (e *Employment) prepare(jobs *[]engine.PlayerJob) {
	e.jobs = *jobs
	for i := range e.jobs {
		job := &e.jobs[i]
		e.jobCodeToJob[job.Code] = job
	}
}

func (e *Employment) retrieveUserJobStatus(email string) *engine.PlayerJobStatus {
	status, ok := e.jobsStatus[email]
	if !ok {
		status = engine.NewPlayerJobStatus()
		e.jobsStatus[email] = status
	}
	status.Update()
	return status
}

func (e *Employment) GetStatus(user *users.User) EmploymentStatus {
	timeNow := time.Now()
	e.mu.Lock()
	status := e.retrieveUserJobStatus(user.Email)
	e.mu.Unlock()
	result := EmploymentStatus{
		IsInProgress:   status.IsInProgress,
		IsComplete:     status.IsComplete,
		CurrentJobCode: status.Code,
	}
	if status.IsInProgress {
		result.TimeLeft = float32(status.CompletionTime.Sub(timeNow).Seconds())
	}
	defer e.mu.RUnlock()
	e.mu.RLock()
	for _, job := range e.jobs {
		if _, ok := status.Countdown[job.Code]; !ok {
			result.AvailableJobs = append(result.AvailableJobs, job)
		}
	}
	return result
}

func (e *Employment) CollectReward(user *users.User) (*engine.PlayerJobStatus, domain.UnitBooty, bool) {
	defer e.mu.Unlock()
	e.mu.Lock()
	status := e.retrieveUserJobStatus(user.Email)
	if status.IsInProgress || !status.IsComplete {
		return nil, domain.UnitBooty{}, false
	}
	config, ok := e.jobCodeToJob[status.Code]
	status.Reset()
	if !ok {
		return status.Clone(), domain.UnitBooty{}, false
	}
	return status.Clone(), config.Reward, true
}

func (e *Employment) QuitJob(user *users.User) (*engine.PlayerJobStatus, bool) {
	defer e.mu.Unlock()
	e.mu.Lock()
	status := e.retrieveUserJobStatus(user.Email)
	if !status.IsInProgress {
		return nil, false
	}
	status.Reset()
	return status.Clone(), true
}

func (e *Employment) ApplyForAJob(user *users.User, code engine.PlayerJobCode) (*engine.PlayerJobStatus, bool) {
	defer e.mu.Unlock()
	e.mu.Lock()
	status := e.retrieveUserJobStatus(user.Email)
	if status.IsInProgress || status.IsComplete {
		return nil, false
	}
	config, ok := e.jobCodeToJob[code]
	if !ok {
		return nil, false
	}
	if config.Requirements != nil && !user.Unit.CheckRequirements(*config.Requirements) {
		return nil, false
	}
	if _, ok := status.Countdown[config.Code]; ok {
		return nil, false
	}
	status.Apply(*config)
	return status.Clone(), true
}
