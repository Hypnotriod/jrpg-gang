package users

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type Users struct {
	mu                sync.RWMutex
	rndGen            *util.RndGen
	users             map[engine.PlayerId]*User
	userNicknameToId  map[string]engine.PlayerId
	userEmailToId     map[string]engine.PlayerId
	userSessionIdToId map[UserSessionId]engine.PlayerId
}

func NewUsers() *Users {
	u := &Users{}
	u.rndGen = util.NewRndGen()
	u.users = make(map[engine.PlayerId]*User)
	u.userNicknameToId = make(map[string]engine.PlayerId)
	u.userEmailToId = make(map[string]engine.PlayerId)
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
	defer u.mu.RUnlock()
	u.mu.RLock()
	result := []User{}
	for _, user := range u.users {
		if util.Contains(playerIds, user.Id) {
			result = append(result, *user)
		}
	}
	return result
}

func (u *Users) Has(playerId engine.PlayerId) bool {
	defer u.mu.RUnlock()
	u.mu.RLock()
	_, exists := u.users[playerId]
	return exists
}

func (u *Users) TotalCount() int {
	defer u.mu.RUnlock()
	u.mu.RLock()
	return len(u.users)
}

func (u *Users) GetByNickname(nickname string) (User, bool) {
	defer u.mu.RUnlock()
	u.mu.RLock()
	playerId, ok := u.userNicknameToId[nickname]
	if !ok {
		return User{}, false
	}
	user, ok := u.users[playerId]
	return *user, ok
}

func (u *Users) GetByEmail(email string) (User, bool) {
	defer u.mu.RUnlock()
	u.mu.RLock()
	playerId, ok := u.userEmailToId[email]
	if !ok {
		return User{}, false
	}
	user, ok := u.users[playerId]
	return *user, ok
}

func (u *Users) GetAndRefreshBySessionId(sessionId UserSessionId) (User, bool) {
	defer u.mu.Unlock()
	u.mu.Lock()
	playerId, ok := u.userSessionIdToId[sessionId]
	if !ok {
		return User{}, false
	}
	user, ok := u.users[playerId]
	delete(u.userSessionIdToId, user.SessionId)
	user.SessionId = UserSessionId(u.rndGen.MakeUUID())
	u.userSessionIdToId[user.SessionId] = user.Id
	return *user, ok
}

func (u *Users) GetIdsByStatus(status UserStatus, onlineOnly bool) []engine.PlayerId {
	defer u.mu.RUnlock()
	u.mu.RLock()
	result := []engine.PlayerId{}
	for _, user := range u.users {
		if user.Status&status != 0 && (!onlineOnly || !user.IsOffline) {
			result = append(result, user.Id)
		}
	}
	return result
}

func (u *Users) GetIdsByStatusExcept(status UserStatus, playerId engine.PlayerId) []engine.PlayerId {
	defer u.mu.RUnlock()
	u.mu.RLock()
	result := []engine.PlayerId{}
	for _, user := range u.users {
		if user.Id != playerId && user.Status&status != 0 {
			result = append(result, user.Id)
		}
	}
	return result
}

func (u *Users) UpdateWithUnitOnGameComplete(playerId engine.PlayerId, unit *domain.Unit) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Unit.Stats.Progress = unit.Stats.Progress
	user.Unit.Booty.Accumulate(unit.Booty)
	user.Unit.Inventory = *unit.Inventory.Clone()
	user.Unit.Inventory.PopulateUids(user.RndGen)
}

func (u *Users) UpdateOnLevelUp(playerId engine.PlayerId, unit *domain.Unit) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Level = user.Unit.PlayerInfo.Level
}

func (u *Users) ResetUser(playerId engine.PlayerId) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.Status = UserStatusJoined
	user.RndGen.ResetUidGen()
}

func (u *Users) AddUser(user *User) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user.Id = engine.PlayerId(u.rndGen.MakeUUID())
	user.SessionId = UserSessionId(u.rndGen.MakeUUID())
	user.Status = UserStatusJoined
	u.users[user.Id] = user
	u.userNicknameToId[user.Nickname] = user.Id
	u.userEmailToId[user.Email] = user.Id
	u.userSessionIdToId[user.SessionId] = user.Id
}

func (u *Users) RemoveUser(playerId engine.PlayerId) *User {
	defer u.mu.Unlock()
	u.mu.Lock()
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
	defer u.mu.Unlock()
	u.mu.Lock()
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
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[playerId]
	if !ok {
		return
	}
	user.IsOffline = isOffline
}
