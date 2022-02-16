package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type User struct {
	Nickname string               `json:"nickname"`
	Class    engine.GameUnitClass `json:"class"`
	Level    uint                 `json:"level"`
	id       engine.UserId
	unit     *engine.GameUnit
}

func NewUser(nickname string,
	class engine.GameUnitClass,
	unit *engine.GameUnit) *User {
	u := &User{}
	u.Nickname = nickname
	u.Class = class
	u.Level = unit.Stats.Progress.Level
	u.unit = unit
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

func (s *Users) AddUser(user *User) {
	defer s.Unlock()
	s.Lock()
	user.id = engine.UserId(s.rndGen.Hash())
	s.users[user.id] = user
	s.userNicknameToId[user.Nickname] = user.id
}
