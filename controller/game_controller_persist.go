package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
	"maps"
	"time"
)

func (c *GameController) persistUser(user *users.User) bool {
	if user.IsGuest {
		return true
	}
	unit := user.Unit.ToPersist()
	userModel := &model.UserModel{
		Email:    user.Email,
		Nickname: user.Nickname,
		Class:    user.Class,
		Unit:     &unit.Unit,
	}
	return c.persistance.UpdateUser(userModel)
}

func (c *GameController) persistJobStatus(user *users.User, jobStatus engine.PlayerJobStatus) bool {
	if user.IsGuest {
		return true
	}
	jobStatusModel := &model.JobStatusModel{
		UserId:         user.UserId,
		IsInProgress:   jobStatus.IsInProgress,
		IsComplete:     jobStatus.IsComplete,
		CompletionTime: jobStatus.CompletionTime,
		Code:           jobStatus.Code,
	}
	jobStatusModel.Countdown = map[engine.PlayerJobCode]time.Time{}
	maps.Copy(jobStatusModel.Countdown, jobStatus.Countdown)
	return c.persistance.UpdateJobStatus(jobStatusModel)
}
