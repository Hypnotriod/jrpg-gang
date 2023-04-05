package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

func (c *GameController) handleCompleteJobRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	jobStatus, reward, ok := c.employment.CollectReward(&user)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if jobStatus != nil {
		c.persistJobStatus(user.Email, *jobStatus)
	}
	c.users.ChangeUserStatus(playerId, users.UserStatusJoined)
	c.users.AccumulateBooty(playerId, reward)
	response.Data[DataKeyReward] = reward
	return response.WithStatus(ResponseStatusOk)
}
