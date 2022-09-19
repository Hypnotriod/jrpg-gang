package controller

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameRoomScenarioId string

type GameRoomInfo struct {
	Uid         uint                `json:"uid"`
	Capacity    uint                `json:"capacity"`
	ScenarioId  GameRoomScenarioId  `json:"scenarioId"`
	JoinedUsers []engine.PlayerInfo `json:"joinedUsers"`
	Host        engine.PlayerInfo   `json:"host"`
}

type GameRoom struct {
	Uid         uint               `json:"uid"`
	Capacity    uint               `json:"capacity"`
	ScenarioId  GameRoomScenarioId `json:"scenarioId"`
	joinedUsers []users.User
	host        users.User
}

func (r *GameRoom) IsFull() bool {
	return len(r.joinedUsers) >= int(r.Capacity)-1
}

func (r *GameRoom) GetUserIds() []engine.UserId {
	result := []engine.UserId{}
	for _, u := range r.joinedUsers {
		result = append(result, u.Id)
	}
	result = append(result, r.host.Id)
	return result
}

func (r *GameRoom) UpdateUserConnectionStatus(userId engine.UserId, isOffline bool) bool {
	if r.host.Id == userId {
		r.host.IsOffline = isOffline
		return true
	}
	for i := range r.joinedUsers {
		if r.joinedUsers[i].Id == userId {
			r.joinedUsers[i].IsOffline = isOffline
			return true
		}
	}
	return false
}

func (r *GameRoom) GetActors() []*engine.GameUnit {
	result := []*engine.GameUnit{}
	r.host.Unit.PlayerInfo = &r.host.PlayerInfo
	r.host.Unit.PlayerInfo.IsHost = true
	result = append(result, &r.host.Unit)
	for i := range r.joinedUsers {
		u := r.joinedUsers[i]
		u.Unit.PlayerInfo = &u.PlayerInfo
		result = append(result, &u.Unit)
	}
	return result
}

func NewGameRoom() *GameRoom {
	r := &GameRoom{}
	r.joinedUsers = []users.User{}
	return r
}

type GameRooms struct {
	mu              sync.RWMutex
	rndGen          *util.RndGen
	rooms           map[uint]*GameRoom
	userIdToRoomUid map[engine.UserId]uint
}

func NewGameRooms() *GameRooms {
	r := &GameRooms{}
	r.rndGen = util.NewRndGen()
	r.rooms = make(map[uint]*GameRoom)
	r.userIdToRoomUid = make(map[engine.UserId]uint)
	return r
}

func (r *GameRooms) Add(room *GameRoom) {
	defer r.mu.Unlock()
	r.mu.Lock()
	room.Uid = r.rndGen.NextUid()
	r.rooms[room.Uid] = room
	r.userIdToRoomUid[room.host.Id] = room.Uid
}

func (r *GameRooms) PopByHostId(hostId engine.UserId) (*GameRoom, bool) {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.userIdToRoomUid[hostId]
	if !ok {
		return nil, false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return nil, false
	}
	if room.host.Id != hostId {
		return nil, false
	}
	for _, u := range room.joinedUsers {
		delete(r.userIdToRoomUid, u.Id)
	}
	delete(r.userIdToRoomUid, hostId)
	delete(r.rooms, roomUid)
	return room, true
}

func (r *GameRooms) AddUser(roomUid uint, user users.User) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	room, ok := r.rooms[roomUid]
	if !ok {
		return false
	}
	if room.IsFull() {
		return false
	}
	room.joinedUsers = append(room.joinedUsers, user)
	r.userIdToRoomUid[user.Id] = roomUid
	return true
}

func (r *GameRooms) RemoveUser(userId engine.UserId) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return false
	}
	if room.host.Id == userId {
		return false
	}
	delete(r.userIdToRoomUid, userId)
	restUsers := []users.User{}
	for _, u := range room.joinedUsers {
		if u.Id != userId {
			restUsers = append(restUsers, u)
		}
	}
	room.joinedUsers = restUsers
	return true
}

func (r *GameRooms) Has(uid uint) bool {
	defer r.mu.RUnlock()
	r.mu.RLock()
	_, ok := r.rooms[uid]
	return ok
}

func (r *GameRooms) ExistsForUserId(userId engine.UserId) bool {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return false
	}
	_, present := r.rooms[roomUid]
	return present
}

func (r *GameRooms) ExistsForHostId(hostId engine.UserId) bool {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.userIdToRoomUid[hostId]
	if !ok {
		return false
	}
	room, present := r.rooms[roomUid]
	if !present {
		return false
	}
	return room.host.Id == hostId
}

func (r *GameRooms) GetByUserId(userId engine.UserId) (GameRoom, bool) {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return GameRoom{}, false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return GameRoom{}, false
	}
	return *room, ok
}

func (r *GameRooms) ConnectionStatusChanged(userId engine.UserId, isOffline bool) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return false
	}
	return room.UpdateUserConnectionStatus(userId, isOffline)
}

func (r *GameRooms) ResponseList() *[]GameRoomInfo {
	defer r.mu.RUnlock()
	r.mu.RLock()
	rooms := []GameRoomInfo{}
	for i := range r.rooms {
		rooms = append(rooms, GameRoomInfo{
			Uid:         r.rooms[i].Uid,
			Host:        r.rooms[i].host.PlayerInfo,
			Capacity:    r.rooms[i].Capacity,
			JoinedUsers: toPlayerInfos(r.rooms[i].joinedUsers),
		})
	}
	return &rooms
}

func toPlayerInfos(users []users.User) []engine.PlayerInfo {
	result := []engine.PlayerInfo{}
	for i := range users {
		result = append(result, users[i].PlayerInfo)
	}
	return result
}
