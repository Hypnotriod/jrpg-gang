package users

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type Users struct {
	mu               sync.RWMutex
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
	s.mu.RLock()
	user, ok := s.users[userId]
	s.mu.RUnlock()
	if ok {
		return *user, ok
	} else {
		return User{}, ok
	}
}

func (s *Users) GetByIds(userIds []engine.UserId) []User {
	defer s.mu.RUnlock()
	s.mu.RLock()
	result := []User{}
	for _, user := range s.users {
		if util.Contains(userIds, user.Id) {
			result = append(result, *user)
		}
	}
	return result
}

func (s *Users) Has(userId engine.UserId) bool {
	defer s.mu.RUnlock()
	s.mu.RLock()
	_, exists := s.users[userId]
	return exists
}

func (s *Users) TotalCount() int {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return len(s.users)
}

func (s *Users) GetByNickname(nickname string) (User, bool) {
	defer s.mu.RUnlock()
	s.mu.RLock()
	userId, ok := s.userNicknameToId[nickname]
	if !ok {
		return User{}, false
	}
	user, ok := s.users[userId]
	return *user, ok
}

func (s *Users) GetIdsByStatus(status UserStatus, onlineOnly bool) []engine.UserId {
	defer s.mu.RUnlock()
	s.mu.RLock()
	result := []engine.UserId{}
	for _, user := range s.users {
		if user.Status&status != 0 && (!onlineOnly || !user.IsOffline) {
			result = append(result, user.Id)
		}
	}
	return result
}

func (s *Users) GetIdsByStatusExcept(status UserStatus, userId engine.UserId) []engine.UserId {
	defer s.mu.RUnlock()
	s.mu.RLock()
	result := []engine.UserId{}
	for _, user := range s.users {
		if user.Id != userId && user.Status&status != 0 {
			result = append(result, user.Id)
		}
	}
	return result
}

func (s *Users) UpdateUnit(userId engine.UserId, unit *engine.GameUnit) {
	defer s.mu.Unlock()
	s.mu.Lock()
	user, ok := s.users[userId]
	if !ok {
		return
	}
	user.Unit = *unit
}

func (s *Users) AddUser(user *User) {
	defer s.mu.Unlock()
	s.mu.Lock()
	for {
		user.Id = engine.UserId(s.rndGen.Hash())
		if _, ok := s.users[user.Id]; !ok {
			break
		}
	}
	user.Unit.SetUserId(user.Id)
	user.Status = UserStatusJoined
	s.users[user.Id] = user
	s.userNicknameToId[user.Nickname] = user.Id
}

func (s *Users) RemoveUser(userId engine.UserId) {
	defer s.mu.Unlock()
	s.mu.Lock()
	user, ok := s.users[userId]
	if !ok {
		return
	}
	delete(s.userNicknameToId, user.Nickname)
	delete(s.users, userId)
}

func (s *Users) ChangeUserStatus(userId engine.UserId, status UserStatus) {
	defer s.mu.Unlock()
	s.mu.Lock()
	user, ok := s.users[userId]
	if !ok {
		return
	}
	user.Status = status
}

func (s *Users) GetUserStatus(userId engine.UserId) UserStatus {
	s.mu.RLock()
	user, ok := s.users[userId]
	s.mu.RUnlock()
	if !ok {
		return UserStatusNotFound
	}
	return user.Status
}

func (s *Users) ConnectionStatusChanged(userId engine.UserId, isOffline bool) {
	defer s.mu.Unlock()
	s.mu.Lock()
	user, ok := s.users[userId]
	if !ok {
		return
	}
	user.IsOffline = isOffline
}
