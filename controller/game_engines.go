package controller

import (
	"jrpg-gang/domain"
	"jrpg-gang/engine"
	"jrpg-gang/util"
	"sync"
)

type GameEngineWrapper struct {
	sync.RWMutex
	engine *engine.GameEngine
	hostId engine.UserId
}

func NewGameEngineWrapper(hostId engine.UserId, engine *engine.GameEngine) *GameEngineWrapper {
	w := &GameEngineWrapper{}
	w.engine = engine
	w.hostId = hostId
	return w
}

type GameEngines struct {
	sync.RWMutex
	rndGen         *util.RndGen
	engines        []*GameEngineWrapper
	userIdToEngine map[engine.UserId]*GameEngineWrapper
}

func NewGameEngines() *GameEngines {
	e := &GameEngines{}
	e.rndGen = util.NewRndGen()
	e.userIdToEngine = make(map[engine.UserId]*GameEngineWrapper)
	return e
}

func (e *GameEngines) Add(hostId engine.UserId, engine *engine.GameEngine) {
	defer e.Unlock()
	e.Lock()
	wrapper := NewGameEngineWrapper(hostId, engine)
	e.engines = append(e.engines, wrapper)
	for _, userId := range engine.GetUserIds() {
		e.userIdToEngine[userId] = wrapper
	}
}

func (e *GameEngines) IsUserInGame(userId engine.UserId) bool {
	defer e.RUnlock()
	e.RLock()
	_, ok := e.userIdToEngine[userId]
	return ok
}

func (e *GameEngines) ExecuteUserAction(action domain.Action, userId engine.UserId) (*engine.GameEvent, []engine.UserId, bool) {
	e.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.RUnlock()
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
	e.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.RUnlock()
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
	e.RLock()
	wrapper, ok := e.userIdToEngine[userId]
	e.RUnlock()
	if !ok {
		return nil, nil, false
	}
	wrapper.Lock()
	result := wrapper.engine.NewGameEvent()
	broadcastUserIds := wrapper.engine.GetRestUserIds(userId)
	wrapper.Unlock()
	return result, broadcastUserIds, true
}
