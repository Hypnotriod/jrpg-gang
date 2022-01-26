package controller

import "jrpg-gang/engine"

func (c *GameController) createGameEngine(request *Request, response *Response) {
	battlefield := engine.NewBattlefield(request.Data.Matrix)
	engine := engine.NewGameEngine(battlefield)
	engineUid := c.uidGen.Next()
	c.engines[engineUid] = engine
	for _, userUid := range request.Data.AllowedUsers {
		c.userIdToEngineId[userUid] = engineUid
	}
}
