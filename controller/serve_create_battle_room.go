package controller

import "jrpg-gang/engine"

func (c *GameController) serveCreateBattleRoom(requestRaw string, response *Response) string {
	request := parseCreateBattleRoomRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	defer c.Unlock()
	c.Lock()
	if !c.isUserExists(request.Data.UserId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	battlefield := engine.NewBattlefield(request.Data.Matrix)
	engine := engine.NewGameEngine(battlefield)
	defer engine.Unlock()
	engine.Lock()
	engineUid := c.uidGen.Next()
	c.engines[engineUid] = engine
	for _, userId := range request.Data.AllowedUsers {
		c.userIdToEngineId[userId] = engineUid
	}
	response.Data[DataKeyGameState] = engine
	return response.WithStatus(ResponseStatusOk)
}
