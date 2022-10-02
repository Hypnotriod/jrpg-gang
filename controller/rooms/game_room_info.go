package rooms

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type GameRoomInfo struct {
	Uid         uint                  `json:"uid"`
	Capacity    uint                  `json:"capacity,omitempty"`
	ScenarioId  engine.GameScenarioId `json:"scenarioId,omitempty"`
	JoinedUsers []engine.PlayerInfo   `json:"joinedUsers"`
	Host        engine.PlayerInfo     `json:"host,omitempty"`
	Inactive    bool                  `json:"inactive,omitempty"`
}

func toInactiveGameRoomInfo(roomUid uint) GameRoomInfo {
	return GameRoomInfo{
		Uid:         roomUid,
		JoinedUsers: []engine.PlayerInfo{},
		Inactive:    true,
	}
}

func toGameRoomInfo(room *GameRoom) GameRoomInfo {
	return GameRoomInfo{
		Uid:         room.Uid,
		Host:        room.host.PlayerInfo,
		ScenarioId:  room.ScenarioId,
		Capacity:    room.Capacity,
		JoinedUsers: toPlayerInfos(room.joinedUsers),
	}
}

func toPlayerInfos(users []users.User) []engine.PlayerInfo {
	result := []engine.PlayerInfo{}
	for i := range users {
		result = append(result, users[i].PlayerInfo)
	}
	return result
}
