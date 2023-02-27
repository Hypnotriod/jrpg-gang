package controller

import (
	"jrpg-gang/auth"
	"jrpg-gang/controller/users"
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
	"regexp"
)

type JoinRequestData struct {
	Token     auth.AuthenticationToken `json:"token,omitempty"`
	Nickname  string                   `json:"nickname,omitempty"`
	Class     domain.UnitClass         `json:"class,omitempty"`
	SessionId users.UserSessionId      `json:"sessionId,omitempty"`
}

func (c *GameController) handleJoinRequest(request *Request, response *Response) (engine.PlayerId, string) {
	data := parseRequestData(&JoinRequestData{}, request.Data)
	if data == nil {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	if data.SessionId != users.UserSessionIdEmpty {
		return c.handleRejoinRequest(request, response, data)
	}
	userModel, ok := c.persistance.GetUserFromCache(data.Token)
	if userModel == nil || !ok {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotFound)
	}
	return c.handleRejoinWithCredentialsRequest(request, response, data, userModel)
}

func (c *GameController) handleRejoinRequest(request *Request, response *Response, data *JoinRequestData) (engine.PlayerId, string) {
	user, ok := c.users.GetAndRefreshBySessionId(data.SessionId)
	if !ok {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusNotFound)
	}
	user.IsOffline = false
	response.fillUserStatus(&user)
	return user.Id, response.WithStatus(ResponseStatusOk)
}

func (c *GameController) handleRejoinWithCredentialsRequest(request *Request, response *Response, data *JoinRequestData, userModel *model.UserModel) (engine.PlayerId, string) {
	var unit *engine.GameUnit
	var nickname string
	var class domain.UnitClass
	if userModel.Unit != nil {
		unit = engine.NewGameUnit(userModel.Unit)
		c.itemsConfig.PopulateFromDescriptor(&unit.Inventory)
		nickname = userModel.Nickname
		class = userModel.Class
	} else {
		if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, data.Nickname); !matched {
			return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
		}
		if ok := c.persistance.HasUserWithNickname(data.Nickname); ok {
			return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusAlreadyExists)
		}
		unit = c.unitsConfig.GetByCode(domain.UnitCode(data.Class)) // todo: allow only specific unit codes
		nickname = data.Nickname
		class = data.Class
	}
	if unit == nil {
		return engine.PlayerIdEmpty, response.WithStatus(ResponseStatusMalformed)
	}
	user := users.NewUser(nickname, userModel.Email, class, unit)
	if userModel.Unit == nil {
		c.persistUser(user)
	}
	c.users.AddUser(user)
	c.persistance.RemoveUserFromCache(data.Token)
	response.fillUserStatus(user)
	return user.Id, response.WithStatus(ResponseStatusOk)
}
