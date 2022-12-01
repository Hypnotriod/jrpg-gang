package rooms

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type GameRoom struct {
	Uid         uint                  `json:"uid"`
	Capacity    uint                  `json:"capacity"`
	ScenarioId  engine.GameScenarioId `json:"scenarioId"`
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
	hostUnit := r.host.Unit.Clone()
	hostUnit.PlayerInfo = r.host.PlayerInfo.Clone()
	result = append(result, hostUnit)
	for i := range r.joinedUsers {
		u := r.joinedUsers[i]
		unit := u.Unit.Clone()
		unit.PlayerInfo = u.PlayerInfo.Clone()
		result = append(result, unit)
	}
	return result
}

func newGameRoom() *GameRoom {
	r := &GameRoom{}
	r.joinedUsers = []users.User{}
	return r
}
