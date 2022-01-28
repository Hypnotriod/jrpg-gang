package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameController struct {
	sync.RWMutex
	uidGen           util.UidGen
	engines          map[uint]*engine.GameEngine
	userIdToNickname map[string]string
	userNicknameToId map[string]string
	userIdToEngineId map[string]uint
}

func NewController() *GameController {
	c := &GameController{}
	c.engines = make(map[uint]*engine.GameEngine)
	c.userIdToNickname = make(map[string]string)
	c.userNicknameToId = make(map[string]string)
	c.userIdToEngineId = make(map[string]uint)
	return c
}

func (c *GameController) HandleRequest(requestRaw string) string {
	response := NewResponse()
	request := parseRequest(requestRaw)
	if request != nil {
		response.Type = request.Type
		response.Id = request.Id
		return c.serveRequest(request.Type, requestRaw, response)
	}
	return response.WithStatus(ResponseStatusMailformed)
}

func (c *GameController) serveRequest(requestType RequestType, requestRaw string, response *Response) string {
	switch requestType {
	case RequestJoin:
		return c.serveJoin(requestRaw, response)
	case RequestCreateBattleRoom:
		return c.serveCreateBattleRoom(requestRaw, response)
	case RequestPlaceUnit:
		return c.servePlaceUnit(requestRaw, response)
	}
	return response.WithStatus(ResponseStatusUnsupported)
}

func (c *GameController) getEngineByUserId(userId string) (*engine.GameEngine, bool) {
	engineUid, ok := c.userIdToEngineId[userId]
	if !ok {
		return nil, false
	}
	engine, ok := c.engines[engineUid]
	if !ok {
		return nil, false
	}
	return engine, true
}

func (c *GameController) isUserExists(userId string) bool {
	_, r := c.userIdToNickname[userId]
	return r
}
