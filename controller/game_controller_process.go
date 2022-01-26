package controller

import "jrpg-gang/engine"

func (c *GameController) processRequest(request *Request, response *Response) {
	switch request.Type {
	case RequestCreateGameEngine:
		c.processCreateGameEngine(request, response)
	case RequestJoin:
		c.processJoin(request, response)
	}
}

func (c *GameController) processCreateGameEngine(request *Request, response *Response) {
	battlefield := engine.NewBattlefield(request.Data.Matrix)
	engine := engine.NewGameEngine(battlefield)
	engineUid := c.uidGen.Next()
	defer c.Unlock()
	c.Lock()
	c.engines[engineUid] = engine
	for _, userUid := range request.Data.AllowedUsers {
		c.userIdToEngineId[userUid] = engineUid
	}
}

func (c *GameController) processJoin(request *Request, response *Response) {
	// todo
}
