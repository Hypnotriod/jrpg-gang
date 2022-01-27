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

func NewGameController() *GameController {
	controller := &GameController{}
	controller.RWMutex = &sync.RWMutex{}
	controller.engines = make(map[uint]*engine.GameEngine)
	controller.userIdToEngineId = make(map[string]uint)
	return controller
}

func (c *GameController) HandleRequest(requestRaw string) string {
	response := NewResponse()
	request := parseRequest(requestRaw)
	if request != nil {
		c.processRequest(request.Type, requestRaw, response)
	} else {
		response.Status = ResponseStatusMailformed
	}
	return util.ObjectToJson(response)
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
