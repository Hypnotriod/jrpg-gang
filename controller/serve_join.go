package controller

import (
	"jrpg-gang/util"
	"regexp"
)

const USER_NICKNAME_REGEX string = `^[a-zA-Z][a-zA-Z0-9-_]+$`

func (c *GameController) serveJoin(requestRaw string, response *Response) string {
	request := parseJoinRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if matched, _ := regexp.MatchString(USER_NICKNAME_REGEX, request.Data.Nickname); !matched {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	defer c.Unlock()
	c.Lock()
	if _, ok := c.userNicknameToId[request.Data.Nickname]; ok {
		return response.WithStatus(ResponseStatusAlreadyExists)
	}
	userId := util.RandomId()
	c.userIdToNickname[userId] = request.Data.Nickname
	c.userNicknameToId[request.Data.Nickname] = userId
	response.UserId = userId
	return response.WithStatus(ResponseStatusOk)
}
