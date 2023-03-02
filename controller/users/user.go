package users

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
)

type UserSessionId string

const (
	UserSessionIdEmpty UserSessionId = ""
)

type UserStatus int

type UserDisplayStatus string

const (
	UserDisplayStatusEmpty   UserDisplayStatus = ""
	UserDisplayStatusJoined  UserDisplayStatus = "joined"
	UserDisplayStatusInLobby UserDisplayStatus = "inLobby"
	UserDisplayStatusInRoom  UserDisplayStatus = "inRoom"
	UserDisplayStatusInGame  UserDisplayStatus = "inGame"
)

const (
	UserStatusNotFound  UserStatus = 0
	UserStatusNotJoined UserStatus = (1 << 0)
	UserStatusJoined    UserStatus = (1 << 1)
	UserStatusInLobby   UserStatus = (1 << 2)
	UserStatusInRoom    UserStatus = (1 << 3)
	UserStatusInGame    UserStatus = (1 << 4)
	UserStatusAtJob     UserStatus = (1 << 5)
)

type User struct {
	engine.PlayerInfo
	RndGen    *util.RndGen
	Email     string
	SessionId UserSessionId
	Status    UserStatus
	Unit      *engine.GameUnit
}

func (s UserStatus) Display() UserDisplayStatus {
	if s.Test(UserStatusJoined) {
		return UserDisplayStatusJoined
	}
	if s.Test(UserStatusInLobby) {
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

func NewUser(nickname string,
	email string,
	class domain.UnitClass,
	unit *engine.GameUnit) *User {
	u := &User{}
	u.RndGen = util.NewRndGen()
	u.Nickname = nickname
	u.Email = email
	u.Class = class
	u.Level = unit.Stats.Progress.Level
	u.Status = UserStatusNotJoined
	u.Unit = unit
	u.Unit.PrepareForUser()
	u.Unit.Inventory.PopulateUids(u.RndGen)
	return u
}
