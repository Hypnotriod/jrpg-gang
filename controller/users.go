package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type UserStatus int

const (
	UserStatusNotJoined UserStatus = (1 << 0)
	UserStatusJoined    UserStatus = (1 << 1)
	UserStatusInRoom    UserStatus = (1 << 2)
	UserStatusInGame    UserStatus = (1 << 3)
)

type User struct {
	Nickname string               `json:"nickname"`
	Class    engine.GameUnitClass `json:"class"`
	Level    uint                 `json:"level"`
	status   UserStatus
	id       engine.UserId
	unit     engine.GameUnit
}

func NewUser(nickname string,
	class engine.GameUnitClass,
	unit *engine.GameUnit) *User {
	u := &User{}
	u.Nickname = nickname
	u.Class = class
	u.Level = unit.Stats.Progress.Level
	u.status = UserStatusNotJoined
	u.unit = *unit
	return u
}

type Users struct {
	sync.RWMutex
	rndGen           *util.RndGen
	users            map[engine.UserId]*User
	userNicknameToId map[string]engine.UserId
}

func NewUsers() *Users {
	s := &Users{}
	s.rndGen = util.NewRndGen()
	s.users = make(map[engine.UserId]*User)
	s.userNicknameToId = make(map[string]engine.UserId)
	return s
}

func (s *Users) Get(userId engine.UserId) (User, bool) {
	defer s.RUnlock()
	s.RLock()
	user, ok := s.users[userId]
	return *user, ok
}

func (s *Users) Has(userId engine.UserId) bool {
	defer s.RUnlock()
	s.RLock()
	_, exists := s.users[userId]
	return exists
}

func (s *Users) TotalCount() int {
	defer s.RUnlock()
	s.RLock()
	return len(s.users)
}

func (s *Users) GetByNickname(nickname string) (User, bool) {
	defer s.RUnlock()
	s.RLock()
	userId, ok := s.userNicknameToId[nickname]
	if !ok {
		return User{}, false
	}
	user, ok := s.users[userId]
	return *user, ok
}

func (s *Users) GetIdsByStatus(status UserStatus) []engine.UserId {
	defer s.RUnlock()
	s.RLock()
	result := []engine.UserId{}
	for _, user := range s.users {
		if user.status&status != 0 {
			result = append(result, user.id)
		}
	}
	return result
}

func (s *Users) GetIdsByStatusExcept(status UserStatus, userId engine.UserId) []engine.UserId {
	defer s.RUnlock()
	s.RLock()
	result := []engine.UserId{}
	for _, user := range s.users {
		if user.id != userId && user.status&status != 0 {
			result = append(result, user.id)
		}
	}
	return result
}

func (s *Users) AddUser(user *User) {
	defer s.Unlock()
	s.Lock()
	for {
		user.id = engine.UserId(s.rndGen.Hash())
		if _, ok := s.users[user.id]; !ok {
			break
		}
	}
	user.unit.UserId = user.id
	user.status = UserStatusJoined
	s.users[user.id] = user
	s.userNicknameToId[user.Nickname] = user.id
}

func (s *Users) RemoveUser(userId engine.UserId) {
	defer s.Unlock()
	s.Lock()
	user, ok := s.users[userId]
	if !ok {
		return
	}
	delete(s.userNicknameToId, user.Nickname)
	delete(s.users, userId)
}

func (s *Users) ChangeUserStatus(userId engine.UserId, status UserStatus) {
	defer s.Unlock()
	s.Lock()
	user, ok := s.users[userId]
	if !ok {
		return
	}
	user.status = status
}

func (s *Users) TestUserStatus(userId engine.UserId, status UserStatus) bool {
	s.RLock()
	user, ok := s.users[userId]
	s.RUnlock()
	if !ok {
		return false
	}
	return user.status&status != 0
}
