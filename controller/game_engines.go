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
}

func NewGameEngineWrapper() *GameEngineWrapper {
	w := &GameEngineWrapper{}
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

func (e *GameEngines) Add(engine *engine.GameEngine) {
	defer e.Unlock()
	e.Lock()
	wrapper := NewGameEngineWrapper()
	wrapper.engine = engine
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
	wrapper.Lock()
	result := wrapper.engine.ExecuteUserAction(action, userId)
	broadcastUserIds := []engine.UserId{}
	if result.UnitActionResult.Result.ResultType == domain.ResultAccomplished {
		broadcastUserIds = wrapper.engine.GetRestUserIds(userId)
	}
	wrapper.Unlock()
	return result, broadcastUserIds, true
}
