package controller

import "sync"

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

type Users struct {
	sync.RWMutex
	users            map[string]*User
	userNicknameToId map[string]string
}

func NewUsers() *Users {
	s := &Users{}
	s.users = make(map[string]*User)
	s.userNicknameToId = make(map[string]string)
	return s
}

func (s *Users) Get(userId string) (User, bool) {
	defer s.RUnlock()
	s.RLock()
	user, ok := s.users[userId]
	return *user, ok
}

func (s *Users) Has(userId string) bool {
	defer s.RUnlock()
	s.RLock()
	_, exists := s.users[userId]
	return exists
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
	s.users[user.Id] = user
	s.userNicknameToId[user.Nickname] = user.Id
}
