package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
	"time"
)

func (c *GameController) persistUser(user *users.User) bool {
	unit := user.Unit.ToPersist()
	userModel := model.UserModel{
		Email:    user.Email,
		Nickname: user.Nickname,
		Class:    user.Class,
		Unit:     &unit.Unit,
	}
	return c.persistance.UpdateUser(userModel)
}

func (c *GameController) persistJobStatus(userId model.UserId, jobStatus engine.PlayerJobStatus) bool {
	jobStatusModel := model.JobStatusModel{
		UserId:         userId,
		IsInProgress:   jobStatus.IsInProgress,
		IsComplete:     jobStatus.IsComplete,
		CompletionTime: jobStatus.CompletionTime,
		Code:           jobStatus.Code,
	}
	jobStatusModel.Countdown = make(map[engine.PlayerJobCode]time.Time)
	for k, v := range jobStatus.Countdown {
		jobStatusModel.Countdown[k] = v
	}
	return c.persistance.UpdateJobStatus(jobStatusModel)
}
