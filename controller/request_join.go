package controller

import (
	"jrpg-gang/auth"
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
)

type JoinRequestData struct {
	Token     auth.AuthenticationToken `json:"token,omitempty"`
	SessionId users.UserSessionId      `json:"sessionId,omitempty"`
}

func (c *GameController) handleJoinRequest(data *JoinRequestData, response *Response) (engine.PlayerId, []byte) {
	if data == nil {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	if data.SessionId != users.UserSessionIdEmpty {
		return c.handleRejoinRequest(response, data)
	}
	userModel, ok := c.persistance.GetUserFromAuthCache(data.Token)
	if userModel == nil || !ok {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotFound)
	}
	return c.handleRejoinWithAuthTokenRequest(response, data, userModel)
}

func (c *GameController) handleRejoinRequest(response *Response, data *JoinRequestData) (engine.PlayerId, []byte) {
	user, ok := c.users.GetAndRefreshBySessionId(data.SessionId)
	if !ok {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotFound)
	}
	user.IsOffline = false
	response.fillUserStatus(&user)
	return user.Id, response.WithStatus(ResponseStatusOk)
}

func (c *GameController) handleRejoinWithAuthTokenRequest(response *Response, data *JoinRequestData, userModel *model.UserModel) (engine.PlayerId, []byte) {
	if len(userModel.Units) == 0 {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotAllowed)
	}
	class := userModel.Class
	nickname := userModel.Nickname
	u, ok := userModel.Units[class]
	if !ok {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotFound)
	}
	unit := engine.NewGameUnit(u)
	c.itemsConfig.PopulateFromDescriptor(&unit.Inventory)
	userId := model.UserId(userModel.Id.Hex())
	user := users.NewUser(nickname, userModel.Email, userId, class, unit)
	c.users.AddUser(user)
	if jobStatus := c.persistance.GetJobStatus(user.UserId); jobStatus != nil {
		if jobStatus.IsInProgress || jobStatus.IsComplete {
			c.users.ChangeUserStatus(user.Id, users.UserStatusAtJob)
		}
		c.employment.SetStatus(user, *jobStatus)
	}
	c.persistance.RemoveUserFromAuthCache(data.Token)
	response.fillUserStatus(user)
	return user.Id, response.WithStatus(ResponseStatusOk)
}
