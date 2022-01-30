package controller

import (
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameRoom struct {
	Uid      uint   `json:"uid"`
	HostId   string `json:"hostId"`
	Capacity uint   `json:"capacity"`
	UserIds  []uint `json:"userIds"`
	Engine   *engine.GameEngine
}

type GameRooms struct {
	sync.RWMutex
	uidGen         util.UidGen
	rooms          map[uint]*GameRoom
	userIdToRoomId map[string]uint
}

func NewGameRooms() *GameRooms {
	r := &GameRooms{}
	r.rooms = make(map[uint]*GameRoom)
	r.userIdToRoomId = make(map[string]uint)
	return r
}

func (r *GameRooms) Add(room *GameRoom) {
	defer r.Unlock()
	r.Lock()
	roomUid := r.uidGen.Next()
	r.rooms[roomUid] = room
	r.userIdToRoomId[room.HostId] = roomUid
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

func (r *GameRooms) RemoveByUserId(userId string) bool {
	defer r.Unlock()
	r.Lock()
	roomUid, ok := r.userIdToRoomId[userId]
	if !ok {
		return false
	}
	delete(r.rooms, roomUid)
	return true
}

func (r *GameRooms) ExistsForUserId(userId string) bool {
	defer r.RUnlock()
	r.RLock()
	roomUid, ok := r.userIdToRoomId[userId]
	if !ok {
		return false
	}
	_, present := r.rooms[roomUid]
	return present
}

func (r *GameRooms) GetByUserId(userId string) (GameRoom, bool) {
	defer r.RUnlock()
	r.RLock()
	roomUid, ok := r.userIdToRoomId[userId]
	if !ok {
		return GameRoom{}, false
	}
	room, ok := r.rooms[roomUid]
	if !ok {
		return GameRoom{}, false
	}
	return *room, ok
}
