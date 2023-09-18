package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
)

type SwitchUnitRequestData struct {
	Class domain.UnitClass `json:"class,omitempty"`
}

func (c *GameController) handleSwitchUnitRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&SwitchUnitRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	userModel := c.persistance.GetUserByEmail(user.Email)
	if userModel == nil {
		return response.WithStatus(ResponseStatusNotFound)
	}
	var unit *engine.GameUnit
	u, ok := userModel.Units[data.Class]
	if ok {
		unit = engine.NewGameUnit(u)
		c.itemsConfig.PopulateFromDescriptor(&unit.Inventory)
	} else {
		unit = c.unitsConfig.GetByCode(domain.UnitCode(data.Class)) // todo: allow only specific unit codes
	}
	c.users.UpdateWithNewGameUnit(playerId, unit)
	response.Data[DataKeyUnit] = user.Unit
	return response.WithStatus(ResponseStatusOk)
}
