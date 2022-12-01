package gameengines

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type gameEngineWrapper struct {
	sync.RWMutex
	engine *engine.GameEngine
}

func newGameEngineWrapper(engine *engine.GameEngine) *gameEngineWrapper {
	w := &gameEngineWrapper{}
	w.engine = engine
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

func (e *GameEngines) Add(engine *engine.GameEngine) {
	defer e.mu.Unlock()
	e.mu.Lock()
	wrapper := newGameEngineWrapper(engine)
	for _, userId := range engine.GetUserIds() {
		e.userIdToEngine[userId] = wrapper
	}
}

func (e *GameEngines) IsUserInGame(userId engine.UserId) bool {
	defer e.mu.RUnlock()
	e.mu.RLock()
	_, ok := e.userIdToEngine[userId]
	return ok
}

func (e *GameEngines) ExecuteUserAction(action domain.Action, userId engine.UserId) (*engine.GameEvent, []engine.UserId, func(), bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, nil, false
	}
	unlock := func() {
		wrapper.Unlock()
	}
	wrapper.Lock()
	result := wrapper.engine.ExecuteUserAction(action, userId)
	broadcastUserIds := []engine.UserId{}
	if result.UnitActionResult.Result.Result == domain.ResultAccomplished {
		broadcastUserIds = wrapper.engine.GetRestUserIds(userId)
	}
	return result, broadcastUserIds, unlock, true
}

func (e *GameEngines) ReadyForNextPhase(userId engine.UserId, isReady bool) (*engine.GameEvent, []engine.UserId, func(), bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, nil, false
	}
	unlock := func() {
		wrapper.Unlock()
	}
	wrapper.Lock()
	if !wrapper.engine.NextPhaseRequired() || wrapper.engine.AllActorsDead() {
		return nil, nil, unlock, false
	}
	wrapper.engine.UpdateActorReady(userId, isReady)
	if wrapper.engine.AllActorsReady() {
		result := wrapper.engine.NextPhase()
		broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
		return result, broadcastUserIds, unlock, true
	}
	result := wrapper.engine.NewGameEventWithPlayersInfo()
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	return result, broadcastUserIds, unlock, true
}

func (e *GameEngines) NextPhase(userId engine.UserId) (*engine.GameEvent, []engine.UserId, func(), bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, nil, false
	}
	unlock := func() {
		wrapper.Unlock()
	}
	wrapper.Lock()
	if !wrapper.engine.NextPhaseRequired() || wrapper.engine.AllActorsDead() {
		return nil, nil, unlock, false
	}
	result := wrapper.engine.NextPhase()
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	return result, broadcastUserIds, unlock, true
}

func (e *GameEngines) GameState(userId engine.UserId) (*engine.GameEvent, []engine.UserId, func(), bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, nil, false
	}
	unlock := func() {
		wrapper.RUnlock()
	}
	wrapper.RLock()
	result := wrapper.engine.NewGameEventWithPlayersInfo()
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	return result, broadcastUserIds, unlock, true
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

func (e *GameEngines) LeaveGame(userId engine.UserId) (*engine.GameEvent, []engine.UserId, domain.Unit, func(), bool) {
	var unit domain.Unit
	e.mu.Lock()
	wrapper, ok := e.userIdToEngine[userId]
	if !ok {
		e.mu.Unlock()
		return nil, nil, unit, nil, false
	}
	delete(e.userIdToEngine, userId)
	e.mu.Unlock()
	unlock := func() {
		wrapper.Unlock()
	}
	wrapper.Lock()
	if u := wrapper.engine.FindActorByUserId(userId); u != nil {
		unit = u.Unit
		if !u.IsDead {
			share := wrapper.engine.TakeAShare()
			unit.Booty.Accumulate(share)
		}
	}
	wrapper.engine.RemoveActor(userId)
	userIds := wrapper.engine.GetUserIds()
	state := wrapper.engine.NewGameEventWithPlayersInfo()
	if len(userIds) == 0 {
		wrapper.engine.Dispose()
	}
	return state, userIds, unit, unlock, true
}

func (e *GameEngines) RemoveUser(userId engine.UserId) (*engine.GameEvent, []engine.UserId, func(), bool) {
	e.mu.Lock()
	wrapper, ok := e.userIdToEngine[userId]
	if !ok {
		e.mu.Unlock()
		return nil, nil, nil, false
	}
	delete(e.userIdToEngine, userId)
	e.mu.Unlock()
	unlock := func() {
		wrapper.Unlock()
	}
	wrapper.Lock()
	wrapper.engine.RemoveActor(userId)
	userIds := wrapper.engine.GetUserIds()
	if len(userIds) == 0 {
		wrapper.engine.Dispose()
		return nil, nil, unlock, false
	}
	state := wrapper.engine.NewGameEventWithPlayersInfo()
	return state, userIds, unlock, true
}

func (e *GameEngines) ConnectionStatusChanged(userId engine.UserId, isOffline bool) (*engine.GameEvent, []engine.UserId, func(), bool) {
	e.mu.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.mu.RUnlock()
	if !ok {
		return nil, nil, nil, false
	}
	unlock := func() {
		wrapper.Unlock()
	}
	wrapper.Lock()
	wrapper.engine.UpdateUserConnectionStatus(userId, isOffline)
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	state := wrapper.engine.NewGameEventWithPlayersInfo()
	return state, broadcastUserIds, unlock, true
}
