package rooms

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

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

func (r *GameRooms) Create(capacity uint, scenarioId GameRoomScenarioId, hostUser users.User) {
	defer r.mu.Unlock()
	r.mu.Lock()
	room := newGameRoom()
	room.Capacity = capacity
	room.ScenarioId = scenarioId
	room.host = hostUser
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

func (r *GameRooms) RemoveUser(userId engine.UserId) (uint, bool) {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return roomUid, false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return roomUid, false
	}
	if room.host.Id == userId {
		return roomUid, false
	}
	delete(r.userIdToRoomUid, userId)
	restUsers := []users.User{}
	for _, u := range room.joinedUsers {
		if u.Id != userId {
			restUsers = append(restUsers, u)
		}
	}
	room.joinedUsers = restUsers
	return roomUid, true
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

func (r *GameRooms) GetUidByUserId(userId engine.UserId) (uint, bool) {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return 0, false
	}
	_, present := r.rooms[roomUid]
	return roomUid, present
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

func (r *GameRooms) ConnectionStatusChanged(userId engine.UserId, isOffline bool) (uint, bool) {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return roomUid, false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return roomUid, false
	}
	return roomUid, room.UpdateUserConnectionStatus(userId, isOffline)
}

func (r *GameRooms) GetAllRoomInfosList() []GameRoomInfo {
	defer r.mu.RUnlock()
	r.mu.RLock()
	rooms := []GameRoomInfo{}
	for i := range r.rooms {
		rooms = append(rooms, toGameRoomInfo(r.rooms[i]))
	}
	return rooms
}

func (r *GameRooms) GetRoomInfoByUid(uid uint) GameRoomInfo {
	defer r.mu.RUnlock()
	r.mu.RLock()
	if room, ok := r.rooms[uid]; ok {
		return toGameRoomInfo(room)
	}
	return toInactiveGameRoomInfo(uid)
}

func (r *GameRooms) GetRoomInfoByUserId(userId engine.UserId) GameRoomInfo {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return GameRoomInfo{}
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return toInactiveGameRoomInfo(roomUid)
	}
	return toGameRoomInfo(room)
}
