package controller

type GameStateRequest struct {
	Request
}

func (c *GameController) handleGameStateRequest(requestRaw string, response *Response) string {
	request := parseRequest(&GameStateRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	result, _, ok := c.engines.GameState(request.UserId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyGameState] = result
	return response.WithStatus(ResponseStatusOk)
}
