package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type UserId string

type User struct {
	Nickname string               `json:"nickname"`
	Class    engine.GameUnitClass `json:"class"`
	id       UserId
	unit     *engine.GameUnit
}

type Users struct {
	sync.RWMutex
	rndGen           *util.RndGen
	users            map[UserId]*User
	userNicknameToId map[string]UserId
}

func NewUsers() *Users {
	s := &Users{}
	s.rndGen = util.NewRndGen()
	s.users = make(map[UserId]*User)
	s.userNicknameToId = make(map[string]UserId)
	return s
}

func (s *Users) Get(userId UserId) (User, bool) {
	defer s.RUnlock()
	s.RLock()
	user, ok := s.users[userId]
	return *user, ok
}

func (s *Users) Has(userId UserId) bool {
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
	user.id = UserId(s.rndGen.Hash())
	s.users[user.id] = user
	s.userNicknameToId[user.Nickname] = user.id
}
