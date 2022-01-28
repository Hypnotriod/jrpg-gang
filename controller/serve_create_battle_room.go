package controller

import "jrpg-gang/engine"

func (c *GameController) serveCreateBattleRoom(requestRaw string, response *Response) *Response {
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
	for _, userId := range request.AllowedUsers {
		c.userIdToEngineId[userId] = engineUid
	}
	response.Data[DataKeyGameState] = engine
	return response.WithStatus(ResponseStatusOk)
}
