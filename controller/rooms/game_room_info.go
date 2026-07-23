package rooms

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
)

type GameRoomInfo struct {
	Uid         uint                  `json:"uid"`
	Capacity    uint                  `json:"capacity,omitempty"`
	ScenarioId  engine.GameScenarioId `json:"scenarioId,omitempty"`
	Host        engine.PlayerInfo     `json:"host"`
	JoinedUsers []engine.PlayerInfo   `json:"joinedUsers"`
	Mercenaries []engine.PlayerInfo   `json:"mercenaries"`
	Inactive    bool                  `json:"inactive,omitempty"`
}

func toInactiveGameRoomInfo(roomUid uint) GameRoomInfo {
	return GameRoomInfo{
		Uid:         roomUid,
		JoinedUsers: []engine.PlayerInfo{},
		Mercenaries: []engine.PlayerInfo{},
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
		Mercenaries: toMercenaryInfos(room.mercenaries),
	}
}

func toPlayerInfos(users []users.User) []engine.PlayerInfo {
	result := []engine.PlayerInfo{}
	for i := range users {
		result = append(result, users[i].PlayerInfo)
	}
	return result
}

func toMercenaryInfos(mercenaries []*engine.GameUnit) []engine.PlayerInfo {
	result := []engine.PlayerInfo{}
	for i := range mercenaries {
		result = append(result, engine.PlayerInfo{
			Nickname: mercenaries[i].Name,
			Class:    mercenaries[i].Class,
			Code:     mercenaries[i].Code,
			Level:    mercenaries[i].Stats.Progress.Level,
		})
	}
	return result
}
