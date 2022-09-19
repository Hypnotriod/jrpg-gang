package users

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

type UserStatus int

type UserDisplayStatus string

const (
	UserDisplayStatusEmpty   UserDisplayStatus = ""
	UserDisplayStatusInLobby UserDisplayStatus = "inLobby"
	UserDisplayStatusInRoom  UserDisplayStatus = "inRoom"
	UserDisplayStatusInGame  UserDisplayStatus = "inGame"
)

const (
	UserStatusNotFound  UserStatus = 0
	UserStatusNotJoined UserStatus = (1 << 0)
	UserStatusJoined    UserStatus = (1 << 1)
	UserStatusInRoom    UserStatus = (1 << 2)
	UserStatusInGame    UserStatus = (1 << 3)
)

func (s UserStatus) Display() UserDisplayStatus {
	if s.Test(UserStatusJoined) {
		return UserDisplayStatusInLobby
	}
	if s.Test(UserStatusInGame) {
		return UserDisplayStatusInGame
	}
	if s.Test(UserStatusInRoom) {
		return UserDisplayStatusInRoom
	}
	return UserDisplayStatusEmpty
}

func (s UserStatus) Test(status UserStatus) bool {
	return s&status != 0
}

type User struct {
	engine.PlayerInfo
	RndGen *util.RndGen
	Status UserStatus
	Id     engine.UserId
	Unit   engine.GameUnit
}

func NewUser(nickname string,
	class engine.GameUnitClass,
	unit *engine.GameUnit) *User {
	u := &User{}
	u.RndGen = util.NewRndGen()
	u.Nickname = nickname
	u.Class = class
	u.Level = unit.Stats.Progress.Level
	u.Status = UserStatusNotJoined
	u.Unit = *unit
	u.Unit.Inventory.Prepare()
	u.Unit.Inventory.PopulateUids(u.RndGen)
	return u
}
