package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameController struct {
	*sync.RWMutex
	uidGen           util.UidGen
	engines          map[uint]*engine.GameEngine
	userIdToEngineId map[string]uint
}

func (c *GameController) HandleRequest(requestRaw string) string {
	request := c.parseRequest(requestRaw)
	response := Response{Status: ResponseStatusOk}
	if request != nil {
		c.processRequest(request, &response)
	} else {
		response.Status = ResponseStatusMailformed
	}
	return util.ObjectToJson(response)
}

func (c *GameController) parseRequest(requestRaw string) *Request {
	if request, ok := util.JsonToObject(&Request{}, requestRaw); ok {
		return request.(*Request)
	}
	return nil
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
