package employment

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
	"jrpg-gang/util"
	"sync"
	"time"
)

type EmploymentStatus struct {
	CurrentJob    *engine.PlayerJob  `json:"currentJob,omitempty"`
	IsInProgress  bool               `json:"isInProgress,omitempty"`
	IsComplete    bool               `json:"isComplete,omitempty"`
	TimeLeft      float32            `json:"timeLeft,omitempty"`
	AvailableJobs []engine.PlayerJob `json:"availableJobs"`
}

type Employment struct {
	mu           sync.RWMutex
	jobs         []engine.PlayerJob
	jobCodeToJob map[engine.PlayerJobCode]*engine.PlayerJob
	jobsStatus   map[model.UserId]*engine.PlayerJobStatus
}

func NewEmployment() *Employment {
	e := &Employment{}
	e.jobCodeToJob = make(map[engine.PlayerJobCode]*engine.PlayerJob)
	e.jobsStatus = make(map[model.UserId]*engine.PlayerJobStatus)
	return e
}

func (e *Employment) Load(path string) error {
	jobs, err := util.ReadJsonFile(&[]engine.PlayerJob{}, path)
	if err != nil {
		return err
	}
	e.mu.Lock()
	defer e.mu.Unlock()
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

func (e *Employment) retrieveUserJobStatus(userId model.UserId) *engine.PlayerJobStatus {
	status, ok := e.jobsStatus[userId]
	if !ok {
		status = engine.NewPlayerJobStatus()
		e.jobsStatus[userId] = status
	}
	status.Update()
	return status
}

func (e *Employment) SetStatus(user *users.User, jobStatusModel model.JobStatusModel) {
	status := &engine.PlayerJobStatus{
		IsInProgress:   jobStatusModel.IsInProgress,
		IsComplete:     jobStatusModel.IsComplete,
		CompletionTime: jobStatusModel.CompletionTime,
		Code:           jobStatusModel.Code,
		Countdown:      jobStatusModel.Countdown,
	}
	e.mu.Lock()
	defer e.mu.Unlock()
	e.jobsStatus[user.UserId] = status
}

func (e *Employment) ClearStatus(user *users.User) {
	e.mu.Lock()
	defer e.mu.Unlock()
	delete(e.jobsStatus, user.UserId)
}

func (e *Employment) GetStatus(user *users.User) EmploymentStatus {
	timeNow := time.Now()
	e.mu.Lock()
	status := e.retrieveUserJobStatus(user.UserId)
	result := EmploymentStatus{
		IsInProgress:  status.IsInProgress,
		IsComplete:    status.IsComplete,
		AvailableJobs: []engine.PlayerJob{},
	}
	if job, ok := e.jobCodeToJob[status.Code]; ok {
		result.CurrentJob = job
	}
	e.mu.Unlock()
	if status.IsInProgress {
		result.TimeLeft = float32(status.CompletionTime.Sub(timeNow).Seconds())
	}
	e.mu.RLock()
	defer e.mu.RUnlock()
	for _, job := range e.jobs {
		if _, ok := status.Countdown[job.Code]; !ok {
			result.AvailableJobs = append(result.AvailableJobs, job)
		}
	}
	return result
}

func (e *Employment) CollectReward(user *users.User) (*engine.PlayerJobStatus, domain.UnitBooty, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	status := e.retrieveUserJobStatus(user.UserId)
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
	e.mu.Lock()
	defer e.mu.Unlock()
	status := e.retrieveUserJobStatus(user.UserId)
	if !status.IsInProgress {
		return nil, false
	}
	status.Reset()
	return status.Clone(), true
}

func (e *Employment) ApplyForAJob(user *users.User, code engine.PlayerJobCode) (*engine.PlayerJobStatus, bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	status := e.retrieveUserJobStatus(user.UserId)
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
