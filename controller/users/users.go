package users

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/persistance/model"
	"jrpg-gang/util"
	"sync"
)

type Users struct {
	mu                sync.RWMutex
	rndGen            *util.RndGen
	users             map[engine.PlayerId]*User
	userNicknameToId  map[string]engine.PlayerId
	userEmailToId     map[model.UserEmail]engine.PlayerId
	userSessionIdToId map[UserSessionId]engine.PlayerId
}

func NewUsers() *Users {
	u := &Users{}
	u.rndGen = util.NewRndGen()
	u.users = make(map[engine.PlayerId]*User)
	u.userNicknameToId = make(map[string]engine.PlayerId)
	u.userEmailToId = make(map[model.UserEmail]engine.PlayerId)
	u.userSessionIdToId = make(map[UserSessionId]engine.PlayerId)
	return u
}

func (u *Users) Get(playerId engine.PlayerId) (User, bool) {
	u.mu.RLock()
	user, ok := u.users[playerId]
	u.mu.RUnlock()
	if ok {
		return *user, ok
	} else {
		return User{}, ok
	}
}

func (u *Users) GetByIds(playerIds []engine.PlayerId) []User {
	u.mu.RLock()
	defer u.mu.RUnlock()
	result := []User{}
	for _, user := range u.users {
		if util.Contains(playerIds, user.Id) {
			result = append(result, *user)
		}
	}
	return result
}

func (u *Users) Has(playerId engine.PlayerId) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	_, exists := u.users[playerId]
	return exists
}

func (u *Users) TotalCount() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return len(u.users)
}

func (u *Users) GetByNickname(nickname string) (User, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	playerId, ok := u.userNicknameToId[nickname]
	if !ok {
		return User{}, false
	}
	user, ok := u.users[playerId]
	return *user, ok
}

func (u *Users) GetByEmail(email model.UserEmail) (User, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()
	playerId, ok := u.userEmailToId[email]
	if !ok {
		return User{}, false
	}
	user, ok := u.users[playerId]
	return *user, ok
}

func (u *Users) GetAndRefreshBySessionId(sessionId UserSessionId) (User, bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	playerId, ok := u.userSessionIdToId[sessionId]
	if !ok {
		return User{}, false
	}
	user, ok := u.users[playerId]
	delete(u.userSessionIdToId, user.SessionId)
	user.SessionId = UserSessionId(u.rndGen.MakeUUIDWithUniquenessCheck(func(value string) bool {
		_, ok := u.userSessionIdToId[UserSessionId(value)]
		return !ok
	}))
	u.userSessionIdToId[user.SessionId] = user.Id
	return *user, ok
}

func (u *Users) GetIdsByStatus(status UserStatus, onlineOnly bool) []engine.PlayerId {
	u.mu.RLock()
	defer u.mu.RUnlock()
	result := []engine.PlayerId{}
	for _, user := range u.users {
		if user.Status&status != 0 && (!onlineOnly || !user.IsOffline) {
			result = append(result, user.Id)
		}
	}
	return result
}

func (u *Users) GetIdsByStatusExcept(status UserStatus, playerId engine.PlayerId) []engine.PlayerId {
	u.mu.RLock()
	defer u.mu.RUnlock()
	result := []engine.PlayerId{}
	for _, user := range u.users {
		if user.Id != playerId && user.Status&status != 0 {
			result = append(result, user.Id)
		}
	}
	return result
}

func (u *Users) UpdateWithUnitOnGameComplete(playerId engine.PlayerId, unit *domain.Unit) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Unit.Stats.Progress = unit.Stats.Progress
	user.Unit.Booty.Accumulate(unit.Booty)
	user.Unit.Achievements = util.CloneMap(unit.Achievements)
	user.Unit.Inventory = *unit.Inventory.Clone()
	user.Unit.Inventory.PopulateUids(user.RndGen)
	user.Unit.PlayerInfo = nil
}

func (u *Users) UpdateWithNewGameUnit(playerId engine.PlayerId, unit *engine.GameUnit) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Unit = unit
	user.Class = unit.Class
	user.Unit.PrepareForUser()
	user.Unit.Inventory.PopulateUids(user.RndGen)
}

func (u *Users) AccumulateBooty(playerId engine.PlayerId, booty domain.UnitBooty) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Unit.Booty.Accumulate(booty)
}

func (u *Users) UpdateOnLevelUp(playerId engine.PlayerId, unit *domain.Unit) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Level = unit.Stats.Progress.Level
}

func (u *Users) ResetUser(playerId engine.PlayerId) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Status = UserStatusJoined
	user.RndGen.ResetUidGen()
}

func (u *Users) AddUser(user *User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user.Id = engine.PlayerId(u.rndGen.MakeUUIDWithUniquenessCheck(func(value string) bool {
		_, ok := u.users[engine.PlayerId(value)]
		return !ok
	}))
	user.SessionId = UserSessionId(u.rndGen.MakeUUIDWithUniquenessCheck(func(value string) bool {
		_, ok := u.userSessionIdToId[UserSessionId(value)]
		return !ok
	}))
	user.Status = UserStatusJoined
	u.users[user.Id] = user
	u.userNicknameToId[user.Nickname] = user.Id
	u.userEmailToId[user.Email] = user.Id
	u.userSessionIdToId[user.SessionId] = user.Id
}

func (u *Users) RemoveUser(playerId engine.PlayerId) *User {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return nil
	}
	delete(u.userNicknameToId, user.Nickname)
	delete(u.userEmailToId, user.Email)
	delete(u.userSessionIdToId, user.SessionId)
	delete(u.users, playerId)
	return user
}

func (u *Users) ChangeUserStatus(playerId engine.PlayerId, status UserStatus) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Status = status
}

func (u *Users) GetUserStatus(playerId engine.PlayerId) UserStatus {
	u.mu.RLock()
	user, ok := u.users[playerId]
	u.mu.RUnlock()
	if !ok {
		return UserStatusNotFound
	}
	return user.Status
}

func (u *Users) ConnectionStatusChanged(playerId engine.PlayerId, isOffline bool) {
	u.mu.Lock()
	defer u.mu.Unlock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.IsOffline = isOffline
}
