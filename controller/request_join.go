package controller

import (
	"jrpg-gang/util"
	"regexp"
)

type JoinRequest struct {
	Request
	Data struct {
		Nickname string `json:"nickname"`
	} `json:"data"`
}

func parseJoinRequest(requestRaw string) *JoinRequest {
	if r, ok := util.JsonToObject(&JoinRequest{}, requestRaw); ok {
		return r.(*JoinRequest)
	}
	return nil
}

const USER_NICKNAME_REGEX string = `^[a-zA-Z][a-zA-Z0-9-_]+$`

func (c *GameController) handleJoinRequest(requestRaw string, response *Response) string {
	request := parseJoinRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, request.Data.Nickname); !matched {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if _, ok := c.users.GetByNickname(request.Data.Nickname); ok {
		return response.WithStatus(ResponseStatusAlreadyExists)
	}
	user := &User{
		Id:       util.RandomId(),
		Nickname: request.Data.Nickname,
	}
	c.users.AddUser(user)
	response.Data[DataKeyUserId] = user.Id
	return response.WithStatus(ResponseStatusOk)
}
