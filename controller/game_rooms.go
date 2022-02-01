package controller

import (
	"jrpg-gang/util"
	"sync"
)

type GameRoom struct {
	*sync.RWMutex
	isFull      bool
	Uid         uint   `json:"uid"`
	Capacity    uint   `json:"capacity"`
	JoinedUsers []User `json:"joinedUsers"`
	Host        User   `json:"host"`
	// engine   *engine.GameEngine
}

func NewGameRoom() *GameRoom {
	r := &GameRoom{}
	r.RWMutex = &sync.RWMutex{}
	r.JoinedUsers = []User{}
	return r
}

type GameRooms struct {
	sync.RWMutex
	uidGen          util.UidGen
	rooms           map[uint]*GameRoom
	userIdToRoomUid map[UserId]uint
}

func NewGameRooms() *GameRooms {
	r := &GameRooms{}
	r.rooms = make(map[uint]*GameRoom)
	r.userIdToRoomUid = make(map[UserId]uint)
	return r
}

func (r *GameRooms) Add(room *GameRoom) {
	defer r.Unlock()
	r.Lock()
	room.Uid = r.uidGen.Next()
	room.isFull = len(room.JoinedUsers) >= int(room.Capacity)-1
	r.rooms[room.Uid] = room
	r.userIdToRoomUid[room.Host.id] = room.Uid
}

func (r *GameRooms) Remove(roomUid uint) bool {
	defer r.Unlock()
	r.Lock()
	if _, ok := r.rooms[roomUid]; !ok {
		return false
	}
	delete(r.rooms, roomUid)
	return true
}

func (r *GameRooms) RemoveByUserId(userId UserId) bool {
	defer r.Unlock()
	r.Lock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return false
	}
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
	if room.isFull {
		return false
	}
	room.JoinedUsers = append(room.JoinedUsers, user)
	room.isFull = len(room.JoinedUsers) >= int(room.Capacity)-1
	r.userIdToRoomUid[user.id] = roomUid
	return true
}

func (r *GameRooms) Has(uid uint) bool {
	defer r.RUnlock()
	r.RLock()
	_, ok := r.rooms[uid]
	return ok
}

func (r *GameRooms) ExistsForUserId(userId UserId) bool {
	defer r.RUnlock()
	r.RLock()
	roomUid, ok := r.userIdToRoomUid[userId]
	if !ok {
		return false
	}
	_, present := r.rooms[roomUid]
	return present
}

func (r *GameRooms) GetByUserId(userId UserId) (GameRoom, bool) {
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
		if !r.rooms[i].isFull {
			rooms = append(rooms, GameRoom{
				Uid:         r.rooms[i].Uid,
				Host:        r.rooms[i].Host,
				Capacity:    r.rooms[i].Capacity,
				JoinedUsers: r.rooms[i].JoinedUsers,
			})
		}
	}
	return &rooms
}
