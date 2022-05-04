package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameRoom struct {
	*sync.RWMutex
	Uid         uint   `json:"uid"`
	Capacity    uint   `json:"capacity"`
	ScenarioUid uint   `json:"scenarioUid"`
	JoinedUsers []User `json:"joinedUsers"`
	Host        User   `json:"host"`
}

func (r *GameRoom) IsFull() bool {
	return len(r.JoinedUsers) >= int(r.Capacity)-1
}

func (r *GameRoom) GetActors() []*engine.GameUnit {
	defer r.RUnlock()
	r.RLock()
	result := []*engine.GameUnit{}
	for _, u := range r.JoinedUsers {
		result = append(result, u.unit)
	}
	result = append(result, r.Host.unit)
	return result
}

func NewGameRoom() *GameRoom {
	r := &GameRoom{}
	r.RWMutex = &sync.RWMutex{}
	r.JoinedUsers = []User{}
	return r
}

type GameRooms struct {
	sync.RWMutex
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
	defer r.Unlock()
	r.Lock()
	room.Uid = r.rndGen.NextUid()
	r.rooms[room.Uid] = room
	r.userIdToRoomUid[room.Host.id] = room.Uid
}

func (r *GameRooms) RemoveByHostId(hostId engine.UserId) bool {
	defer r.Unlock()
	r.Lock()
	roomUid, ok := r.userIdToRoomUid[hostId]
	if !ok {
		return false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return false
	}
	for _, u := range room.JoinedUsers {
		delete(r.userIdToRoomUid, u.id)
	}
	delete(r.userIdToRoomUid, hostId)
	delete(r.rooms, roomUid)
	return true
}

func (r *GameRooms) AddUser(roomUid uint, user User) bool {
	defer r.Unlock()
	r.Lock()
	room, ok := r.rooms[roomUid]
	if !ok {
		return false
	}
	if room.IsFull() {
		return false
	}
	room.JoinedUsers = append(room.JoinedUsers, user)
	r.userIdToRoomUid[user.id] = roomUid
	return true
}

func (r *GameRooms) RemoveUser(userId engine.UserId) bool {
	defer r.Unlock()
	r.Lock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return false
	}
	if room.Host.id == userId {
		return false
	}
	delete(r.userIdToRoomUid, userId)
	restUsers := []User{}
	for _, u := range room.JoinedUsers {
		if u.id != userId {
			restUsers = append(restUsers, u)
		}
	}
	room.JoinedUsers = restUsers
	return true
}

func (r *GameRooms) Has(uid uint) bool {
	defer r.RUnlock()
	r.RLock()
	_, ok := r.rooms[uid]
	return ok
}

func (r *GameRooms) ExistsForUserId(userId engine.UserId) bool {
	defer r.RUnlock()
	r.RLock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return false
	}
	_, present := r.rooms[roomUid]
	return present
}

func (r *GameRooms) ExistsForHostId(hostId engine.UserId) bool {
	defer r.RUnlock()
	r.RLock()
	roomUid, ok := r.userIdToRoomUid[hostId]
	if !ok {
		return false
	}
	room, present := r.rooms[roomUid]
	if !present {
		return false
	}
	return room.Host.id == hostId
}

func (r *GameRooms) GetByUserId(userId engine.UserId) (GameRoom, bool) {
	defer r.RUnlock()
	r.RLock()
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

func (r *GameRooms) ResponseList() *[]GameRoom {
	defer r.RUnlock()
	r.RLock()
	rooms := []GameRoom{}
	for i := range r.rooms {
		rooms = append(rooms, GameRoom{
			Uid:         r.rooms[i].Uid,
			Host:        r.rooms[i].Host,
			Capacity:    r.rooms[i].Capacity,
			JoinedUsers: r.rooms[i].JoinedUsers,
		})
	}
	return &rooms
}
