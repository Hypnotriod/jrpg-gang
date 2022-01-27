package controller

import (
	"jrpg-gang/engine"
)

func (c *GameController) processRequest(requestType RequestType, requestRaw string, response *Response) {
	switch requestType {
	case RequestCreateBattleRoom:
		c.processCreateBattleRoom(requestRaw, response)
	case RequestPlaceUnit:
		c.processPlaceUnit(requestRaw, response)
	}
}

func (c *GameController) processCreateBattleRoom(requestRaw string, response *Response) *Response {
	request := parseCreateBattleRoomRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	// todo: verify request data
	defer c.Unlock()
	c.Lock()
	battlefield := engine.NewBattlefield(request.Matrix)
	engine := engine.NewGameEngine(battlefield)
	engineUid := c.uidGen.Next()
	c.engines[engineUid] = engine
	for _, userUid := range request.AllowedUsers {
		c.userIdToEngineId[userUid] = engineUid
	}
	response.Data[DataKeyGameState] = engine
	return response.WithStatus(ResponseStatusOk)
}

func (c *GameController) processPlaceUnit(requestRaw string, response *Response) *Response {
	request := parsePlaceUnitRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	// todo: verify request data
	defer c.RUnlock()
	c.RLock()
	engine, ok := c.getEngineByUserId(request.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusError)
	}
	defer engine.Unlock()
	engine.Lock()
	actionResult := engine.Battlefield.PlaceUnit(&request.Unit)
	response.Data[DataKeyActionResult] = actionResult
	response.Data[DataKeyGameState] = engine
	return response.WithStatus(ResponseStatusOk)
}
