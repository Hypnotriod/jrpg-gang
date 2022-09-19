package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type gameEngineWrapper struct {
	sync.RWMutex
	engine *engine.GameEngine
	hostId engine.UserId
}

func newGameEngineWrapper(hostId engine.UserId, engine *engine.GameEngine) *gameEngineWrapper {
	w := &gameEngineWrapper{}
	w.engine = engine
	w.hostId = hostId
	return w
}

type GameEngines struct {
	mu             sync.RWMutex
	rndGen         *util.RndGen
	userIdToEngine map[engine.UserId]*gameEngineWrapper
}

func NewGameEngines() *GameEngines {
	e := &GameEngines{}
	e.rndGen = util.NewRndGen()
	e.userIdToEngine = make(map[engine.UserId]*gameEngineWrapper)
	return e
}

func (e *GameEngines) Add(hostId engine.UserId, engine *engine.GameEngine) {
	defer e.mu.Unlock()
	e.mu.Lock()
	wrapper := newGameEngineWrapper(hostId, engine)
	for _, userId := range engine.GetUserIds() {
		e.userIdToEngine[userId] = wrapper
	}
}

func (e *GameEngines) ChangeHost(oldHostId engine.UserId, newHostId engine.UserId) bool {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[oldHostId]
	e.mu.RUnlock()
	if !ok {
		return false
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	if wrapper.hostId != oldHostId {
		return false
	}
	wrapper.hostId = newHostId
	if u := wrapper.engine.FindActorByUserId(oldHostId); u != nil {
		u.PlayerInfo.IsHost = false
	}
	if u := wrapper.engine.FindActorByUserId(newHostId); u != nil {
		u.PlayerInfo.IsHost = true
	}
	return true
}

func (e *GameEngines) IsUserInGame(userId engine.UserId) bool {
	defer e.mu.RUnlock()
	e.mu.RLock()
	_, ok := e.userIdToEngine[userId]
	return ok
}

func (e *GameEngines) ExecuteUserAction(action domain.Action, userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, false
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	result := wrapper.engine.ExecuteUserAction(action, userId)
	broadcastUserIds := []engine.UserId{}
	if result.UnitActionResult.Result.Result == domain.ResultAccomplished {
		broadcastUserIds = wrapper.engine.GetRestUserIds(userId)
	}
	return result, broadcastUserIds, true
}

func (e *GameEngines) NextPhase(userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, false
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	if wrapper.hostId != userId || !wrapper.engine.NextPhaseRequired() {
		return nil, nil, false
	}
	result := wrapper.engine.NextPhase()
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	return result, broadcastUserIds, true
}

func (e *GameEngines) GameState(userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, false
	}
	defer wrapper.RUnlock()
	wrapper.RLock()
	result := wrapper.engine.NewGameEvent()
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	return result, broadcastUserIds, true
}

func (e *GameEngines) PlayerInfo(userId engine.UserId) (engine.PlayerInfo, bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return engine.PlayerInfo{}, false
	}
	defer wrapper.RUnlock()
	wrapper.RLock()
	unit := wrapper.engine.FindActorByUserId(userId)
	if unit == nil || unit.PlayerInfo == nil {
		return engine.PlayerInfo{}, false
	}
	return *unit.PlayerInfo, true
}

func (e *GameEngines) RemoveUser(userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	e.mu.Lock()
	wrapper, ok := e.userIdToEngine[userId]
	if !ok {
		e.mu.Unlock()
		return nil, nil, false
	}
	delete(e.userIdToEngine, userId)
	e.mu.Unlock()
	defer wrapper.Unlock()
	wrapper.Lock()
	wrapper.engine.RemoveActor(userId)
	userIds := wrapper.engine.GetUserIds()
	if len(userIds) == 0 {
		wrapper.engine.Dispose()
		return nil, nil, false
	}
	if wrapper.hostId == userId {
		wrapper.hostId = userIds[0]
		if u := wrapper.engine.FindActorByUserId(wrapper.hostId); u != nil {
			u.PlayerInfo.IsHost = true
		}
	}
	broadcastUserIds := wrapper.engine.GetUserIds()
	state := wrapper.engine.NewGameEvent()
	return state, broadcastUserIds, true
}

func (e *GameEngines) ConnectionStatusChanged(userId engine.UserId, isOffline bool) (*engine.GameEvent, []engine.UserId, bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, false
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	wrapper.engine.UpdateUserConnectionStatus(userId, isOffline)
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	state := wrapper.engine.NewGameEvent()
	return state, broadcastUserIds, true
}
