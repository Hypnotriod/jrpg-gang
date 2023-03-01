package employment

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
	"time"
)

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

func (e *Employment) ApplyForAJob(user *users.User, code engine.PlayerJobCode) (*engine.PlayerJobStatus, bool) {
	defer e.mu.Unlock()
	e.mu.Lock()
	status, ok := e.jobsStatus[user.Email]
	if !ok {
		status = engine.NewPlayerJobStatus()
		e.jobsStatus[user.Email] = status
	}
	status.Update()
	if status.IsInProgress || status.IsComplete {
		return status.Clone(), false
	}
	config, ok := e.jobCodeToJob[code]
	if !ok {
		return status.Clone(), false
	}
	if _, ok := status.Countdown[config.Code]; ok {
		return status.Clone(), false
	}
	timeNow := time.Now()
	status.CompletionTime = timeNow.Add(time.Duration(config.Duration) * time.Second)
	status.Countdown[config.Code] = status.CompletionTime.Add(time.Duration(config.Countdown) * time.Second)
	return status.Clone(), true
}