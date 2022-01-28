package controller

func (c *GameController) servePlaceUnit(requestRaw string, response *Response) string {
	request := parsePlaceUnitRequest(requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	// todo: verify request data
	defer c.RUnlock()
	c.RLock()
	engine, ok := c.getEngineByUserId(request.Data.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusError)
	}
	defer engine.Unlock()
	engine.Lock()
	actionResult := engine.Battlefield.PlaceUnit(&request.Data.Unit)
	response.Data[DataKeyActionResult] = actionResult
	response.Data[DataKeyGameState] = engine
	return response.WithStatus(ResponseStatusOk)
}
