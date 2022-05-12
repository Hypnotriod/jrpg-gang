package controller

type GameNextPhaseRequest struct {
	Request
}

func (c *GameController) handleGameNextPhaseRequest(requestRaw string, response *Response) string {
	request := parseRequest(&GameNextPhaseRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	result, broadcastUserIds, ok := c.engines.NextPhase(request.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyActionResult] = result
	if len(broadcastUserIds) > 0 {
		c.broadcastGameAction(broadcastUserIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
