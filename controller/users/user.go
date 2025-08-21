package users

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
	"jrpg-gang/util"
	"net"
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
	UserDisplayStatusAtJob   UserDisplayStatus = "atJob"
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
	model.UserCredentials
	RndGen    *util.RndGen
	SessionId UserSessionId
	Status    UserStatus
	Unit      *engine.GameUnit
	Ip        net.IP
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
	if s.Test(UserStatusAtJob) {
		return UserDisplayStatusAtJob
	}
	return UserDisplayStatusEmpty
}

func (s UserStatus) Test(status UserStatus) bool {
	return s&status != 0
}

func NewUser(
	ip net.IP,
	nickname string,
	email model.UserEmail,
	userId model.UserId,
	class domain.UnitClass,
	unit *engine.GameUnit) *User {
	u := &User{}
	u.Ip = ip
	u.RndGen = util.NewRndGen()
	u.Nickname = nickname
	u.Email = email
	u.UserId = userId
	u.Class = class
	u.Level = unit.Stats.Progress.Level
	u.Status = UserStatusNotJoined
	u.Unit = unit
	u.Unit.PrepareForUser()
	u.Unit.Inventory.PopulateUids(u.RndGen)
	return u
}
