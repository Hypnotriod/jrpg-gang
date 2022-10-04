package users

import (
	"jrpg-gang/domain"
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
	u := &Users{}
	u.rndGen = util.NewRndGen()
	u.users = make(map[engine.UserId]*User)
	u.userNicknameToId = make(map[string]engine.UserId)
	return u
}

func (u *Users) Get(userId engine.UserId) (User, bool) {
	u.mu.RLock()
	user, ok := u.users[userId]
	u.mu.RUnlock()
	if ok {
		return *user, ok
	} else {
		return User{}, ok
	}
}

func (u *Users) GetByIds(userIds []engine.UserId) []User {
	defer u.mu.RUnlock()
	u.mu.RLock()
	result := []User{}
	for _, user := range u.users {
		if util.Contains(userIds, user.Id) {
			result = append(result, *user)
		}
	}
	return result
}

func (u *Users) Has(userId engine.UserId) bool {
	defer u.mu.RUnlock()
	u.mu.RLock()
	_, exists := u.users[userId]
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
	userId, ok := u.userNicknameToId[nickname]
	if !ok {
		return User{}, false
	}
	user, ok := u.users[userId]
	return *user, ok
}

func (u *Users) GetIdsByStatus(status UserStatus, onlineOnly bool) []engine.UserId {
	defer u.mu.RUnlock()
	u.mu.RLock()
	result := []engine.UserId{}
	for _, user := range u.users {
		if user.Status&status != 0 && (!onlineOnly || !user.IsOffline) {
			result = append(result, user.Id)
		}
	}
	return result
}

func (u *Users) GetIdsByStatusExcept(status UserStatus, userId engine.UserId) []engine.UserId {
	defer u.mu.RUnlock()
	u.mu.RLock()
	result := []engine.UserId{}
	for _, user := range u.users {
		if user.Id != userId && user.Status&status != 0 {
			result = append(result, user.Id)
		}
	}
	return result
}

func (u *Users) UpdateWithUnitOnGameComplete(userId engine.UserId, unit *domain.Unit) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[userId]
	if !ok {
		return
	}
	user.Unit.Stats.Progress = unit.Stats.Progress
	user.Unit.Booty.Accumulate(unit.Booty)
	user.Unit.Inventory = *unit.Inventory.Clone()
	user.Unit.Inventory.Prepare()
	user.Unit.Inventory.PopulateUids(user.RndGen)
}

func (u *Users) ResetUser(userId engine.UserId) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[userId]
	if !ok {
		return
	}
	user.Status = UserStatusJoined
	user.RndGen.ResetUidGen()
}

func (u *Users) AddUser(user *User) {
	defer u.mu.Unlock()
	u.mu.Lock()
	for {
		user.Id = engine.UserId(u.rndGen.MakeId())
		if _, ok := u.users[user.Id]; !ok {
			break
		}
	}
	user.Status = UserStatusJoined
	u.users[user.Id] = user
	u.userNicknameToId[user.Nickname] = user.Id
}

func (u *Users) RemoveUser(userId engine.UserId) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[userId]
	if !ok {
		return
	}
	delete(u.userNicknameToId, user.Nickname)
	delete(u.users, userId)
}

func (u *Users) ChangeUserStatus(userId engine.UserId, status UserStatus) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[userId]
	if !ok {
		return
	}
	user.Status = status
}

func (u *Users) GetUserStatus(userId engine.UserId) UserStatus {
	u.mu.RLock()
	user, ok := u.users[userId]
	u.mu.RUnlock()
	if !ok {
		return UserStatusNotFound
	}
	return user.Status
}

func (u *Users) ConnectionStatusChanged(userId engine.UserId, isOffline bool) {
	defer u.mu.Unlock()
	u.mu.Lock()
	user, ok := u.users[userId]
	if !ok {
		return
	}
	user.IsOffline = isOffline
}
