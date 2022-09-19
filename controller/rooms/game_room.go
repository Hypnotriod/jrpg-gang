package rooms

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type GameRoomScenarioId string

type GameRoomInfo struct {
	Uid         uint                `json:"uid"`
	Capacity    uint                `json:"capacity"`
	ScenarioId  GameRoomScenarioId  `json:"scenarioId"`
	JoinedUsers []engine.PlayerInfo `json:"joinedUsers"`
	Host        engine.PlayerInfo   `json:"host"`
}

type GameRoom struct {
	Uid         uint               `json:"uid"`
	Capacity    uint               `json:"capacity"`
	ScenarioId  GameRoomScenarioId `json:"scenarioId"`
	joinedUsers []users.User
	host        users.User
}

func (r *GameRoom) IsFull() bool {
	return len(r.joinedUsers) >= int(r.Capacity)-1
}

func (r *GameRoom) GetUserIds() []engine.UserId {
	result := []engine.UserId{}
	for _, u := range r.joinedUsers {
		result = append(result, u.Id)
	}
	result = append(result, r.host.Id)
	return result
}

func (r *GameRoom) UpdateUserConnectionStatus(userId engine.UserId, isOffline bool) bool {
	if r.host.Id == userId {
		r.host.IsOffline = isOffline
		return true
	}
	for i := range r.joinedUsers {
		if r.joinedUsers[i].Id == userId {
			r.joinedUsers[i].IsOffline = isOffline
			return true
		}
	}
	return false
}

func (r *GameRoom) GetActors() []*engine.GameUnit {
	result := []*engine.GameUnit{}
	r.host.Unit.PlayerInfo = &r.host.PlayerInfo
	r.host.Unit.PlayerInfo.IsHost = true
	result = append(result, &r.host.Unit)
	for i := range r.joinedUsers {
		u := r.joinedUsers[i]
		u.Unit.PlayerInfo = &u.PlayerInfo
		result = append(result, &u.Unit)
	}
	return result
}

func NewGameRoom(capacity uint, scenarioId GameRoomScenarioId, host users.User) *GameRoom {
	r := &GameRoom{}
	r.Capacity = capacity
	r.ScenarioId = scenarioId
	r.host = host
	r.joinedUsers = []users.User{}
	return r
}
