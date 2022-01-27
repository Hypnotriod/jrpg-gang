package controller

import "jrpg-gang/engine"

func (c *GameController) processRequest(request *Request, response *Response) {
	switch request.Type {
	case RequestCreateBattleRoom:
		c.processCreateBattleRoom(request, response)
	case RequestplaceUnit:
		c.processPlaceUnit(request, response)
	}
}

func (c *GameController) processCreateBattleRoom(request *Request, response *Response) *Response {
	// todo: verify request data
	defer c.Unlock()
	c.Lock()
	battlefield := engine.NewBattlefield(request.Data.Matrix)
	engine := engine.NewGameEngine(battlefield)
	engineUid := c.uidGen.Next()
	c.engines[engineUid] = engine
	for _, userUid := range request.Data.AllowedUsers {
		c.userIdToEngineId[userUid] = engineUid
	}
	response.GameState = *engine
	return response.WithStatus(ResponseStatusOk)
}

func (c *GameController) processPlaceUnit(request *Request, response *Response) *Response {
	// todo: verify request data
	defer c.RUnlock()
	c.RLock()
	engine, ok := c.getEngineByUserId(request.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusError)
	}
	defer engine.Unlock()
	engine.Lock()
	response.ActionResult = engine.Battlefield.PlaceUnit(&request.Data.Unit)
	response.GameState = *engine
	return response.WithStatus(ResponseStatusOk)
}
