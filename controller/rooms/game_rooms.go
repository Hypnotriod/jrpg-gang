package rooms

import (
	"jrpg-gang/controller/users"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameRooms struct {
	mu                sync.RWMutex
	rndGen            *util.RndGen
	rooms             map[uint]*GameRoom
	playerIdToRoomUid map[engine.PlayerId]uint
}

func NewGameRooms() *GameRooms {
	r := &GameRooms{}
	r.rndGen = util.NewRndGen()
	r.rooms = make(map[uint]*GameRoom)
	r.playerIdToRoomUid = make(map[engine.PlayerId]uint)
	return r
}

func (r *GameRooms) Create(capacity uint, scenarioId engine.GameScenarioId, hostUser users.User) {
	defer r.mu.Unlock()
	r.mu.Lock()
	room := newGameRoom()
	room.Capacity = capacity
	room.ScenarioId = scenarioId
	room.host = hostUser
	room.Uid = r.rndGen.NextUid()
	r.rooms[room.Uid] = room
	r.playerIdToRoomUid[room.host.Id] = room.Uid
}

func (r *GameRooms) PopByHostId(hostId engine.PlayerId) (*GameRoom, bool) {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.playerIdToRoomUid[hostId]
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
		delete(r.playerIdToRoomUid, u.Id)
	}
	delete(r.playerIdToRoomUid, hostId)
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
	r.playerIdToRoomUid[user.Id] = roomUid
	return true
}

func (r *GameRooms) RemoveUser(playerId engine.PlayerId) (uint, bool) {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.playerIdToRoomUid[playerId]
	if !ok {
		return roomUid, false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return roomUid, false
	}
	if room.host.Id == playerId {
		return roomUid, false
	}
	delete(r.playerIdToRoomUid, playerId)
	restUsers := []users.User{}
	for _, u := range room.joinedUsers {
		if u.Id != playerId {
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

func (r *GameRooms) ExistsForPlayerId(playerId engine.PlayerId) bool {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.playerIdToRoomUid[playerId]
	if !ok {
		return false
	}
	_, present := r.rooms[roomUid]
	return present
}

func (r *GameRooms) GetUidByPlayerId(playerId engine.PlayerId) (uint, bool) {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.playerIdToRoomUid[playerId]
	if !ok {
		return 0, false
	}
	_, present := r.rooms[roomUid]
	return roomUid, present
}

func (r *GameRooms) ExistsForHostId(hostId engine.PlayerId) bool {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.playerIdToRoomUid[hostId]
	if !ok {
		return false
	}
	room, present := r.rooms[roomUid]
	if !present {
		return false
	}
	return room.host.Id == hostId
}

func (r *GameRooms) ConnectionStatusChanged(playerId engine.PlayerId, isOffline bool) (uint, bool) {
	defer r.mu.Unlock()
	r.mu.Lock()
	roomUid, ok := r.playerIdToRoomUid[playerId]
	if !ok {
		return roomUid, false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return roomUid, false
	}
	return roomUid, room.UpdateUserConnectionStatus(playerId, isOffline)
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

func (r *GameRooms) GetRoomInfoByPlayerId(playerId engine.PlayerId) GameRoomInfo {
	defer r.mu.RUnlock()
	r.mu.RLock()
	roomUid, ok := r.playerIdToRoomUid[playerId]
	if !ok {
		return GameRoomInfo{}
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return toInactiveGameRoomInfo(roomUid)
	}
	return toGameRoomInfo(room)
}
