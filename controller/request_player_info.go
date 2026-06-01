package controller

import (
	"jrpg-gang/auth"
	"jrpg-gang/controller/users"
	"jrpg-gang/domain"
	"jrpg-gang/persistance/model"
	"regexp"
)

type SetPlayerInfoRequestData struct {
	Token    auth.AuthenticationToken `json:"token"`
	Nickname string                   `json:"nickname"`
	Class    domain.UnitClass         `json:"class"`
}

func (c *GameController) handleSetPlayerInfoRequest(request *Request, response *Response) []byte {
	data := parseRequestData(&SetPlayerInfoRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	userModel, ok := c.persistance.GetUserFromAuthCache(data.Token)
	if userModel == nil || !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if len(userModel.Units) != 0 {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	nickname := data.Nickname
	if userModel.IsGuest {
		nickname = userModel.Nickname
	} else {
		if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, nickname); !matched {
			return response.WithStatus(ResponseStatusMalformed)
		}
		if ok := c.persistance.HasUserWithNickname(nickname); ok {
			return response.WithStatus(ResponseStatusAlreadyExists)
		}
	}
	unit := c.unitsConfig.GetByClass(domain.UnitClass(data.Class))
	if unit == nil {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if !c.persistance.SetUserInfoToAuthCache(data.Token, nickname, data.Class, &unit.Unit) {
		return response.WithStatus(ResponseStatusNotFound)
	}
	userId := model.UserId(userModel.Id.Hex())
	user := users.NewUser(nickname, userModel.Email, userModel.IsGuest, userId, data.Class, unit)
	c.persistUser(user)
	return response.WithStatus(ResponseStatusOk)
}
